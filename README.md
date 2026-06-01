<p align="center">
  <a href="https://runapi.ai"><img src="https://runapi.ai/icon.svg" height="56" alt="RunAPI"></a>
</p>

<h3 align="center">
  <a href="https://github.com/runapi-ai/gemini-omni-sdk">Gemini Omni API SDK for RunAPI</a>
</h3>

<p align="center">
  Gemini Omni API SDKs for JavaScript, Ruby, and Go on RunAPI, covering audio, character, and video workflows.
</p>

<div align="center">

[![npm](https://img.shields.io/npm/v/@runapi.ai/gemini-omni)](https://www.npmjs.com/package/@runapi.ai/gemini-omni)
[![RubyGems](https://img.shields.io/gem/v/runapi-gemini-omni)](https://rubygems.org/gems/runapi-gemini-omni)
[![Go Reference](https://pkg.go.dev/badge/github.com/runapi-ai/gemini-omni-sdk/go.svg)](https://pkg.go.dev/github.com/runapi-ai/gemini-omni-sdk/go)
[![License](https://img.shields.io/github/license/runapi-ai/gemini-omni-sdk)](https://github.com/runapi-ai/gemini-omni-sdk/blob/main/LICENSE)

</div>
<br/>

The gemini omni api SDK packages JavaScript, Ruby, and Go clients for Gemini Omni on RunAPI. Use this gemini omni api SDK for voice resources, character resources, and multimodal video generation workflows that need typed installs, JSON request bodies, and consistent RunAPI errors across services.

Gemini Omni belongs to the Google catalog on RunAPI. The public model page is https://runapi.ai/models/gemini-omni. The public `gemini-omni-sdk` repository groups the JavaScript, Ruby, and Go packages for this model.

## Install

```bash
npm install @runapi.ai/gemini-omni
gem install runapi-gemini-omni
go get github.com/runapi-ai/gemini-omni-sdk/go@latest
```

## JavaScript quick start

```typescript
import { GeminiOmniClient } from '@runapi.ai/gemini-omni';

const client = new GeminiOmniClient();

const voice = await client.createAudio.run({
  audio_id: 'achernar',
  name: 'Acher Narrator',
  voice_description: 'A calm, clear voice',
  example_dialogue: 'Hello, I am achernar',
});

const character = await client.createCharacter.run({
  descriptions: 'A silver-haired cyberpunk guide',
  reference_image_url: 'https://file.runapi.ai/demo/character.png',
  audio_ids: [voice.id],
  character_name: 'Jenny',
});

const video = await client.textToVideo.run({
  prompt: 'Create a neon city tracking shot with the character walking forward.',
  duration_seconds: 8,
  aspect_ratio: '16:9',
  output_resolution: '1080p',
  reference_image_urls: ['https://file.runapi.ai/demo/scene.png'],
  audio_ids: [voice.id],
  character_ids: [character.id],
});
```

## Public links

- Model page: https://runapi.ai/models/gemini-omni
- SDK docs: https://runapi.ai/docs#sdk-gemini-omni
- Product docs: https://runapi.ai/docs#gemini-omni
- SDK repository: https://github.com/runapi-ai/gemini-omni-sdk
- Skill repository: https://github.com/runapi-ai/gemini-omni
- Provider comparison: https://runapi.ai/providers/google
- Full catalog: https://runapi.ai/models

## License

Licensed under the Apache License, Version 2.0.
