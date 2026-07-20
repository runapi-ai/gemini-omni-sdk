# Gemini Omni Python SDK for RunAPI

The Gemini Omni Python SDK is the language-specific package for Gemini Omni on RunAPI. Use this package for voice resources, character resources, and multimodal video generation workflows when your application needs request bodies, task status lookup, and consistent RunAPI errors in Python.

## Install

```bash
pip install runapi-gemini-omni
```

## Quick start

```python
from runapi.gemini_omni import GeminiOmniClient

client = GeminiOmniClient()  # reads RUNAPI_API_KEY, or pass api_key="sk-..."

# Reusable voice (returns immediately)
voice = client.create_audio.run(audio_id="kore", name="Narrator")

# Reusable character (returns immediately)
character = client.create_character.run(
    descriptions="A friendly robot guide",
    reference_image_url="https://cdn.runapi.ai/public/samples/reference-1.jpg",
)

# Text-to-video (create + poll until complete)
result = client.text_to_video.run(
    model="gemini-omni-flash-preview",
    prompt="A paper airplane glides through a sunlit studio.",
    aspect_ratio="16:9",
    output_resolution="720p",
)
print(result.videos[0].url)
```

`create_audio` and `create_character` are synchronous: `run` submits and returns the result directly. `text_to_video` is asynchronous: use `create` to submit and return quickly, `get` to fetch the latest task state, and `run` to create and poll until completion.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Language notes

Pass parameters as keyword arguments and catch the `runapi.gemini_omni` error classes. The available resources are `create_audio`, `create_character`, and `text_to_video`. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

## Links

- Model page: https://runapi.ai/models/gemini-omni
- SDK docs: https://runapi.ai/docs#sdk-gemini-omni
- Product docs: https://runapi.ai/docs#gemini-omni
- Flash Preview pricing and rate limits: https://runapi.ai/models/gemini-omni/flash-preview
- Provider comparison: https://runapi.ai/providers/google
- Full catalog: https://runapi.ai/models

## License

Licensed under the Apache License, Version 2.0.
