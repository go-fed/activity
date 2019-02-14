package pub

import (
	"bytes"
	"context"
	"crypto"
	"fmt"
	"github.com/go-fed/httpsig"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	// acceptHeaderValue is the Accept header value indicating that the
	// response should contain an ActivityStreams object.
	acceptHeaderValue = "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""
)

// Transport makes ActivityStreams calls to other servers in order to POST or
// GET ActivityStreams data.
//
// It may be reused multiple times, but never concurrently.
type Transport interface {
	// Dereference fetches the ActivityStreams object located at this IRI
	// with a GET request.
	Dereference(c context.Context, iri *url.URL) ([]byte, error)
	// Deliver sends an ActivityStreams object.
	Deliver(c context.Context, b []byte, to *url.URL) error
	// BatchDeliver sends an ActivityStreams object to multiple recipients.
	BatchDeliver(c context.Context, b []byte, recipients []*url.URL) error
}

// Transport must be implemented by HttpSigTransport.
var _ Transport = &HttpSigTransport{}

// HttpSigTransport makes a dereference call using HTTP signatures to
// authenticate the request on behalf of a particular actor.
//
// No rate limiting is applied.
//
// Only one request is tried per call.
type HttpSigTransport struct {
	client     HttpClient
	appAgent   string
	gofedAgent string
	clock      Clock
	signer     httpsig.Signer
	pubKeyId   string
	privKey    crypto.PrivateKey
}

// NewHttpSigTransport returns a new HttpSigTransport.
func NewHttpSigTransport(
	client HttpClient,
	appAgent, gofedAgent string,
	clock Clock,
	signer httpsig.Signer,
	pubKeyId string,
	privKey crypto.PrivateKey) *HttpSigTransport {
	return &HttpSigTransport{
		client:     client,
		appAgent:   appAgent,
		gofedAgent: gofedAgent,
		clock:      clock,
		signer:     signer,
		pubKeyId:   pubKeyId,
		privKey:    privKey,
	}
}

// Dereferences with a request signed with an HTTP Signature.
func (h HttpSigTransport) Dereference(c context.Context, iri *url.URL) ([]byte, error) {
	req, err := http.NewRequest("GET", iri.String(), nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(c)
	req.Header.Add(acceptHeader, acceptHeaderValue)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Date", h.clock.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	req.Header.Add("User-Agent", fmt.Sprintf("%s %s", h.appAgent, h.gofedAgent))
	err = h.signer.SignRequest(h.privKey, h.pubKeyId, req)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET request to %s failed (%d): %s", iri.String(), resp.StatusCode, resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}

// Deliver sends a POST request with an HTTP Signature.
func (h HttpSigTransport) Deliver(c context.Context, b []byte, to *url.URL) error {
	byteCopy := make([]byte, len(b))
	copy(byteCopy, b)
	buf := bytes.NewBuffer(byteCopy)
	req, err := http.NewRequest("POST", to.String(), buf)
	if err != nil {
		return err
	}
	req.WithContext(c)
	req.Header.Add(contentTypeHeader, contentTypeHeaderValue)
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Date", h.clock.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05")+" GMT")
	req.Header.Add("User-Agent", fmt.Sprintf("%s %s", h.appAgent, h.gofedAgent))
	err = h.signer.SignRequest(h.privKey, h.pubKeyId, req)
	if err != nil {
		return err
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("POST request to %s failed (%d): %s", to.String(), resp.StatusCode, resp.Status)
	}
	return nil
}

// BatchDeliver sends concurrent POST requests. Returns an error if any of the
// requests had an error.
func (h HttpSigTransport) BatchDeliver(c context.Context, b []byte, recipients []*url.URL) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(recipients))
	for _, recipient := range recipients {
		wg.Add(1)
		go func(r *url.URL) {
			defer wg.Done()
			if err := h.Deliver(c, b, r); err != nil {
				errCh <- err
			}
		}(recipient)
	}
	wg.Wait()
	errs := make([]string, 0, len(recipients))
	for {
		select {
		case e := <-errCh:
			errs = append(errs, e.Error())
		default:
			break
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("batch deliver had at least one failure: %s", strings.Join(errs, "; "))
	}
	return nil
}

// HttpClient sends http requests, and is an abstraction only needed by the
// HttpSigTransport. The standard library's Client satisfies this interface.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// HttpClient must be implemented by http.Client.
var _ HttpClient = &http.Client{}
