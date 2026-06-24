"""Gemini Omni create-character resource (synchronous)."""

from __future__ import annotations

from typing import Any, Dict

from runapi.core import Resource, ValidationError

from ..contract_gen import CONTRACT
from ..types import CreateCharacterResponse


class CreateCharacter(Resource):
    """Create a reusable character. Synchronous: ``run()`` returns the result directly."""

    ENDPOINT = "/api/v1/gemini_omni/create_character"

    RESPONSE_CLASS = CreateCharacterResponse

    MODEL = "gemini-omni-character"

    DESCRIPTIONS_MAX_LENGTH = 20_000
    CHARACTER_NAME_MAX_LENGTH = 210

    def run(self, **params: Any) -> Any:
        """Create a reusable character (synchronous).

        Args:
            **params: Reusable character parameters (descriptions, reference_image_url, ...).

        Returns:
            The result.
        """
        compacted = self._compact_params(params)
        self._validate_params(compacted)
        return self._request("post", self.ENDPOINT, body=compacted)

    def _validate_params(self, params: Dict[str, Any]) -> None:
        self._validate_contract(CONTRACT["create-character"], {**params, "model": self.MODEL})
        if params.get("audio_ids") is not None:
            self._validate_array(params, "audio_ids")
        self._validate_length(params, "descriptions", self.DESCRIPTIONS_MAX_LENGTH)
        self._validate_length(params, "character_name", self.CHARACTER_NAME_MAX_LENGTH)

    @staticmethod
    def _validate_array(params: Dict[str, Any], key: str) -> None:
        if not isinstance(params.get(key), list):
            raise ValidationError(f"{key} must be an array")

    @staticmethod
    def _validate_length(params: Dict[str, Any], key: str, max_length: int) -> None:
        value = params.get(key)
        if value is None or len(str(value)) <= max_length:
            return
        raise ValidationError(f"{key} must be at most {max_length} characters")
