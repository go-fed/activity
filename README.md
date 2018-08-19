# activity

`go get github.com/go-fed/activity`

This repository supports `vgo` and is remotely verifiable.

This repository contains three libraries for use in your golang applications:

* `vocab`: An ActivityStreams Vocabulary library
* `streams`: A convenience library for the ActivityStreams Vocabulary
* `pub`: ActivityPub SocialAPI (Client-to-Server) and FederateAPI
  (Server-to-Server)

This library is biased. It forgoes understanding JSON-LD in exchange for static
typing. It provides a large amount of default behavior to let Social,
Federated, or both kinds of ActivityPub applications just work.

## Status

**0.2.0** ([Semantic Versioning](https://semver.org/))

An [official implementation report](https://activitypub.rocks/implementation-report/)
was last submitted for version **0.2.0** [here](https://github.com/w3c/activitypub/issues/318).

Previous unofficial implementation reports are available in [issue #46](https://github.com/go-fed/activity/issues/46).

Please see CHANGELOG for changes between versions.

## Getting Started

See each subdirectory for its own README for further elaboration. The
recommended reading order is `vocab`, `streams`, and then `pub`. Others are
optional.

## How can I get help, file issues, or contribute?

Please see the CONTRIBUTING.md file!

## How well tested are these libraries?

I took great care to add numerous tests using examples directly from
specifications, official test repositories, and my own end-to-end tests.

## Who is using this library currently?

| Application | Description                                       | Repository                                                                 | Point Of Contact                                                                                                    | Homepage                             |
|:-----------:|:-------------------------------------------------:|:--------------------------------------------------------------------------:|:-------------------------------------------------------------------------------------------------------------------:|:------------------------------------:|
| Anancus     | Self-hosted and federated social link aggregation | [https://gitlab.com/tuxether/anancus](https://gitlab.com/tuxether/anancus) | [@tuxether@floss.social](https://floss.social/@tuxether) or [tuxether@protonmail.ch](mailto:tuxether@protonmail.ch) | N/A                                  |
| Write.as    | Simple, privacy-focused blogging platform         | [https://github.com/writeas/Write.as](https://github.com/writeas/Write.as) | [@write_as@writing.exchange](https://writing.exchange/@write_as) or [hello@write.as](mailto:hello@write.as)         | [https://write.as](https://write.as) |

## How do I use these libraries?

There is a [tutorial](https://go-fed.org/tutorial) available at
[https://go-fed.org](https://go-fed.org).

Please see each subdirectory for its own README for further elaboration. The
recommended reading order is `vocab`, `streams`, and then `pub`. Others are
optional.

Passing familiarity with ActivityStreams and ActivityPub is recommended.

## Other Libraries

* `tools` - Code generation wizardry and ActivityPub-spec-as-data.
* `deliverer` - Provides an asynchronous `Deliverer` for use with the `pub` lib

## FAQ

Please see the CONTRIBUTING.md file!

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
