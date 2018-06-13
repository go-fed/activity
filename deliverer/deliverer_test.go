package deliverer

import (
	"fmt"
	"github.com/go-test/deep"
	"golang.org/x/time/rate"
	"net/url"
	"sync"
	"testing"
	"time"
)

const (
	id1           = "id1"
	id2           = "id2"
	sending       = "sending"
	cancel        = "cancel"
	successful    = "successful"
	retrying      = "retrying"
	undeliverable = "undeliverable"
	noState       = "noState"
)

var (
	testBytes []byte = []byte{0, 1, 2, 3}
	testURL   *url.URL
)

func init() {
	var err error
	testURL, err = url.Parse("example.com")
	if err != nil {
		panic(err)
	}
}

var _ DeliveryPersister = &mockDeliveryPersister{}

type mockDeliveryPersister struct {
	t        *testing.T
	i        int
	mu       *sync.Mutex
	id1State string
	id2State string
}

func newMockDeliveryPersister(t *testing.T) *mockDeliveryPersister {
	return &mockDeliveryPersister{
		t:        t,
		mu:       &sync.Mutex{},
		id1State: noState,
		id2State: noState,
	}
}

func (m *mockDeliveryPersister) Sending(b []byte, to *url.URL) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.i == 0 {
		m.i++
		return id1
	} else if m.i == 1 {
		m.i++
		return id2
	} else {
		m.t.Fatal("too many calls to Sending")
	}
	return ""
}

func (m *mockDeliveryPersister) Cancel(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if id == id1 {
		m.id1State = cancel
	} else if id == id2 {
		m.id2State = cancel
	} else {
		m.t.Fatalf("unknown Cancel id: %s", id)
	}
}

func (m *mockDeliveryPersister) Successful(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if id == id1 {
		m.id1State = successful
	} else if id == id2 {
		m.id2State = successful
	} else {
		m.t.Fatalf("unknown Successful id: %s", id)
	}
}

func (m *mockDeliveryPersister) Retrying(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if id == id1 {
		m.id1State = retrying
	} else if id == id2 {
		m.id2State = retrying
	} else {
		m.t.Fatalf("unknown Retrying id: %s", id)
	}
}

func (m *mockDeliveryPersister) Undeliverable(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if id == id1 {
		m.id1State = undeliverable
	} else if id == id2 {
		m.id2State = undeliverable
	} else {
		m.t.Fatalf("unknown Retrying id: %s", id)
	}
}

func TestDelivererPoolSuccessNoPersister(t *testing.T) {
	testSendFn := func(b []byte, u *url.URL) error {
		if diff := deep.Equal(b, testBytes); diff != nil {
			t.Fatal(diff)
		} else if u != testURL {
			t.Fatal("wrong testURL")
		}
		return nil
	}
	pool := NewDelivererPool(DeliveryOptions{
		InitialRetryTime: time.Microsecond,
		MaximumRetryTime: time.Microsecond,
		BackoffFactor:    2,
		MaxRetries:       1,
		RateLimit:        rate.NewLimiter(1, 1),
	})
	pool.Do(testBytes, testURL, testSendFn)
	time.Sleep(time.Microsecond * 500)
}

func TestDelivererPoolSuccessPersister(t *testing.T) {
	testSendFn := func(b []byte, u *url.URL) error {
		if diff := deep.Equal(b, testBytes); diff != nil {
			t.Fatal(diff)
		} else if u != testURL {
			t.Fatal("wrong testURL")
		}
		return nil
	}
	p := newMockDeliveryPersister(t)
	pool := NewDelivererPool(DeliveryOptions{
		InitialRetryTime: time.Microsecond,
		MaximumRetryTime: time.Microsecond,
		BackoffFactor:    2,
		MaxRetries:       1,
		RateLimit:        rate.NewLimiter(1, 1),
		Persister:        p,
	})
	pool.Do(testBytes, testURL, testSendFn)
	time.Sleep(time.Microsecond * 500)
	if p.id1State != successful {
		t.Fatalf("want: %s, got %s", successful, p.id1State)
	}
}

func TestRestartSuccess(t *testing.T) {
	testSendFn := func(b []byte, u *url.URL) error {
		if diff := deep.Equal(b, testBytes); diff != nil {
			t.Fatal(diff)
		} else if u != testURL {
			t.Fatal("wrong testURL")
		}
		return nil
	}
	p := newMockDeliveryPersister(t)
	pool := NewDelivererPool(DeliveryOptions{
		InitialRetryTime: time.Microsecond,
		MaximumRetryTime: time.Microsecond,
		BackoffFactor:    2,
		MaxRetries:       1,
		RateLimit:        rate.NewLimiter(1, 1),
		Persister:        p,
	})
	pool.Restart(testBytes, testURL, id2, testSendFn)
	time.Sleep(time.Microsecond * 500)
	if p.id2State != successful {
		t.Fatalf("want: %s, got %s", successful, p.id1State)
	}
}

