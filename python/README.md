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
    reference_image_url="https://cdn.runapi.ai/public/samples/portrait.jpg",
    body_reference_image_url="https://cdn.runapi.ai/public/samples/image.jpg",
)

# Text-to-video (create + poll until complete)
result = client.text_to_video.run(
    model="gemini-omni-flash-1-1",
    prompt="A paper airplane travels from dawn into dusk.",
    duration_seconds=6,
    first_frame_image_url="https://cdn.runapi.ai/public/samples/first-frame.jpg",
    last_frame_image_url="https://cdn.runapi.ai/public/samples/last-frame.jpg",
    aspect_ratio="16:9",
    output_resolution="360p",
)
print(result.videos[0].url)
```

`character.images` returns the portrait first and the optional full-body image second. A character created with both references consumes two of the seven reference units in a multimodal video request.

For `gemini-omni-flash-1-1`, `first_frame_image_url` cannot be combined with reference images, audio IDs, video clips, or character IDs. `last_frame_image_url` requires `first_frame_image_url`.

`create_audio` returns its result directly. `create_character.run` returns the
terminal character result and follows an accepted Task when needed.
`text_to_video` is asynchronous: use `create` to submit and return quickly,
`get` to fetch the latest task state, and `run` to create and poll until
completion.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Language notes

Pass parameters as keyword arguments and catch the `runapi.gemini_omni` error classes. The available resources are `create_audio`, `create_character`, and `text_to_video`. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

## Links

- Model page: https://runapi.ai/models/gemini-omni
- SDK docs: https://runapi.ai/docs/resources/sdks
- Product docs: https://runapi.ai/docs/api/gemini-omni/text-to-video
- Flash 1.1 pricing and rate limits: https://runapi.ai/models/gemini-omni/flash-1-1
- Flash Preview pricing and rate limits: https://runapi.ai/models/gemini-omni/flash-preview
- Full catalog: https://runapi.ai/models

## License

Licensed under the Apache License, Version 2.0.
