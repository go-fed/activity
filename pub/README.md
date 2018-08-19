# pub

Implements both the SocialAPI and FederateAPI in the ActivityPub specification.

## Disclaimer

This library is designed with flexibility in mind. The cost of doing so is that
writing an ActivityPub application requires a lot of careful considerations that
are not trivial. ActivityPub is an Application transport layer that is also tied
to a specific data model, making retrofits nontrivial as well.

## How To Use

There are two ActivityPub APIs: the SocialAPI between a user and your
ActivityPub server, and the FederateAPI between your ActivityPub server and
another server peer. This library lets you choose one or both.

*Lightning intro to ActivityPub: ActivityPub uses ActivityStreams as data. This
lives in `go-fed/activity/vocab`. ActivityPub has a concept of `actors` who can
send, receive, and read their messages. When sending and receiving messages from
a client (such as on their phone) to an ActivityPub server, it is via the
SocialAPI. When it is between two ActivityPub servers, it is via the
FederateAPI.*

Next, there are two kinds of ActivityPub requests to handle:

1. Requests that `GET` or `POST` to stuff owned by an `actor` like their `inbox`
 or `outbox`.
1. Requests that `GET` ActivityStream objects hosted on your server.

The first is the most complex, and requires the creation of a `Pubber`. It is
created depending on which APIs are to be supported:

```golang
// Only support SocialAPI
s := pub.NewSocialPubber(...)
// Only support FederateAPI
f := pub.NewFederatingPubber(...)
// Support both APIs
sf := pub.NewPubber(...)
```

Note that *only* the creation of the `Pubber` is affected by the decision of
which API to support. Once created, the `Pubber` should be used in the same
manner regardless of the API it is supporting. This allows your application
to easily adopt one API first and migrate to both later by simply changing how
the `Pubber` is created.

To use the `Pubber`, call its methods in the HTTP handlers responsible for an
`actor`'s `inbox` and `outbox`:

```golang
// Given:
//     var myPubber pub.Pubber
var outboxHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
  c := context.Background()
  // Populate c with application specific information
  if handled, err := myPubber.PostOutbox(c, w, r); err != nil {
    // Write to w
  } else if handled {
    return
  }
  if handled, err := myPubber.GetOutbox(c, w, r); err != nil {
    // Write to w
  } else if handled {
    return
  }
  // Handle non-ActivityPub request, such as responding with an HTML
  // representation with correct view permissions.
}
var inboxHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
  c := context.Background()
  // Populate c with application specific information
  if handled, err := myPubber.PostInbox(c, w, r); err != nil {
    // Write to w
  } else if handled {
    return
  }
  if handled, err := myPubber.GetInbox(c, w, r); err != nil {
    // Write to w
  } else if handled {
    return
  }
  // Handle non-ActivityPub request, such as responding with an HTML
  // representation with correct view permissions.
}
```

Finally, to handle the second kind of request, use the `HandlerFunc` within HTTP
handler functions in a similar way. There are two ways to create `HandlerFunc`,
which depend on decisions we will address later:

```golang
asHandler := pub.ServeActivityPubObject(...)
var activityStreamHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
  c := context.Background()
  // Populate c with application specific information
  if handled, err := asHandler(c, w, r); err != nil {
    // Write to w
  } else if handled {
    return
  }
  // Handle non-ActivityPub request, such as responding with an HTML
  // representation with correct view permissions.
}
```

That's all that's required to support ActivityPub.

## How To Create

You may have noticed that using the library is deceptively straightforward. This
is because *creating* the `Pubber` and `HandlerFunc` types is not trivial and
requires forethought.

There are a lot of interfaces that must be satisfied in order to have a complete
working ActivityPub server.

Note that `context.Context` is passed everywhere possible, to allow your
implementation to keep a request-specific context throughout the lifecycle of
an ActivityPub request.

### Application Interface

Regardless of which of the SocialAPI and FederateAPI chosen, the `Application`
interface contains the set of core methods fundamental to the functionality of
this library. It contains a lot of the storage fetching and writing, all of
which is keyed by `*url.URL`. To protect against race conditions, this library
will inform whether it is fetching data to read-only or fetching for read-or-
write.

Note that under some conditions, ActivityPub verifies the peer's request. It
does so using HTTP Signatures. However, this requires knowing the other party's
public key, and fetching this remotely is do-able. However, this library assumes
this server already has it locally; at this time it is up to implementations to
remotely fetch it if needed.

### SocialAPI and FederateAPI Interfaces

These interfaces capture additional behaviors required by the SocialAPI and the
FederateAPI.

The SocialAPI can additionally provide a mechanism for client authentication and
authorization using frameworks like Oauth 2.0. Such frameworks are not natively
supported in this library and must be supplied.

### Callbacker Interface

One of these is needed per ActivityPub API supported. For example, if both the
SocialAPI and FederateAPI are supported, then two of these are needed.

Upon receiving one of these activities from a `POST` to the inbox or outbox, the
correct callbacker will be called to handle either a SocialAPI activity or a
FederateAPI activity.

This is where the bulk of implementation-specific logic is expected to reside.

Do note that for some of these activities, default actions will already occur.
For example, if receiving an `Accept` in response to a sent `Follow`, this
library automatically handles adding the correct actor into the correct
`following` collection. This means a lot of the social and federate
functionality is provided out of the box.

### Deliverer Interface

This is an optional interface. Since this library needs to send HTTP requests,
it would be unwise for it to provide no way of allowing implementations to
rate limit, persist across downtime, back off, etc. This interface is satisfied
by the `go-fed/activity/deliverer` package which has an implementation that can
remember to send requests across downtime.

If an implementation does not care to have this level of control, a synchronous
implementation is very straightforward to make.

### Other Interfaces

Other interfaces such as `Typer` and `PubObject` are meant to limit modification
scope or require minimal ActivityStream compatibility to be used by this
library. As long as the `go-fed/activity/vocab` or `go-fed/activity/streams`
packages are being used, these interfaces will be natively supported.

## Other Considerations

Please see the README for `go-fed/activity` for the status of the latest
official implementation report.

The [go-fed.org](https://go-fed.org) website has a tutorial and documentation
for this library.
