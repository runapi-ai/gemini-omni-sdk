// Package geminiomni provides the Gemini Omni API client.
package geminiomni

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const createAudioPath = "/api/v1/gemini_omni/create_audio"
const createCharacterPath = "/api/v1/gemini_omni/create_character"
const textToVideoPath = "/api/v1/gemini_omni/text_to_video"

type Client struct {
	CreateAudio     *CreateAudio
	CreateCharacter *CreateCharacter
	TextToVideo     *TextToVideo
}

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

func NewClientWithHTTP(httpClient core.HTTPClient) *Client {
	return &Client{
		CreateAudio:     &CreateAudio{http: httpClient},
		CreateCharacter: &CreateCharacter{http: httpClient},
		TextToVideo:     &TextToVideo{http: httpClient},
	}
}

type CreateAudio struct{ http core.HTTPClient }

func (r *CreateAudio) Run(ctx context.Context, params CreateAudioParams, opts ...option.RequestOption) (*CreateAudioResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[CreateAudioResponse](ctx, r.http, createAudioPath, core.CompactParams(params), requestOptions)
}

type CreateCharacter struct{ http core.HTTPClient }

func (r *CreateCharacter) Run(ctx context.Context, params CreateCharacterParams, opts ...option.RequestOption) (*CreateCharacterResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[CreateCharacterResponse](ctx, r.http, createCharacterPath, core.CompactParams(params), requestOptions)
}

type TextToVideo struct{ http core.HTTPClient }

func (r *TextToVideo) Create(ctx context.Context, params TextToVideoParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, textToVideoPath, core.CompactParams(params), requestOptions)
}

func (r *TextToVideo) Get(ctx context.Context, id string, opts ...option.RequestOption) (*TextToVideoResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[TextToVideoResponse](ctx, r.http, core.ResourcePath(textToVideoPath, id), requestOptions)
}

func (r *TextToVideo) Run(ctx context.Context, params TextToVideoParams, opts ...option.RequestOption) (*TextToVideoResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) {
		return r.Create(ctx, params, opts...)
	}, func(ctx context.Context, id string) (*TextToVideoResponse, error) {
		return r.Get(ctx, id, opts...)
	}, pollingOptions)
}
