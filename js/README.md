# Gemini Omni JavaScript SDK for RunAPI

The Gemini Omni JavaScript SDK is the language-specific package for Gemini Omni on RunAPI. Use this package for voice resources, character resources, and multimodal video generation workflows when your application needs request bodies, task status lookup, and consistent RunAPI errors in JavaScript.

## Install

```bash
npm install @runapi.ai/gemini-omni
```

## Quick start

```typescript
import { GeminiOmniClient } from '@runapi.ai/gemini-omni';

const client = new GeminiOmniClient();
const voice = await client.createAudio.run({
  audio_id: 'achernar',
  name: 'Acher Narrator',
});

const character = await client.createCharacter.run({
  descriptions: 'A silver-haired cyberpunk guide',
  reference_image_url: 'https://cdn.runapi.ai/public/samples/reference-1.jpg',
});

const video = await client.textToVideo.run({
  model: 'gemini-omni-flash-preview',
  prompt: 'A paper airplane glides through a sunlit studio.',
  aspect_ratio: '16:9',
  output_resolution: '720p',
});
```

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Links

- Model page: https://runapi.ai/models/gemini-omni
- SDK docs: https://runapi.ai/docs#sdk-gemini-omni
- Product docs: https://runapi.ai/docs#gemini-omni
- Flash Preview pricing and rate limits: https://runapi.ai/models/gemini-omni/flash-preview
- Provider comparison: https://runapi.ai/providers/google
- Full catalog: https://runapi.ai/models

## License

Licensed under the Apache License, Version 2.0.
