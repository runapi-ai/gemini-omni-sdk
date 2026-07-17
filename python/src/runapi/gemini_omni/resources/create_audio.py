"""Gemini Omni create-audio resource (synchronous)."""

from __future__ import annotations

from typing import Any, Dict, Optional

from runapi.core import Resource, ValidationError, RequestOptions

from ..contract_gen import CONTRACT
from ..types import CreateAudioResponse


class CreateAudio(Resource):
    """Create a reusable voice. Synchronous: ``run()`` returns the result directly."""

    ENDPOINT = "/api/v1/gemini_omni/create_audio"

    RESPONSE_CLASS = CreateAudioResponse

    MODEL = "gemini-omni-audio"

    NAME_MAX_LENGTH = 210
    VOICE_DESCRIPTION_MAX_LENGTH = 20_000
    EXAMPLE_DIALOGUE_MAX_LENGTH = 120

    def run(self, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        """Create a reusable voice (synchronous).

        Args:
            **params: Reusable voice parameters (audio_id, name, ...).

        Returns:
            The result.
        """
        compacted = self._compact_params(params)
        self._validate_params(compacted)
        return self._request("post", self.ENDPOINT, body=compacted, options=options)

    def _validate_params(self, params: Dict[str, Any]) -> None:
        self._validate_contract(CONTRACT["create-audio"], {**params, "model": self.MODEL})
        self._validate_length(params, "name", self.NAME_MAX_LENGTH)
        self._validate_length(params, "voice_description", self.VOICE_DESCRIPTION_MAX_LENGTH)
        self._validate_length(params, "example_dialogue", self.EXAMPLE_DIALOGUE_MAX_LENGTH)

    @staticmethod
    def _validate_length(params: Dict[str, Any], key: str, max_length: int) -> None:
        value = params.get(key)
        if value is None or len(str(value)) <= max_length:
            return
        raise ValidationError(f"{key} must be at most {max_length} characters")
