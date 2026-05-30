package geminiomni

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/runapi-ai/core-sdk/go/core"
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
		return json.RawMessage(`{"id":"character-kie-123","character":{"id":"character-kie-123","name":"Jenny","image_url":"https://file.runapi.ai/gemini/jenny.png"}}`), nil
	}
	if path == "/api/v1/gemini_omni/text_to_video" {
		return json.RawMessage(`{"id":"task-local-123","status":"processing"}`), nil
	}
	if path == "/api/v1/gemini_omni/text_to_video/task-local-123" {
		return json.RawMessage(`{"id":"task-local-123","status":"completed","videos":[{"url":"https://tempfile.runapi.ai/gemini/output.mp4"}]}`), nil
	}
	return json.RawMessage(`{"id":"audio-kie-123","audio":{"id":"audio-kie-123","name":"Acher Narrator"}}`), nil
}

func TestCreateAudioRunSendsCorrectRequest(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	resp, err := client.CreateAudio.Run(context.Background(), CreateAudioParams{
		AudioID:          VoiceAchernar,
		Name:             "Acher Narrator",
		VoiceDescription: "A calm, clear voice",
		ExampleDialogue:  "Hello, I am achernar",
	})
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
	if resp.ID != "audio-kie-123" {
		t.Fatalf("unexpected response id: %s", resp.ID)
	}
}

func TestCreateCharacterRunSendsCorrectRequest(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	resp, err := client.CreateCharacter.Run(context.Background(), CreateCharacterParams{
		Descriptions:  "A silver-haired cyberpunk guide",
		ImageURLs:     []string{"https://file.runapi.ai/demo/character.png"},
		AudioIDs:      []string{"audio-kie-123"},
		CharacterName: "Jenny",
	})
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
	if _, ok := body["imageUrls"]; ok {
		t.Fatalf("unexpected camelCase key in body: %#v", body)
	}
	if resp.ID != "character-kie-123" {
		t.Fatalf("unexpected response id: %s", resp.ID)
	}
}

func TestTextToVideoCreateAndGet(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	seed := 12345
	created, err := client.TextToVideo.Create(context.Background(), TextToVideoParams{
		Prompt:       "Create a neon city tracking shot",
		Duration:     "8",
		AspectRatio:  "16:9",
		Resolution:   "1080p",
		ImageURLs:    []string{"https://file.runapi.ai/demo/scene.png"},
		AudioIDs:     []string{"audio-kie-123"},
		CharacterIDs: []string{"character-kie-123"},
		Seed:         &seed,
	})
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
	if body["prompt"] != "Create a neon city tracking shot" || body["duration"] != "8" || body["resolution"] != "1080p" {
		t.Fatalf("unexpected body: %#v", body)
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
