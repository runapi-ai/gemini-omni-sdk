# Gemini Omni Ruby SDK for RunAPI

The Gemini Omni Ruby SDK is the language-specific package for Gemini Omni on RunAPI. Use this package for voice resources, character resources, and multimodal video generation workflows when your application needs request bodies, task status lookup, and consistent RunAPI errors in Ruby.

## Install

```bash
gem install runapi-gemini-omni
```

## Quick start

```ruby
require "runapi/gemini_omni"

client = RunApi::GeminiOmni::Client.new
voice = client.create_audio.run(
  audio_id: "achernar",
  name: "Acher Narrator"
)

character = client.create_character.run(
  descriptions: "A silver-haired cyberpunk guide",
  reference_image_url: "https://cdn.runapi.ai/public/samples/reference-1.jpg"
)

video = client.text_to_video.run(
  model: "gemini-omni-flash-preview",
  prompt: "A paper airplane glides through a sunlit studio.",
  aspect_ratio: "16:9",
  output_resolution: "720p"
)
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