func TestDelivererPoolRetrying(t *testing.T) {
	testSendFn := func(b []byte, u *url.URL) error {
		if diff := deep.Equal(b, testBytes); diff != nil {
			t.Fatal(diff)
		} else if u != testURL {
			t.Fatal("wrong testURL")
		}
		return fmt.Errorf("expected")
	}
	p := newMockDeliveryPersister(t)
	pool := NewDelivererPool(DeliveryOptions{
		InitialRetryTime: time.Microsecond,
		MaximumRetryTime: time.Microsecond,
		BackoffFactor:    2,
		MaxRetries:       1,
		RateLimit:        rate.NewLimiter(1000000, 10000000),
		Persister:        p,
	})
	pool.Do(testBytes, testURL, testSendFn)
	time.Sleep(time.Microsecond * 500)
	select {
	case <-pool.Errors():
	default:
		t.Fatal("expected error")
	}
	time.Sleep(time.Microsecond * 500)
	if p.id1State != retrying {
		t.Fatalf("want: %s, got %s", retrying, p.id1State)
	}
}

func TestDelivererPoolUndeliverable(t *testing.T) {
	testSendFn := func(b []byte, u *url.URL) error {
		if diff := deep.Equal(b, testBytes); diff != nil {
			t.Fatal(diff)
		} else if u != testURL {
			t.Fatal("wrong testURL")
		}
		return fmt.Errorf("expected")
	}
	p := newMockDeliveryPersister(t)
	pool := NewDelivererPool(DeliveryOptions{
		InitialRetryTime: time.Microsecond,
		MaximumRetryTime: time.Microsecond,
		BackoffFactor:    2,
		MaxRetries:       1,
		RateLimit:        rate.NewLimiter(1000000, 10000000),
		Persister:        p,
	})
	pool.Do(testBytes, testURL, testSendFn)
	time.Sleep(time.Microsecond * 500)
	<-pool.Errors()
	time.Sleep(time.Microsecond * 500)
	<-pool.Errors()
	time.Sleep(time.Microsecond * 500)
	<-pool.Errors()
	time.Sleep(time.Microsecond * 500)
	if p.id1State != undeliverable {
		t.Fatalf("want: %s, got %s", undeliverable, p.id1State)
	}
}

func TestRestartRetrying(t *testing.T) {
	testSendFn := func(b []byte, u *url.URL) error {
		if diff := deep.Equal(b, testBytes); diff != nil {
			t.Fatal(diff)
		} else if u != testURL {
			t.Fatal("wrong testURL")
		}
		return fmt.Errorf("expected")
	}
	p := newMockDeliveryPersister(t)
	pool := NewDelivererPool(DeliveryOptions{
		InitialRetryTime: time.Microsecond,
		MaximumRetryTime: time.Microsecond,
		BackoffFactor:    2,
		MaxRetries:       1,
		RateLimit:        rate.NewLimiter(1000000, 10000000),
		Persister:        p,
	})
	pool.Restart(testBytes, testURL, id2, testSendFn)
	time.Sleep(time.Microsecond * 500)
	select {
	case <-pool.Errors():
	default:
		t.Fatal("expected error")
	}
	time.Sleep(time.Microsecond * 500)
	if p.id2State != retrying {
		t.Fatalf("want: %s, got %s", retrying, p.id2State)
	}
}

func TestRestartUndeliverable(t *testing.T) {
	testSendFn := func(b []byte, u *url.URL) error {
		if diff := deep.Equal(b, testBytes); diff != nil {
			t.Fatal(diff)
		} else if u != testURL {
			t.Fatal("wrong testURL")
		}
		return fmt.Errorf("expected")
	}
	p := newMockDeliveryPersister(t)
	pool := NewDelivererPool(DeliveryOptions{
		InitialRetryTime: time.Microsecond,
		MaximumRetryTime: time.Microsecond,
		BackoffFactor:    2,
		MaxRetries:       1,
		RateLimit:        rate.NewLimiter(1000000, 10000000),
		Persister:        p,
	})
	pool.Restart(testBytes, testURL, id2, testSendFn)
	time.Sleep(time.Microsecond * 500)
	<-pool.Errors()
	time.Sleep(time.Microsecond * 500)
	<-pool.Errors()
	time.Sleep(time.Microsecond * 500)
	<-pool.Errors()
	time.Sleep(time.Microsecond * 500)
	if p.id2State != undeliverable {
		t.Fatalf("want: %s, got %s", undeliverable, p.id2State)
	}
}
