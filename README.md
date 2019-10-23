# activity

`go get github.com/go-fed/activity`

This repository contains two libraries and a tool:

* `astool`: A linked-data aware tool to generate golang native types for any
ActivityStreams vocabulary.
* `streams`: The ActivityStreams native types generated with the `astool`.
* `pub`: ActivityPub Social Protocol (Client-to-Server or C2S) and Federating
Protocol (Server-to-Server or S2S)

## Status

**Preview (unstable) 1.0.0** ([Semantic Versioning](https://semver.org/))

An [official implementation report](https://activitypub.rocks/implementation-report/)
was last submitted for version **0.2.0** [here](https://github.com/w3c/activitypub/issues/318).

Previous unofficial implementation reports are available in [issue #46](https://github.com/go-fed/activity/issues/46).

Please see CHANGELOG for changes between versions.

## Getting Started

See `astool`, `streams`, or `pub` for their own README.

## How can I get help, file issues, or contribute?

Please see the CONTRIBUTING.md file!

## How well tested are these libraries?

I took great care to add numerous tests using examples directly from
specifications, official test repositories, and my own end-to-end tests.

**v1.0.0** still has a lot of unit tests to be written. Bug fixes and small
backwards-incompatible behavior is expected before it is blessed as being
released.

## Who is using this library currently?

| Application | Description                                       | Repository                                                                 | Point Of Contact                                                                                                    | Homepage                             |
|:-----------:|:-------------------------------------------------:|:--------------------------------------------------------------------------:|:-------------------------------------------------------------------------------------------------------------------:|:------------------------------------:|
| Anancus       | Self-hosted and federated social link aggregation              | [https://gitlab.com/tuxether/anancus](https://gitlab.com/tuxether/anancus)       | [@tuxether@floss.social](https://floss.social/@tuxether) or [tuxether@protonmail.ch](mailto:tuxether@protonmail.ch) | N/A                                                |
| WriteFreely   | Simple, open-source, privacy-focused blogging platform         | [https://github.com/writeas/writefreely](https://github.com/writeas/writefreely) | [@write_as@writing.exchange](https://writing.exchange/@write_as) or [hello@write.as](mailto:hello@write.as)         | [https://writefreely.org](https://writefreely.org) |
| Read.as       | Long-form reader built on open protocols                       | [https://github.com/writeas/Read.as](https://github.com/writeas/Read.as)         | [@write_as@writing.exchange](https://writing.exchange/@write_as) or [hello@write.as](mailto:hello@write.as)         | [https://read.as](https://read.as)                 |
| go-fed/apcore | Generic ActivityPub server framework in Go                     | [https://github.com/go-fed/apcore](https://github.com/go-fed/apcore)             | [@cj@mastodon.technology](https://mastodon.technology/@cj) or [cjslep@gmail.com](mailto:cjslep@gmail.com)           | [https://go-fed.org](https://go-fed.org)           |

## How do I use these libraries?

Please see each subdirectory for its own README for further elaboration.

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

Thanks to those who have been early adopters with v0 and/or provided early
feedback.
