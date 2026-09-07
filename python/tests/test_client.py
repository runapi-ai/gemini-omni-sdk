import pytest

from runapi.core import ApiResponse, config
from runapi.core.errors import AuthenticationError, ValidationError
from runapi.gemini_omni import GeminiOmniClient
from runapi.gemini_omni.resources.create_audio import CreateAudio
from runapi.gemini_omni.resources.create_character import CreateCharacter
from runapi.gemini_omni.resources.text_to_video import TextToVideo
from runapi.gemini_omni.types import (
    CompletedTextToVideoResponse,
    CreateAudioResponse,
    CreateCharacterResponse,
)


class FakeHttp:
    def __init__(self, *responses):
        self._responses = list(responses)
        self.calls = []

    def request(self, method, path, body=None, options=None):
        self.calls.append((method, path, body))
        if self._responses:
            return self._responses.pop(0)
        return {"id": "task_1", "status": "pending"}


@pytest.fixture(autouse=True)
def reset_config(monkeypatch):
    monkeypatch.delenv("RUNAPI_API_KEY", raising=False)
    monkeypatch.setattr(config, "api_key", None)
    yield


# --- auth -----------------------------------------------------------------


def test_accepts_api_key_parameter():
    assert isinstance(GeminiOmniClient(api_key="k", http_client=FakeHttp()), GeminiOmniClient)


def test_falls_back_to_global(monkeypatch):
    monkeypatch.setattr(config, "api_key", "global-key")
    assert isinstance(GeminiOmniClient(http_client=FakeHttp()), GeminiOmniClient)


def test_falls_back_to_env(monkeypatch):
    monkeypatch.setenv("RUNAPI_API_KEY", "env-key")
    assert isinstance(GeminiOmniClient(http_client=FakeHttp()), GeminiOmniClient)


def test_raises_without_api_key():
    with pytest.raises(AuthenticationError, match="API key is required"):
        GeminiOmniClient()


def test_uses_injected_http_client_and_accessors():
    fake = FakeHttp()
    client = GeminiOmniClient(api_key="k", http_client=fake)
    assert isinstance(client.create_audio, CreateAudio)
    assert isinstance(client.create_character, CreateCharacter)
    assert isinstance(client.text_to_video, TextToVideo)
    assert client.text_to_video._http is fake


# --- terminal resources ---------------------------------------------------


def test_create_audio_posts_once_and_returns_typed():
    fake = FakeHttp({"id": "a1", "audio": {"id": "kore", "name": "Narrator"}})
    client = GeminiOmniClient(api_key="k", http_client=fake)
    result = client.create_audio.run(audio_id="kore", name="Narrator")
    assert fake.calls == [("post", "/api/v1/gemini_omni/create_audio", {"audio_id": "kore", "name": "Narrator"})]
    assert isinstance(result, CreateAudioResponse)
    assert result.audio.id == "kore"


def test_create_audio_requires_audio_id_and_name():
    client = GeminiOmniClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="audio_id is required"):
        client.create_audio.run(name="Narrator")
    with pytest.raises(ValidationError, match="name is required"):
        client.create_audio.run(audio_id="kore")


def test_create_audio_name_length():
    client = GeminiOmniClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="name must be at most 210 characters"):
        client.create_audio.run(audio_id="kore", name="x" * 211)


def test_create_character_returns_typed_and_validates():
    fake = FakeHttp({"id": "c1", "character": {"id": "c1", "name": "Robot", "images": [{"url": "https://x/r.png"}, {"url": "https://x/b.png"}]}})
    client = GeminiOmniClient(api_key="k", http_client=fake)
    result = client.create_character.run(
        descriptions="A robot",
        reference_image_url="https://x/r.png",
        body_reference_image_url="https://x/b.png",
    )
    assert fake.calls[0] == (
        "post",
        "/api/v1/gemini_omni/create_character",
        {
            "descriptions": "A robot",
            "reference_image_url": "https://x/r.png",
            "body_reference_image_url": "https://x/b.png",
        },
    )
    assert isinstance(result, CreateCharacterResponse)
    assert result.character.id == "c1"
    assert result.character.images[1].url == "https://x/b.png"


def test_create_character_follows_accepted_task_result():
    location = "https://runapi.ai/api/v1/tasks/task_1/result"
    fake = FakeHttp(
        ApiResponse(
            {"id": "task_1", "status": "processing"},
            {"Location": location, "Retry-After": "0"},
            status_code=202,
        ),
        ApiResponse(
            {
                "id": "task_1",
                "status": "completed",
                "response": {
                    "status": 200,
                    "content_type": "application/json",
                    "headers": {},
                    "body": {
                        "id": "character_1",
                        "character": {
                            "id": "character_1",
                            "name": "Robot",
                            "images": [],
                        },
                    },
                },
            }
        ),
    )
    client = GeminiOmniClient(api_key="k", http_client=fake)

    result = client.create_character.run(
        descriptions="A robot", reference_image_url="https://x/r.png"
    )

    assert isinstance(result, CreateCharacterResponse)
    assert result.id == "character_1"
    assert result.character.id == "character_1"
    assert [call[:2] for call in fake.calls] == [
        ("post", "/api/v1/gemini_omni/create_character"),
        ("get", location),
    ]


