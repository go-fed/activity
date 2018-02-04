package pub

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"math"
	"net/url"
	"sync"
	"time"
)

// DeliveryPersister allows applications to keep track of delivery states of
// the messages being sent, including during retries. This permits clients to
// also resume delivery of messages that were in the process of being delivered
// when the application server was shut down.
type DeliveryPersister interface {
	// Sending informs the delivery persister that the provided bytes are
	// being delivered to the specified url. It must return a unique id for
	// this delivery.
	Sending(b []byte, to url.URL) string
	// Cancel informs the delivery persister that the provided delivery was
	// interrupted by the server cancelling. These should be retried once
	// the server is back online.
	Cancel(id string)
	// Successful informs the delivery persister that the request has been
	// successfully delivered and no further retries are needed.
	Successful(id string)
	// Retrying indicates the specified delivery is being retried.
	Retrying(id string)
	// Undeliverable indicates the specified delivery has failed and is no
	// longer being retried.
	Undeliverable(id string)
}

type delivererPool struct {
	// When present, permits clients to be notified of all state changes
	// when delivering a request to another federated server.
	//
	// Optional.
	persister DeliveryPersister
	// Limit speed of retries.
	initialRetryTime time.Duration
	maxRetryTime     time.Duration
	retryTimeFactor  float64
	// Limit total number of retries.
	maxNumberRetries int
	// Enforces speed limit of retries
	limiter *rate.Limiter
	// Allow graceful cancelling
	ctx      context.Context
	cancel   context.CancelFunc
	timerId  uint64
	timerMap map[uint64]*time.Timer
	mu       sync.Mutex // Limits concurrent access to timerId and timerMap
	// Allow graceful error handling
	errChan chan error
}

func newDelivererPool(initial, maxRetry time.Duration, factor float64, maxN int, qps *rate.Limiter, persister DeliveryPersister) *delivererPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &delivererPool{
		persister:        persister,
		initialRetryTime: initial,
		maxRetryTime:     maxRetry,
		retryTimeFactor:  factor,
		maxNumberRetries: maxN,
		limiter:          qps,
		ctx:              ctx,
		cancel:           cancel,
		timerId:          0,
		timerMap:         make(map[uint64]*time.Timer, 0),
		mu:               sync.Mutex{},
		errChan:          make(chan error, 0),
	}
}

type retryData struct {
	nextWait time.Duration
	n        int
	f        func() error
	id       string
}

func (r retryData) NextRetry(factor float64, max time.Duration) retryData {
	w := time.Duration(int64(math.Floor((float64(r.nextWait) * factor) + 0.5)))
	if w > max {
		w = max
	}
	return retryData{
		nextWait: w,
		n:        r.n + 1,
		f:        r.f,
		id:       r.id,
	}
}

func (r retryData) ShouldRetry(max int) bool {
	return r.n < max
}

// Do spawns a goroutine that retries f until it returns no error.
func (d *delivererPool) Do(b []byte, to url.URL, sendFn func([]byte, url.URL) error) {
	f := func() error {
		return sendFn(b, to)
	}
	go func() {
		id := ""
		if d.persister != nil {
			id = d.persister.Sending(b, to)
		}
		d.do(retryData{
			nextWait: d.initialRetryTime,
			n:        0,
			f:        f,
			id:       id,
		})
	}()
}

func (d *delivererPool) Restart(b []byte, to url.URL, id string, sendFn func([]byte, url.URL) error) {
	f := func() error {
		return sendFn(b, to)
	}
	go func() {
		d.do(retryData{
			nextWait: d.initialRetryTime,
			n:        0,
			f:        f,
			id:       id,
		})
	}()
}

func (d *delivererPool) Stop() {
	d.cancel()
	d.closeTimers()
}

func (d *delivererPool) Errors() <-chan error {
	return d.errChan
}

func (d *delivererPool) do(r retryData) {
	if err := d.limiter.Wait(d.ctx); err != nil {
		if d.persister != nil {
			d.persister.Cancel(r.id)
		}
		d.errChan <- err
		return
	}
	if err := r.f(); err != nil {
		d.errChan <- err
		if r.ShouldRetry(d.maxNumberRetries) {
			if d.persister != nil {
				d.persister.Retrying(r.id)
			}
			d.addClosableTimer(r)
		} else {
			d.errChan <- fmt.Errorf("delivery tried maximum number of times")
			if d.persister != nil {
				d.persister.Undeliverable(r.id)
			}
		}
	}
	if d.persister != nil {
		d.persister.Successful(r.id)
	}
}

func (d *delivererPool) addClosableTimer(r retryData) {
	d.mu.Lock()
	defer d.mu.Unlock()
	id := d.timerId
	d.timerId++
	d.timerMap[id] = time.AfterFunc(r.nextWait, func() {
		d.do(r.NextRetry(d.retryTimeFactor, d.maxRetryTime))
		d.removeTimer(id)
	})
}

func (d *delivererPool) removeTimer(id uint64) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.timerMap[id]; ok {
		delete(d.timerMap, id)
	}
}

func (d *delivererPool) closeTimers() {
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, v := range d.timerMap {
		v.Stop()
	}
	d.timerMap = make(map[uint64]*time.Timer, 0)
}
