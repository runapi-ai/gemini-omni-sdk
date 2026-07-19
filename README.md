<p align="center">
  <a href="https://runapi.ai"><img src="https://runapi.ai/icon.svg" height="56" alt="RunAPI"></a>
</p>

<h3 align="center">
  <a href="https://github.com/runapi-ai/gemini-omni-sdk">Gemini Omni API SDK for RunAPI</a>
</h3>

<p align="center">
  Gemini Omni API SDKs for JavaScript, Python, Ruby, Go, Java, and PHP on RunAPI, covering audio, character, and video workflows.
</p>

<div align="center">

[![npm](https://img.shields.io/npm/v/@runapi.ai/gemini-omni)](https://www.npmjs.com/package/@runapi.ai/gemini-omni)
[![PyPI](https://img.shields.io/pypi/v/runapi-gemini-omni)](https://pypi.org/project/runapi-gemini-omni/)
[![RubyGems](https://img.shields.io/gem/v/runapi-gemini-omni)](https://rubygems.org/gems/runapi-gemini-omni)
[![Go Reference](https://pkg.go.dev/badge/github.com/runapi-ai/gemini-omni-sdk/go.svg)](https://pkg.go.dev/github.com/runapi-ai/gemini-omni-sdk/go)
[![Maven Central](https://img.shields.io/maven-central/v/ai.runapi/runapi-gemini-omni)](https://central.sonatype.com/artifact/ai.runapi/runapi-gemini-omni)
[![License](https://img.shields.io/github/license/runapi-ai/gemini-omni-sdk)](https://github.com/runapi-ai/gemini-omni-sdk/blob/main/LICENSE)

</div>
<br/>

The Gemini Omni API SDK packages JavaScript, Python, Ruby, Go, Java, and PHP clients for Gemini Omni on RunAPI. Use it for audio, character, and video workflows when your app needs typed request builders, predictable task polling, file upload helpers, account helpers, and consistent RunAPI errors.

Gemini Omni is listed in the RunAPI model catalog at https://runapi.ai/models/gemini-omni. Variant pages below carry pricing, rate-limit, and commercial-usage details. The public `gemini-omni-sdk` repository groups the non-PHP language packages, examples, CI, and release tags for this model. The PHP package is released from a split Composer repository.

## Install

```bash
npm install @runapi.ai/gemini-omni
pip install runapi-gemini-omni
gem install runapi-gemini-omni
go get github.com/runapi-ai/gemini-omni-sdk/go@latest
```

Gradle:

```kotlin
dependencies {
  implementation("ai.runapi:runapi-gemini-omni:0.1.1")
}
```

Maven:

```xml
<dependency>
  <groupId>ai.runapi</groupId>
  <artifactId>runapi-gemini-omni</artifactId>
  <version>0.1.1</version>
</dependency>
```

Use the Java BOM when installing multiple RunAPI Java modules:

```kotlin
dependencies {
  implementation(platform("ai.runapi:runapi-bom:0.1.7"))
  implementation("ai.runapi:runapi-gemini-omni")
}
```

The PHP package is published from the split Composer repository as `runapi-ai/gemini-omni`; see https://github.com/runapi-ai/gemini-omni-php for PHP install and examples.

## What you can build

- Build apps, agent workflows, batch jobs, and production services around Gemini Omni requests.
- Install only the language package your app needs while keeping one model-specific repository for docs and releases.
- Use `create` for submit-only jobs, `get` for status lookup, and `run` for submit-and-poll scripts.
- Upload local files, URL files, or base64 files through shared RunAPI file helpers.
- Handle validation, authentication, rate limits, insufficient credits, task failures, and polling timeouts through RunAPI SDK errors.

## Java quick start

```java
import ai.runapi.geminiomni.GeminiOmniClient;
import ai.runapi.geminiomni.types.TextToVideoParams;
import ai.runapi.geminiomni.types.CompletedTextToVideoResponse;
import ai.runapi.geminiomni.types.TextToVideoModel;

GeminiOmniClient client = GeminiOmniClient.builder()
    .apiKey(System.getenv("RUNAPI_API_KEY"))
    .build();

CompletedTextToVideoResponse result = client.textToVideo().run(
    TextToVideoParams.builder()
        .model(TextToVideoModel.GEMINI_OMNI_TEXT_TO_VIDEO)
        .prompt("A tiny paper boat floating through a glowing cave")
        .durationSeconds(8)
        .aspectRatio("16:9")
        .outputResolution("720p")
        .build()
);
```

Java packages target Java 8 bytecode and are tested on Java 8, 11, 17, and 21. Each model artifact depends on `ai.runapi:runapi-core`, so application code normally installs only `ai.runapi:runapi-gemini-omni`.

## Task lifecycle

Most media endpoints are asynchronous. `create()` submits a task and returns its id, `get(id)` fetches the latest task state, and `run(params)` creates the task and polls until it reaches a terminal state. In web request handlers, prefer `create()` plus webhook or later `get()` polling so the server does not hold a worker open.

## Repository layout

- `js/` publishes `@runapi.ai/gemini-omni`.
- `python/` publishes `runapi-gemini-omni`.
- `ruby/` publishes `runapi-gemini-omni`.
- `go/` publishes `github.com/runapi-ai/gemini-omni-sdk/go` and depends on `github.com/runapi-ai/core-sdk/go`.
- `java/` publishes `ai.runapi:runapi-gemini-omni` and depends on `ai.runapi:runapi-core`.

## Public links

- Model page: https://runapi.ai/models/gemini-omni
- SDK docs: https://runapi.ai/docs#sdk-gemini-omni
- Product docs: https://runapi.ai/docs#gemini-omni
- SDK repository: https://github.com/runapi-ai/gemini-omni-sdk
- PHP package repository: https://github.com/runapi-ai/gemini-omni-php
- Skill repository: https://github.com/runapi-ai/gemini-omni
- Provider comparison: https://runapi.ai/providers/google
- Full catalog: https://runapi.ai/models

## Pricing and variants

Use the most specific Gemini Omni variant page for pricing, rate limits, and commercial usage:
- [Audio](https://runapi.ai/models/gemini-omni)
- [Character](https://runapi.ai/models/gemini-omni/character)
- [Text to video](https://runapi.ai/models/gemini-omni/text-to-video)

Default pricing link for the Gemini Omni SDK: https://runapi.ai/models/gemini-omni

## File storage

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## FAQ

### Which package should I install for Gemini Omni work?

Install the model package for your language: `@runapi.ai/gemini-omni` on npm, `runapi-gemini-omni` on PyPI, `runapi-gemini-omni` on RubyGems, `github.com/runapi-ai/gemini-omni-sdk/go`, `ai.runapi:runapi-gemini-omni` on Maven Central, or `runapi-ai/gemini-omni` on Packagist. Install core SDK packages only when you are building shared SDK infrastructure.

### Where should public links point?

Primary Gemini Omni links point to https://runapi.ai/models/gemini-omni. Pricing and usage-policy links point to variant pages such as https://runapi.ai/models/gemini-omni. Provider comparisons point to https://runapi.ai/providers/google, and broad browsing points to https://runapi.ai/models.

## License

Licensed under the Apache License, Version 2.0.