def test_create_character_requires_fields():
    client = GeminiOmniClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="descriptions is required"):
        client.create_character.run(reference_image_url="https://x/r.png")
    with pytest.raises(ValidationError, match="reference_image_url is required"):
        client.create_character.run(descriptions="A robot")


def test_create_character_audio_ids_must_be_array():
    client = GeminiOmniClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="audio_ids must be an array"):
        client.create_character.run(
            descriptions="A robot", reference_image_url="https://x/r.png", audio_ids="not-a-list"
        )


# --- async text_to_video --------------------------------------------------


def test_text_to_video_create_and_get_shapes():
    fake = FakeHttp({"id": "t1", "status": "pending"}, {"id": "t1", "status": "processing"})
    client = GeminiOmniClient(api_key="k", http_client=fake)
    client.text_to_video.create(prompt="a fox", duration_seconds=8)
    client.text_to_video.get("t1")
    assert fake.calls == [
        ("post", "/api/v1/gemini_omni/text_to_video", {"prompt": "a fox", "duration_seconds": 8}),
        ("get", "/api/v1/gemini_omni/text_to_video/t1", None),
    ]


def test_text_to_video_flash_preview_sends_model_without_duration():
    fake = FakeHttp({"id": "t-flash", "status": "pending"})
    client = GeminiOmniClient(api_key="k", http_client=fake)
    client.text_to_video.create(
        model="gemini-omni-flash-preview",
        prompt="A paper airplane flying through a sunlit studio",
        aspect_ratio="9:16",
        output_resolution="720p",
    )
    assert fake.calls == [
        (
            "post",
            "/api/v1/gemini_omni/text_to_video",
            {
                "model": "gemini-omni-flash-preview",
                "prompt": "A paper airplane flying through a sunlit studio",
                "aspect_ratio": "9:16",
                "output_resolution": "720p",
            },
        )
    ]


def test_text_to_video_flash_1_1_sends_frame_fields_and_360p():
    fake = FakeHttp({"id": "t-flash-1-1", "status": "pending"})
    client = GeminiOmniClient(api_key="k", http_client=fake)
    client.text_to_video.create(
        model="gemini-omni-flash-1-1",
        prompt="A paper airplane crosses from dawn into dusk",
        duration_seconds=6,
        first_frame_image_url="https://cdn.runapi.ai/public/samples/first-frame.jpg",
        last_frame_image_url="https://cdn.runapi.ai/public/samples/last-frame.jpg",
        aspect_ratio="16:9",
        output_resolution="360p",
    )
    assert fake.calls == [
        (
            "post",
            "/api/v1/gemini_omni/text_to_video",
            {
                "model": "gemini-omni-flash-1-1",
                "prompt": "A paper airplane crosses from dawn into dusk",
                "duration_seconds": 6,
                "first_frame_image_url": "https://cdn.runapi.ai/public/samples/first-frame.jpg",
                "last_frame_image_url": "https://cdn.runapi.ai/public/samples/last-frame.jpg",
                "aspect_ratio": "16:9",
                "output_resolution": "360p",
            },
        )
    ]


def test_text_to_video_flash_1_1_enforces_frame_rules():
    client = GeminiOmniClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(
        ValidationError,
        match="reference_image_urls is not allowed when model is gemini-omni-flash-1-1 and first_frame_image_url is present",
    ):
        client.text_to_video.create(
            model="gemini-omni-flash-1-1",
            prompt="A paper airplane crosses from dawn into dusk",
            duration_seconds=6,
            first_frame_image_url="https://cdn.runapi.ai/public/samples/first-frame.jpg",
            reference_image_urls=["https://cdn.runapi.ai/public/samples/reference-1.jpg"],
        )
    with pytest.raises(
        ValidationError,
        match="first_frame_image_url is required when model is gemini-omni-flash-1-1 and last_frame_image_url is present",
    ):
        client.text_to_video.create(
            model="gemini-omni-flash-1-1",
            prompt="A paper airplane crosses from dawn into dusk",
            duration_seconds=6,
            last_frame_image_url="https://cdn.runapi.ai/public/samples/last-frame.jpg",
        )


def test_text_to_video_requires_prompt_and_duration():
    client = GeminiOmniClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="prompt is required"):
        client.text_to_video.create(duration_seconds=8)
    with pytest.raises(ValidationError, match="duration_seconds is required"):
        client.text_to_video.create(prompt="a fox")


def test_text_to_video_duration_enum():
    client = GeminiOmniClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="duration_seconds must be one of: 4, 6, 8, 10"):
        client.text_to_video.create(prompt="a fox", duration_seconds=7)


def test_text_to_video_reference_units_cap():
    client = GeminiOmniClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="reference units"):
        client.text_to_video.create(
            prompt="a fox",
            duration_seconds=8,
            reference_image_urls=["a", "b", "c", "d", "e"],  # 5, within array max (7)
            character_ids=["x", "y", "z"],  # 3, within array max (3); 5 + 3 = 8 > 7 units
        )


def test_text_to_video_run_narrows_completed():
    fake = FakeHttp(
        {"id": "t1", "status": "pending"},
        {"id": "t1", "status": "completed", "videos": [{"url": "https://x/v.mp4"}]},
    )
    client = GeminiOmniClient(api_key="k", http_client=fake)
    result = client.text_to_video.run(prompt="a fox", duration_seconds=8)
    assert isinstance(result, CompletedTextToVideoResponse)
    assert result.videos[0].url == "https://x/v.mp4"
