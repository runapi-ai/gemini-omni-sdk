<p align="center">
  <a href="https://github.com/runapi-ai/gemini-omni">
    <h3 align="center">Gemini Omni API Skill for RunAPI</h3>
  </a>
</p>

<p align="center">
  Install this agent skill, inspect Gemini Omni fields, then create voice, character, and video resources through the RunAPI CLI.
</p>

<p align="center">
  <a href="https://runapi.ai/models/gemini-omni"><strong>Model Reference</strong></a> · <a href="https://github.com/runapi-ai/cli"><strong>CLI</strong></a> · <a href="https://github.com/runapi-ai/gemini-omni-sdk"><strong>SDK</strong></a>
</p>

<div align="center">

[![skills.sh](https://www.skills.sh/b/runapi-ai/gemini-omni)](https://www.skills.sh/runapi-ai/gemini-omni/gemini-omni)
[![ClawHub](https://img.shields.io/badge/ClawHub-runapi--gemini--omni-111827)](https://clawhub.ai/runapi-ai/runapi-gemini-omni)
[![License](https://img.shields.io/github/license/runapi-ai/gemini-omni)](https://github.com/runapi-ai/gemini-omni/blob/main/LICENSE)

</div>
<br/>

Create Gemini Omni voice resources, character resources, and text-to-video tasks through RunAPI. This skill helps Claude Code, Codex, Gemini CLI, Cursor, and 50+ agents use Gemini Omni through RunAPI.

The canonical agent file is `skills/gemini-omni/SKILL.md`.

## Variants

- Audio: create reusable voice resources with `runapi gemini-omni create-audio`.
- Character: create reusable character resources with `runapi gemini-omni create-character`.
- Video: create text-to-video tasks with `runapi gemini-omni text-to-video`.

## Install

```bash
npx skills add runapi-ai/gemini-omni -g
```

Or paste this prompt to your AI agent:

```text
Install the gemini-omni skill for me:

1. Clone https://github.com/runapi-ai/gemini-omni
2. Copy the skills/gemini-omni/ directory into your
   user-level skills directory.
3. Verify that SKILL.md is present.
4. Confirm the install path when done.
```

## Quick example

```shell
runapi gemini-omni create-audio --input-file request.json
runapi gemini-omni create-character --input-file request.json
runapi gemini-omni text-to-video --async --input-file request.json
runapi wait <task-id> --service gemini-omni --action text-to-video
```

## Routing

- Model page: https://runapi.ai/models/gemini-omni
- Product docs: https://runapi.ai/docs#gemini-omni
- SDK docs: https://runapi.ai/docs#sdk-gemini-omni
- SDK repository: https://github.com/runapi-ai/gemini-omni-sdk
- Pricing and rate limits: https://runapi.ai/models/gemini-omni
- Provider comparison: https://runapi.ai/providers/google
- Browse all RunAPI models and skills: https://runapi.ai/models

## Agent rules

- Integration work uses the target language SDK; one-off generation, manual smoke tests, debugging, or user-requested CLI runs use the RunAPI CLI skill: https://github.com/runapi-ai/cli-skill
- RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## License

Licensed under the Apache License, Version 2.0.
