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
video, err := client.TextToVideo.Run(context.Background(), geminiomni.TextToVideoParams{
  Model:            geminiomni.ModelGeminiOmniFlashPreview,
  Prompt:           "A paper airplane glides through a sunlit studio.",
  AspectRatio:      "16:9",
  OutputResolution: "720p",
})
if err != nil {
  panic(err)
}
fmt.Println(video.Videos[0].URL)
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
