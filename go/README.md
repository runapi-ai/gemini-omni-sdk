# Gemini Omni Go SDK for RunAPI

The Gemini Omni Go SDK is the language-specific package for Gemini Omni on RunAPI. Use this package for voice resources, character resources, and multimodal video generation workflows when your application needs request bodies, task status lookup, and consistent RunAPI errors in Go.

## Install

```bash
go get github.com/runapi-ai/gemini-omni-sdk/go@latest
```

## Quick start

```go
import (
  "context"
  "fmt"

  "github.com/runapi-ai/gemini-omni-sdk/go/geminiomni"
)

client, err := geminiomni.NewClient()
character, err := client.CreateCharacter.Run(context.Background(), geminiomni.CreateCharacterParams{
  Descriptions:          "A silver-haired cyberpunk guide",
  ReferenceImageURL:     "https://cdn.runapi.ai/public/samples/portrait.jpg",
  BodyReferenceImageURL: "https://cdn.runapi.ai/public/samples/image.jpg",
})
if err != nil {
  panic(err)
}

video, err := client.TextToVideo.Run(context.Background(), geminiomni.TextToVideoParams{
  Model:              geminiomni.ModelGeminiOmniFlash11,
  Prompt:             "A paper airplane travels from dawn into dusk.",
  DurationSeconds:    6,
  FirstFrameImageURL: "https://cdn.runapi.ai/public/samples/first-frame.jpg",
  LastFrameImageURL:  "https://cdn.runapi.ai/public/samples/last-frame.jpg",
  AspectRatio:        "16:9",
  OutputResolution:   "360p",
})
if err != nil {
  panic(err)
}
fmt.Println(character.Character.Images[0].URL, character.Character.Images[1].URL)
fmt.Println(video.Videos[0].URL)
```

Character images are ordered portrait first and optional full-body image second. A character created with both references consumes two of the seven reference units in a multimodal video request.

For `gemini-omni-flash-1-1`, `FirstFrameImageURL` cannot be combined with reference images, audio IDs, video clips, or character IDs. `LastFrameImageURL` requires `FirstFrameImageURL`.

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
