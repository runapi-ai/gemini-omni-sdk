---
name: gemini-omni
description: Create Gemini Omni voice resources, character resources, and text-to-video tasks through RunAPI. Use when the user asks an agent to create or manage Gemini Omni audio voices, character resources, or multimodal video. Default to the RunAPI CLI for one-off calls; use SDKs only when integrating RunAPI into an app or backend.
documentation: https://runapi.ai/models/gemini-omni.md
provider_page: https://runapi.ai/providers/google.md
catalog: https://runapi.ai/models.md
metadata:
  openclaw:
    homepage: https://runapi.ai/models/gemini-omni
    requires:
      bins:
      - runapi
    install:
    - kind: brew
      formula: runapi-ai/tap/runapi
      bins:
      - runapi
    envVars:
    - name: RUNAPI_API_KEY
      required: false
      description: Optional RunAPI API key; agents should prefer environment auth or saved CLI config.
---

# Gemini Omni On RunAPI

Create Gemini Omni voice resources, character resources, and text-to-video tasks through RunAPI. The default path for one-off agent tasks is the `runapi` CLI; SDKs are for application integration.

## Routing decision

- One-off voice, character, or video generation for the user → use the **CLI path** with the `runapi` binary.
- Building an app, backend, worker, library, or production codebase → use the **SDK integration path**.

## Variants

- Audio: create reusable voice resources with `runapi gemini-omni create-audio`.
- Character: create reusable character resources with `runapi gemini-omni create-character`.
- Video: create text-to-video tasks with `runapi gemini-omni text-to-video`.

## CLI path

The `runapi` binary is the runtime dependency. Run `runapi auth status` first. For agents and headless runs, prefer `RUNAPI_API_KEY` or import it into saved config.

Inspect the available commands and request fields with CLI help:

```shell
runapi gemini-omni --help
runapi gemini-omni create-audio --help
runapi gemini-omni create-character --help
runapi gemini-omni text-to-video --help
```

Run a one-off request:

```shell
runapi gemini-omni create-audio --input-file request.json
```

For video tasks, submit asynchronously and wait for completion:

```shell
runapi gemini-omni text-to-video --async --input-file request.json
runapi wait <task-id> --service gemini-omni --action text-to-video
```

Available commands: `create-audio`, `create-character`, `text-to-video`.

## SDK integration path

When integrating Gemini Omni into an app, backend, worker, or library — not for one-off tasks — use a RunAPI SDK package:

- JavaScript / TypeScript: `@runapi.ai/gemini-omni`
- Ruby: `runapi-gemini-omni`
- Go: `github.com/runapi-ai/gemini-omni-sdk/go`

## References

- Model overview, pricing, and rate limits: https://runapi.ai/models/gemini-omni.md
- Provider comparison: https://runapi.ai/providers/google.md
- Full model catalog: https://runapi.ai/models.md
