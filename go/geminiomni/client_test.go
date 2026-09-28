package geminiomni

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

type stubHTTPClient struct {
	method string
	path   string
	body   any
}

func (s *stubHTTPClient) Request(_ context.Context, method, path string, opts *core.HTTPRequestOptions) (json.RawMessage, error) {
	s.method = method
	s.path = path
	if opts != nil {
		s.body = opts.Body
	}
	if path == "/api/v1/gemini_omni/create_character" {
		return json.RawMessage(`{"id":"character-runapi-123","character":{"id":"character-runapi-123","name":"Jenny","images":[{"url":"https://file.runapi.ai/gemini/jenny.png"},{"url":"https://file.runapi.ai/gemini/jenny-body.png"}]}}`), nil
	}
	if path == "/api/v1/gemini_omni/text_to_video" {
		return json.RawMessage(`{"id":"task-local-123","status":"processing"}`), nil
	}
	if path == "/api/v1/gemini_omni/text_to_video/task-local-123" {
		return json.RawMessage(`{"id":"task-local-123","status":"completed", "usage": {"cost": 0.05},"videos":[{"url":"https://tempfile.runapi.ai/gemini/output.mp4"}]}`), nil
	}
	return json.RawMessage(`{"id":"audio-runapi-123","audio":{"id":"audio-runapi-123","name":"Acher Narrator"}}`), nil
}

func TestCreateAudioRunSendsCorrectRequest(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	resp, err := client.CreateAudio.Run(context.Background(), CreateAudioParams{
		AudioID:          VoiceAchernar,
		Name:             "Acher Narrator",
		VoiceDescription: "A calm, clear voice",
		ExampleDialogue:  "Hello, I am achernar"})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/gemini_omni/create_audio" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body, ok := stub.body.(map[string]any)
	if !ok {
		t.Fatalf("expected flat body map, got %T", stub.body)
	}
	if body["audio_id"] != string(VoiceAchernar) || body["voice_description"] != "A calm, clear voice" {
		t.Fatalf("unexpected body: %#v", body)
	}
	if _, ok := body["audioId"]; ok {
		t.Fatalf("unexpected camelCase key in body: %#v", body)
	}
	if resp.ID != "audio-runapi-123" {
		t.Fatalf("unexpected response id: %s", resp.ID)
	}
}

func TestCreateCharacterRunSendsCorrectRequest(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	resp, err := client.CreateCharacter.Run(context.Background(), CreateCharacterParams{
		Descriptions:          "A silver-haired cyberpunk guide",
		ReferenceImageURL:     "https://file.runapi.ai/demo/character.png",
		BodyReferenceImageURL: "https://file.runapi.ai/demo/character-body.png",
		AudioIDs:              []string{"audio-runapi-123"},
		CharacterName:         "Jenny"})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/gemini_omni/create_character" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body, ok := stub.body.(map[string]any)
	if !ok {
		t.Fatalf("expected flat body map, got %T", stub.body)
	}
	if body["descriptions"] != "A silver-haired cyberpunk guide" {
		t.Fatalf("unexpected body: %#v", body)
	}
	if _, ok := body["description"]; ok {
		t.Fatalf("unexpected singular description key in body: %#v", body)
	}
	if body["reference_image_url"] != "https://file.runapi.ai/demo/character.png" {
		t.Fatalf("unexpected reference image in body: %#v", body)
	}
	if body["body_reference_image_url"] != "https://file.runapi.ai/demo/character-body.png" {
		t.Fatalf("unexpected body reference image in body: %#v", body)
	}
	if _, ok := body["image_urls"]; ok {
		t.Fatalf("unexpected image_urls key in body: %#v", body)
	}
	if _, ok := body["imageUrls"]; ok {
		t.Fatalf("unexpected camelCase key in body: %#v", body)
	}
	if resp.ID != "character-runapi-123" {
		t.Fatalf("unexpected response id: %s", resp.ID)
	}
	if resp.Character == nil || len(resp.Character.Images) != 2 || resp.Character.Images[1].URL != "https://file.runapi.ai/gemini/jenny-body.png" {
		t.Fatalf("unexpected character images: %#v", resp.Character)
	}
}

