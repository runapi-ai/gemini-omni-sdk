# Gemini Omni API Ruby SDK for RunAPI

The gemini omni api Ruby SDK is the language-specific package for Gemini Omni on RunAPI. Use this package for voice resources, character resources, and multimodal video workflows that need JSON request bodies, task status lookup, and consistent RunAPI errors in Ruby.

## Install

```bash
gem install runapi-gemini-omni
```

## Quick start

```ruby
require "runapi-gemini_omni"

client = RunApi::GeminiOmni::Client.new
voice = client.create_audio.run(
  audio_id: "achernar",
  name: "Acher Narrator"
)

character = client.create_character.run(
  descriptions: "A silver-haired cyberpunk guide",
  image_urls: ["https://file.runapi.ai/demo/character.png"]
)

video = client.text_to_video.run(
  prompt: "Create a neon city tracking shot with the character walking forward.",
  duration: "8",
  resolution: "1080p",
  character_ids: [character.id]
)
```

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Links

- Model page: https://runapi.ai/models/gemini-omni
- SDK docs: https://runapi.ai/docs#sdk-gemini-omni
- Product docs: https://runapi.ai/docs#gemini-omni
- Pricing and rate limits: https://runapi.ai/models/gemini-omni
- Provider comparison: https://runapi.ai/providers/google
- Full catalog: https://runapi.ai/models

## License

Licensed under the Apache License, Version 2.0.
