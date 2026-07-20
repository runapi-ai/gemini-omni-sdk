// Package geminiomni provides the Gemini Omni API client.
package geminiomni

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/base"
	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const createAudioPath = "/api/v1/gemini_omni/create_audio"
const createCharacterPath = "/api/v1/gemini_omni/create_character"
const textToVideoPath = "/api/v1/gemini_omni/text_to_video"

// Fixed-model endpoints inject their model only into a validation copy. The
// text-to-video endpoint preserves an explicitly selected model on the wire.
const createAudioModel = "gemini-omni-audio"
const createCharacterModel = "gemini-omni-character"
const textToVideoModel = "gemini-omni-text-to-video"

// validateAction validates a compacted request body against one contract
// action, injecting the endpoint's fixed model (never posted) so contract
// model-membership and per-field checks apply.
func validateAction(action, model string, body map[string]any) error {
	withModel := make(map[string]any, len(body)+1)
	for key, value := range body {
		withModel[key] = value
	}
	if selected, ok := withModel["model"]; !ok || selected == nil || selected == "" {
		withModel["model"] = model
	}
	return core.ValidateParams(contractSchema[action], withModel)
}

// Client provides Gemini Omni multimodal generation: voice presets, character creation, and text-to-video.
type Client struct {
	base.Base
	CreateAudio     *CreateAudio
	CreateCharacter *CreateCharacter
	TextToVideo     *TextToVideo
}

// NewClient creates a Gemini Omni client with the given options.
func NewClient(opts ...option.ClientOption) (*Client, error) {
	resolved, err := option.ResolveClientOptions(opts...)
	if err != nil {
		return nil, err
	}
	httpClient, err := core.NewHTTPClient(resolved)
	if err != nil {
		return nil, err
	}
	return NewClientWithHTTP(httpClient), nil
}

// NewClientWithHTTP creates a Gemini Omni client with a pre-configured HTTP transport.
func NewClientWithHTTP(httpClient core.HTTPClient) *Client {
	return &Client{
		Base:            base.New(httpClient),
		CreateAudio:     &CreateAudio{http: httpClient},
		CreateCharacter: &CreateCharacter{http: httpClient},
		TextToVideo:     &TextToVideo{http: httpClient},
	}
}

// CreateAudio registers a reusable voice preset from a built-in [AudioVoice] identity.
// This is synchronous -- only Run is available (no Create/Get polling).
type CreateAudio struct{ http core.HTTPClient }

// Run submits a Gemini Omni audio creation task and returns the result.
func (r *CreateAudio) Run(ctx context.Context, params CreateAudioParams, opts ...option.RequestOption) (*CreateAudioResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	body := core.CompactParams(params)
	if err := validateAction("create-audio", createAudioModel, body); err != nil {
		return nil, err
	}
	return core.PostJSON[CreateAudioResponse](ctx, r.http, createAudioPath, body, requestOptions)
}

// CreateCharacter builds a reusable character from a reference image and description.
// Attach AudioIDs to give the character a specific voice. This is synchronous -- only Run is available.
type CreateCharacter struct{ http core.HTTPClient }

// Run submits a Gemini Omni character creation task and returns the result.
func (r *CreateCharacter) Run(ctx context.Context, params CreateCharacterParams, opts ...option.RequestOption) (*CreateCharacterResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	body := core.CompactParams(params)
	if err := validateAction("create-character", createCharacterModel, body); err != nil {
		return nil, err
	}
	return core.PostJSON[CreateCharacterResponse](ctx, r.http, createCharacterPath, body, requestOptions)
}

// TextToVideo generates video from a prompt with optional characters, audio voices, reference images, and video clips.
// This is async -- use Create+Get for manual polling, or Run for automatic polling.
type TextToVideo struct{ http core.HTTPClient }

// Create submits a Gemini Omni text-to-video task and returns immediately with a task id.
func (r *TextToVideo) Create(ctx context.Context, params TextToVideoParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	body := core.CompactParams(params)
	if err := validateAction("text-to-video", textToVideoModel, body); err != nil {
		return nil, err
	}
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, textToVideoPath, body, requestOptions)
}

// Get fetches the current status of a Gemini Omni text-to-video task by id.
func (r *TextToVideo) Get(ctx context.Context, id string, opts ...option.RequestOption) (*TextToVideoResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[TextToVideoResponse](ctx, r.http, core.ResourcePath(textToVideoPath, id), requestOptions)
}

// Run submits a Gemini Omni text-to-video task and polls until it completes.
func (r *TextToVideo) Run(ctx context.Context, params TextToVideoParams, opts ...option.RequestOption) (*TextToVideoResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) {
		return r.Create(ctx, params, opts...)
	}, func(ctx context.Context, id string) (*TextToVideoResponse, error) {
		return r.Get(ctx, id, opts...)
	}, pollingOptions)
}
