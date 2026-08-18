# Changelog

## [ruby/v0.3.4](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/ruby%2Fv0.3.4) - 2026-08-18

### Changed
- Allow Ruby clients to install the core SDK release that adds persistent Files and multipart Uploads alongside this model SDK.


## [js/v0.3.3](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/js%2Fv0.3.3), [ruby/v0.3.3](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/ruby%2Fv0.3.3), [go/v0.3.3](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/go%2Fv0.3.3), [python/v0.3.1](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/python%2Fv0.3.1), [java/v0.2.2](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/java%2Fv0.2.2) - 2026-07-28

### Changed
- Describe verified mixed-media fields, required inputs, and documented collection and length limits, including video trim validation precedence.


## [java/v0.2.1](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/java%2Fv0.2.1) - 2026-07-28

### Added
- Decode typed Task Billing Facts on synchronous audio and character responses.

## [go/v0.3.2](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/go%2Fv0.3.2) - 2026-07-28

### Added
- Expose persisted billing facts on task responses.

## [js/v0.3.2](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/js%2Fv0.3.2), [ruby/v0.3.2](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/ruby%2Fv0.3.2) - 2026-07-28

### Added
- Type task billing facts on task-backed responses.


## [python/v0.3.0](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/python%2Fv0.3.0) - 2026-07-24

### Added
- Expose shared Files, Account, and Pricing resources plus typed Task Billing Facts through the Provider Client.


## [js/v0.3.1](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/js%2Fv0.3.1), [ruby/v0.3.1](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/ruby%2Fv0.3.1), [go/v0.3.1](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/go%2Fv0.3.1), [python/v0.2.1](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/python%2Fv0.2.1) - 2026-07-20

### Changed
- Validate generated limits for reference images, audio IDs, video items, and character IDs.


## [js/v0.3.0](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/js%2Fv0.3.0), [ruby/v0.3.0](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/ruby%2Fv0.3.0), [go/v0.3.0](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/go%2Fv0.3.0), [python/v0.2.0](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/python%2Fv0.2.0), [java/v0.2.0](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/java%2Fv0.2.0) - 2026-07-20

### Added
- Add prompt-only Gemini Omni Flash Preview text-to-video requests with model-specific validation.


## [js/v0.2.8](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/js%2Fv0.2.8), [go/v0.2.8](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/go%2Fv0.2.8) - 2026-07-17

### Fixed
- Align JavaScript and Go with the current RunAPI core parameter validation APIs.

## [java/v0.1.1](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/java%2Fv0.1.1) - 2026-06-25

### Fixed
- Fixed Java retry handling for Retry-After response headers.
- Fixed Java contract validation for action-level conditional rules.
- Refreshed Java SDK metadata for v0.1.1.

## [java/v0.1.0](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/java%2Fv0.1.0) - 2026-06-24

### Added
- Publish `ai.runapi:runapi-gemini-omni` for Java SDK consumers.
- Include typed Java builders, synchronous client resources, sources, and Javadocs.

## [js/v0.2.7](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/js%2Fv0.2.7), [ruby/v0.2.7](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/ruby%2Fv0.2.7), [go/v0.2.7](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/go%2Fv0.2.7), [python/v0.1.0](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/python%2Fv0.1.0) - 2026-06-18

### Changed
- Per-method documentation for all resource methods

## [js/v0.2.6](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/js%2Fv0.2.6), [ruby/v0.2.6](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/ruby%2Fv0.2.6), [go/v0.2.6](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/go%2Fv0.2.6) - 2026-06-01

### Changed
- Align SDK with upstream Input Contract and public API vocabulary changes
- Update endpoint definitions and field constraints

## [js/v0.2.5](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/js%2Fv0.2.5), [ruby/v0.2.5](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/ruby%2Fv0.2.5), [go/v0.2.5](https://github.com/runapi-ai/gemini-omni-sdk/releases/tag/go%2Fv0.2.5) - 2026-05-30

### Added
- Publish the initial Gemini Omni SDK for JavaScript, Ruby, and Go.
- Include language README files, package metadata, and RunAPI catalog links.