func TestCreateCharacterRunFollowsAcceptedTaskLocation(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		switch requests {
		case 1:
			if r.Method != http.MethodPost || r.URL.Path != createCharacterPath {
				t.Fatalf("unexpected create request: %s %s", r.Method, r.URL.Path)
			}
			w.Header().Set("Location", "/api/v1/tasks/task_pending")
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusAccepted)
			_, _ = w.Write([]byte(`{"id":"task_pending","status":"pending"}`))
		case 2:
			if r.Method != http.MethodGet || r.URL.Path != "/api/v1/tasks/task_pending" {
				t.Fatalf("unexpected task result request: %s %s", r.Method, r.URL.Path)
			}
			_, _ = w.Write([]byte(`{"id":"task_pending","status":"completed", "usage": {"cost": 0.05},"response":{"status":200,"content_type":"application/json","headers":{},"body":{"id":"character-runapi-123","character":{"id":"character-runapi-123","name":"Jenny","images":[]}}}}`))
		default:
			t.Fatalf("unexpected request %d", requests)
		}
	}))
	defer server.Close()

	client, err := NewClient(option.WithAPIKey("test-key"), option.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.CreateCharacter.Run(
		context.Background(),
		CreateCharacterParams{
			Descriptions:      "A silver-haired cyberpunk guide",
			ReferenceImageURL: "https://file.runapi.ai/demo/character.png",
			CharacterName:     "Jenny"},
		option.WithHeader("Prefer", "wait=0"),
		option.WithPollInterval(time.Millisecond),
	)
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != "character-runapi-123" || result.Character == nil || result.Character.ID != "character-runapi-123" {
		t.Fatalf("unexpected terminal result: %#v", result)
	}
}

func TestTextToVideoCreateAndGet(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	seed := 12345
	created, err := client.TextToVideo.Create(context.Background(), TextToVideoParams{
		Prompt:             "Create a neon city tracking shot",
		DurationSeconds:    8,
		AspectRatio:        "16:9",
		OutputResolution:   "1080p",
		ReferenceImageURLs: []string{"https://file.runapi.ai/demo/scene.png"},
		AudioIDs:           []string{"audio-runapi-123"},
		CharacterIDs:       []string{"character-runapi-123"},
		Seed:               &seed})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/gemini_omni/text_to_video" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body, ok := stub.body.(map[string]any)
	if !ok {
		t.Fatalf("expected flat body map, got %T", stub.body)
	}
	if body["prompt"] != "Create a neon city tracking shot" || body["duration_seconds"] != float64(8) || body["output_resolution"] != "1080p" {
		t.Fatalf("unexpected body: %#v", body)
	}
	if _, ok := body["image_urls"]; ok {
		t.Fatalf("unexpected image_urls key in body: %#v", body)
	}
	if values, ok := body["reference_image_urls"].([]any); !ok || len(values) != 1 || values[0] != "https://file.runapi.ai/demo/scene.png" {
		t.Fatalf("unexpected reference_image_urls in body: %#v", body)
	}
	if _, ok := body["characterIds"]; ok {
		t.Fatalf("unexpected camelCase key in body: %#v", body)
	}
	if created.ID != "task-local-123" {
		t.Fatalf("unexpected task id: %s", created.ID)
	}

	got, err := client.TextToVideo.Get(context.Background(), "task-local-123")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != "/api/v1/gemini_omni/text_to_video/task-local-123" {
		t.Fatalf("unexpected get request: %s %s", stub.method, stub.path)
	}
	if got.GetStatus() != "completed" || got.Videos[0].URL != "https://tempfile.runapi.ai/gemini/output.mp4" {
		t.Fatalf("unexpected response: %#v", got)
	}
}

func TestTextToVideoRejectsInvalidVideoClip(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	_, err := client.TextToVideo.Create(context.Background(), TextToVideoParams{
		Prompt:          "Create a neon city tracking shot",
		DurationSeconds: 8,
		VideoList:       []VideoClip{{URL: "", Start: 0, Ends: 10}}})
	if err == nil || err.Error() != "video_list[0].url is required" {
		t.Fatalf("expected video clip validation error, got %v", err)
	}
	if stub.method != "" {
		t.Fatalf("request should not be sent, got %s %s", stub.method, stub.path)
	}
}

