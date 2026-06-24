"""Gemini Omni enums and response models."""

from __future__ import annotations

from runapi.core import BaseModel, TaskResponse, optional, required

AUDIO_VOICES = [
    "achernar", "achird", "algenib", "algieba", "alnilam", "aoede", "autonoe",
    "callirrhoe", "charon", "despina", "enceladus", "erinome", "fenrir", "gacrux",
    "iapetus", "kore", "laomedeia", "leda", "orus", "puck", "pulcherrima",
    "rasalgethi", "sadachbia", "sadaltager", "schedar", "sulafat", "umbriel",
    "vindemiatrix", "zephyr", "zubenelgenubi",
]

SEED_MIN = 0
SEED_MAX = 2_147_483_647


class Audio(BaseModel):
    id = required(str)
    name = optional(str)


class CreateAudioResponse(BaseModel):
    """Result of creating a Gemini Omni reusable voice."""
    id = required(str)
    audio = optional(lambda: Audio)
    error = optional(str)


class Image(BaseModel):
    url = optional(str)


class Character(BaseModel):
    id = required(str)
    name = optional(str)
    images = optional([lambda: Image])


class CreateCharacterResponse(BaseModel):
    """Result of creating a Gemini Omni reusable character."""
    id = required(str)
    character = optional(lambda: Character)
    error = optional(str)


class Video(BaseModel):
    url = optional(str)


class TextToVideoResponse(TaskResponse):
    """Task status/result for Gemini Omni text-to-video."""
    videos = optional([lambda: Video])


class CompletedTextToVideoResponse(TextToVideoResponse):
    """Narrowed Gemini Omni text-to-video response once polling observes completion."""
    videos = required([lambda: Video])
