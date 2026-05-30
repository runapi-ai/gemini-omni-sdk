# Gemini Omni API Go SDK for RunAPI

The gemini omni api Go SDK is the language-specific package for Gemini Omni on RunAPI. Use this package for voice resources, character resources, and multimodal video workflows that need JSON request bodies, task status lookup, and consistent RunAPI errors in Go.

## Install

```bash
go get github.com/runapi-ai/gemini-omni-sdk/go@latest
```

## Quick start

```go
import (
  "context"

  "github.com/runapi-ai/gemini-omni-sdk/go/geminiomni"
)

client, err := geminiomni.NewClient()
voice, err := client.CreateAudio.Run(context.Background(), geminiomni.CreateAudioParams{
  AudioID: "achernar",
  Name:    "Acher Narrator",
})

character, err := client.CreateCharacter.Run(context.Background(), geminiomni.CreateCharacterParams{
  Descriptions: "A silver-haired cyberpunk guide",
  ImageURLs:    []string{"https://file.runapi.ai/demo/character.png"},
})

video, err := client.TextToVideo.Run(context.Background(), geminiomni.TextToVideoParams{
  Prompt:       "Create a neon city tracking shot with the character walking forward.",
  Duration:     "8",
  Resolution:   "1080p",
  CharacterIDs: []string{character.ID},
})
```

## Links

- Model page: https://runapi.ai/models/gemini-omni
- SDK docs: https://runapi.ai/docs#sdk-gemini-omni
- Product docs: https://runapi.ai/docs#gemini-omni
- Pricing and rate limits: https://runapi.ai/models/gemini-omni
- Provider comparison: https://runapi.ai/providers/google
- Full catalog: https://runapi.ai/models

## License

Licensed under the Apache License, Version 2.0.
