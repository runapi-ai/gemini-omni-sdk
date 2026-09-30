"""Gemini Omni create-character terminal-or-accepted Task resource."""

from __future__ import annotations

from typing import Any, Optional

from runapi.core import Resource, RequestOptions

from ..types import CreateCharacterResponse


class CreateCharacter(Resource):
    """Create a character from a portrait and optional full-body reference image."""

    ENDPOINT = "/api/v1/gemini_omni/create_character"

    RESPONSE_CLASS = CreateCharacterResponse

    MODEL = "gemini-omni-character"

    def run(self, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        """Create a reusable character and follow an accepted Task to completion.

        Args:
            **params: Reusable character parameters. ``reference_image_url`` is the
                portrait; ``body_reference_image_url`` optionally adds a full-body reference.

        Returns:
            The result.
        """
        compacted = self._compact_params(params)
        return self._run_hybrid("post", self.ENDPOINT, body=compacted, options=options)
