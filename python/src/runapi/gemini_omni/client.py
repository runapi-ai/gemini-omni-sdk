"""Gemini Omni client."""

from __future__ import annotations

from typing import Any, Optional

from runapi.core import ProviderClient

from .resources.create_audio import CreateAudio
from .resources.create_character import CreateCharacter
from .resources.text_to_video import TextToVideo


class GeminiOmniClient(ProviderClient):
    """Gemini Omni client: reusable voices, characters, and text-to-video.

    Example::

        client = GeminiOmniClient(api_key="sk-...")
        result = client.text_to_video.run(prompt="A fox in snow", duration_seconds=8)
    """

    def __init__(self, api_key: Optional[str] = None, **options: Any) -> None:
        super().__init__(api_key, **options)
        http = self._http
        self.create_audio = CreateAudio(http)
        self.create_character = CreateCharacter(http)
        self.text_to_video = TextToVideo(http)