func TestTextToVideoCreateFlashPreviewSendsModelWithoutDuration(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	created, err := client.TextToVideo.Create(context.Background(), TextToVideoParams{
		Model:            ModelGeminiOmniFlashPreview,
		Prompt:           "A paper airplane flying through a sunlit studio",
		AspectRatio:      "9:16",
		OutputResolution: "720p"})
	if err != nil {
		t.Fatal(err)
	}
	body, ok := stub.body.(map[string]any)
	if !ok {
		t.Fatalf("expected flat body map, got %T", stub.body)
	}
	if body["model"] != string(ModelGeminiOmniFlashPreview) {
		t.Fatalf("expected Flash Preview model in body: %#v", body)
	}
	if _, ok := body["duration_seconds"]; ok {
		t.Fatalf("unexpected duration_seconds in Flash Preview body: %#v", body)
	}
	if created.ID != "task-local-123" {
		t.Fatalf("unexpected task id: %s", created.ID)
	}
}

func TestTextToVideoCreateFlash11SendsFrameFieldsAnd360p(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	created, err := client.TextToVideo.Create(context.Background(), TextToVideoParams{
		Model:              ModelGeminiOmniFlash11,
		Prompt:             "A paper airplane crosses from dawn into dusk",
		DurationSeconds:    6,
		FirstFrameImageURL: "https://cdn.runapi.ai/public/samples/first-frame.jpg",
		LastFrameImageURL:  "https://cdn.runapi.ai/public/samples/last-frame.jpg",
		AspectRatio:        "16:9",
		OutputResolution:   "360p"})
	if err != nil {
		t.Fatal(err)
	}
	body, ok := stub.body.(map[string]any)
	if !ok {
		t.Fatalf("expected flat body map, got %T", stub.body)
	}
	if body["model"] != string(ModelGeminiOmniFlash11) || body["first_frame_image_url"] != "https://cdn.runapi.ai/public/samples/first-frame.jpg" || body["last_frame_image_url"] != "https://cdn.runapi.ai/public/samples/last-frame.jpg" {
		t.Fatalf("unexpected Flash 1.1 body: %#v", body)
	}
	if body["output_resolution"] != "360p" {
		t.Fatalf("expected 360p output resolution: %#v", body)
	}
	if _, ok := body["first_frame_url"]; ok {
		t.Fatalf("unexpected internal frame key in body: %#v", body)
	}
	if created.ID != "task-local-123" {
		t.Fatalf("unexpected task id: %s", created.ID)
	}
}

func TestTextToVideoFlash11EnforcesFrameRules(t *testing.T) {
	tests := []struct {
		name   string
		params TextToVideoParams
		error  string
	}{
		{
			name: "first frame forbids reference images",
			params: TextToVideoParams{
				Model:              ModelGeminiOmniFlash11,
				Prompt:             "A paper airplane crosses from dawn into dusk",
				DurationSeconds:    6,
				FirstFrameImageURL: "https://cdn.runapi.ai/public/samples/first-frame.jpg",
				ReferenceImageURLs: []string{"https://cdn.runapi.ai/public/samples/reference-1.jpg"}},
			error: "reference_image_urls is not allowed when first_frame_image_url is present and model is gemini-omni-flash-1-1"},
		{
			name: "last frame requires first frame",
			params: TextToVideoParams{
				Model:             ModelGeminiOmniFlash11,
				Prompt:            "A paper airplane crosses from dawn into dusk",
				DurationSeconds:   6,
				LastFrameImageURL: "https://cdn.runapi.ai/public/samples/last-frame.jpg"},
			error: "first_frame_image_url is required when last_frame_image_url is present and model is gemini-omni-flash-1-1"}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stub := &stubHTTPClient{}
			client := NewClientWithHTTP(stub)
			_, err := client.TextToVideo.Create(context.Background(), test.params)
			if err == nil || err.Error() != test.error {
				t.Fatalf("expected %q, got %v", test.error, err)
			}
			if stub.method != "" {
				t.Fatalf("request should not be sent, got %s %s", stub.method, stub.path)
			}
		})
	}
}
