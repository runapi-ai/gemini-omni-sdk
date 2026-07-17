"""Gemini Omni text-to-video resource."""

from __future__ import annotations

from typing import Any, Dict, List, Optional

from runapi.core import Resource, ValidationError, RequestOptions

from ..contract_gen import CONTRACT
from ..types import (
    SEED_MAX,
    SEED_MIN,
    CompletedTextToVideoResponse,
    TextToVideoResponse,
)


class TextToVideo(Resource):
    """Generate video from a prompt, reference media, and characters."""

    ENDPOINT = "/api/v1/gemini_omni/text_to_video"

    RESPONSE_CLASS = TextToVideoResponse
    COMPLETED_RESPONSE_CLASS = CompletedTextToVideoResponse

    MODEL = "gemini-omni-text-to-video"

    PROMPT_MAX_LENGTH = 20_000
    REFERENCE_IMAGE_URLS_MAX = 7
    AUDIO_IDS_MAX = 3
    VIDEO_LIST_MAX = 1
    CHARACTER_IDS_MAX = 3
    REFERENCE_UNITS_MAX = 7
    VIDEO_REFERENCE_UNITS = 2
    MAX_TRIM_SECONDS = 10

    def run(self, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        """Create a text-to-video task and poll until it completes.

        Args:
            **params: Text-to-video parameters (model, prompt, ...).

        Returns:
            The completed text-to-video response.
        """
        task = self.create(options=options, **params)
        return self._poll_until_complete(lambda: self.get(task.id, options=options))

    def create(self, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        """Create a text-to-video task and return immediately with an id.

        Args:
            **params: Text-to-video parameters (model, prompt, ...).

        Returns:
            The task creation result with an id.
        """
        compacted = self._compact_params(params)
        self._validate_params(compacted)
        return self._request("post", self.ENDPOINT, body=compacted, options=options)

    def get(self, id: str, options: Optional[RequestOptions] = None) -> Any:
        """Fetch the current status of a text-to-video task.

        Args:
            id: Task id.

        Returns:
            The current text-to-video status.
        """
        return self._request("get", f"{self.ENDPOINT}/{id}", options=options)

    def _validate_params(self, params: Dict[str, Any]) -> None:
        self._validate_contract(CONTRACT["text-to-video"], {**params, "model": self.MODEL})
        self._validate_length(params, "prompt", self.PROMPT_MAX_LENGTH)
        if params.get("reference_image_urls") is not None:
            self._validate_array(params, "reference_image_urls", self.REFERENCE_IMAGE_URLS_MAX)
        if params.get("audio_ids") is not None:
            self._validate_array(params, "audio_ids", self.AUDIO_IDS_MAX)
        if params.get("video_list") is not None:
            self._validate_array(params, "video_list", self.VIDEO_LIST_MAX)
        if params.get("character_ids") is not None:
            self._validate_array(params, "character_ids", self.CHARACTER_IDS_MAX)
        if params.get("video_list") is not None:
            self._validate_video_list(params.get("video_list"))
        self._validate_reference_units(params)
        self._validate_seed(params)

    @staticmethod
    def _validate_length(params: Dict[str, Any], key: str, max_length: int) -> None:
        value = params.get(key)
        if value is None or len(str(value)) <= max_length:
            return
        raise ValidationError(f"{key} must be at most {max_length} characters")

    @staticmethod
    def _validate_array(params: Dict[str, Any], key: str, max_length: int) -> None:
        value = params.get(key)
        if not isinstance(value, list):
            raise ValidationError(f"{key} must be an array")
        if len(value) <= max_length:
            return
        raise ValidationError(f"{key} accepts at most {max_length} items")

    def _validate_video_list(self, items: List[Dict[str, Any]]) -> None:
        for index, item in enumerate(items):
            url = item.get("url") if isinstance(item, dict) else None
            if not url:
                raise ValidationError(f"video_list[{index}].url is required")
            start_time = self._to_float(item.get("start"))
            end_time = self._to_float(item.get("ends"))
            if start_time is None or end_time is None:
                raise ValidationError(f"video_list[{index}] start and ends must be numbers")
            if start_time < 0:
                raise ValidationError(f"video_list[{index}].start must be 0 or greater")
            if not end_time > start_time:
                raise ValidationError(f"video_list[{index}].ends must be greater than start")
            if (end_time - start_time) > self.MAX_TRIM_SECONDS:
                raise ValidationError(
                    f"video_list[{index}] trim range must be {self.MAX_TRIM_SECONDS} seconds or less"
                )

    def _validate_reference_units(self, params: Dict[str, Any]) -> None:
        units = (
            self._count(params.get("reference_image_urls"))
            + self._count(params.get("video_list")) * self.VIDEO_REFERENCE_UNITS
            + self._count(params.get("character_ids"))
        )
        if units <= self.REFERENCE_UNITS_MAX:
            return
        raise ValidationError(
            f"reference_image_urls + video_list*2 + character_ids must use {self.REFERENCE_UNITS_MAX} reference units or fewer"
        )

    @staticmethod
    def _validate_seed(params: Dict[str, Any]) -> None:
        value = params.get("seed")
        if value is None:
            return
        try:
            seed = int(value)
        except (TypeError, ValueError):
            seed = None
        if seed is not None and SEED_MIN <= seed <= SEED_MAX:
            return
        raise ValidationError(f"seed must be an integer between {SEED_MIN} and {SEED_MAX}")

    @staticmethod
    def _count(value: Any) -> int:
        if value is None:
            return 0
        return len(value) if isinstance(value, list) else 1

    @staticmethod
    def _to_float(value: Any) -> Optional[float]:
        if isinstance(value, bool):
            return None
        try:
            return float(value)
        except (TypeError, ValueError):
            return None
