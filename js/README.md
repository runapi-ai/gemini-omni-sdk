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
  reference_image_url: 'https://cdn.runapi.ai/public/samples/portrait.jpg',
  body_reference_image_url: 'https://cdn.runapi.ai/public/samples/image.jpg',
});

const video = await client.textToVideo.run({
  model: 'gemini-omni-flash-1-1',
  prompt: 'A paper airplane travels from dawn into dusk.',
  duration_seconds: 6,
  first_frame_image_url: 'https://cdn.runapi.ai/public/samples/first-frame.jpg',
  last_frame_image_url: 'https://cdn.runapi.ai/public/samples/last-frame.jpg',
  aspect_ratio: '16:9',
  output_resolution: '360p',
});
```

`character.images` returns the portrait first and the optional full-body image second. A character created with both references consumes two of the seven reference units in a multimodal video request.

For `gemini-omni-flash-1-1`, `first_frame_image_url` cannot be combined with reference images, audio IDs, video clips, or character IDs. `last_frame_image_url` requires `first_frame_image_url`.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Links

- Model page: https://runapi.ai/models/gemini-omni
- SDK docs: https://runapi.ai/docs/resources/sdks
- Product docs: https://runapi.ai/docs/api/gemini-omni/text-to-video
- Flash 1.1 pricing and rate limits: https://runapi.ai/models/gemini-omni/flash-1-1
- Flash Preview pricing and rate limits: https://runapi.ai/models/gemini-omni/flash-preview
- Full catalog: https://runapi.ai/models

## License

Licensed under the Apache License, Version 2.0.
