"""Gemini Omni create-audio resource (synchronous)."""

from __future__ import annotations

from typing import Any, Optional

from runapi.core import Resource, RequestOptions

from ..types import CreateAudioResponse


class CreateAudio(Resource):
    """Create a reusable voice. Synchronous: ``run()`` returns the result directly."""

    ENDPOINT = "/api/v1/gemini_omni/create_audio"

    RESPONSE_CLASS = CreateAudioResponse

    MODEL = "gemini-omni-audio"

    def run(self, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        """Create a reusable voice (synchronous).

        Args:
            **params: Reusable voice parameters (audio_id, name, ...).

        Returns:
            The result.
        """
        compacted = self._compact_params(params)
        return self._request("post", self.ENDPOINT, body=compacted, options=options)
