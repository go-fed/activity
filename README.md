# activity

`go get github.com/go-fed/activity`

This repository contains three libraries for use in your golang applications:

* An ActivityStreams Vocabulary library
* A convenience library for the ActivityStreams Vocabulary
* ActivityPub Social API (Client-to-Server) and Federation Protocol
  (Server-to-Server).

This library is biased. It forgoes understanding JSON-LD in exchange for static
typing. It provides a large amount of default behavior to let Social,
Federated, or both kinds of ActivityPub applications just work.

## Status

There is no stable version of this library (yet).

See each subdirectory for its own README for further elaboration.

### Core ActivityPub Libraries

* `vocab` - ActivityStreams Vocabulary: Functional and tested
* `streams` - ActivityStreams Convenience Library: Functional and tested
* `pub` - ActivityPub: Under development and testing

### Supplemental Libraries

* `tools` - Code generation wizardry and ActivityPub-spec-as-data.
* `deliverer` - Provides an asynchronous `Deliverer` for use with the `pub` lib

## How To Use This Library

This section will be fleshed out once the library is approaching its first
stable release.

## FAQ

*Why does compilation take so long?*

The `vocab` and `streams` packages are code generated on order of hundreds of
thousands to a million lines long. Use `go install` or `go build -i` to cache
the build artifacts and do incremental builds.

## Useful References

* [ActivityPub Specification](https://www.w3.org/TR/activitypub)
* [ActivityPub GitHub Repo](https://github.com/w3c/activitypub)
* [ActivityStreams Core Specification](https://www.w3.org/TR/activitystreams-core)
* [ActivityStreams Vocabulary Specification](https://www.w3.org/TR/activitystreams-vocabulary)
* [ActivityStreams GitHub Repo](https://github.com/w3c/activitystreams)

## Thanks

I would like to thank those that have worked hard to create the technologies
and standards that created the opportunity to implement this suite of
libraries.
