package geminiomni

import "github.com/runapi-ai/core-sdk/go/core"

// AudioVoice selects one of the 30 preset voice identities for audio creation.
// Each voice has a distinct pitch, cadence, and personality. Pass the chosen constant
// to [CreateAudioParams].AudioID.
type AudioVoice string

const (
	VoiceAchernar      AudioVoice = "achernar"
	VoiceAchird        AudioVoice = "achird"
	VoiceAlgenib       AudioVoice = "algenib"
	VoiceAlgieba       AudioVoice = "algieba"
	VoiceAlnilam       AudioVoice = "alnilam"
	VoiceAoede         AudioVoice = "aoede"
	VoiceAutonoe       AudioVoice = "autonoe"
	VoiceCallirrhoe    AudioVoice = "callirrhoe"
	VoiceCharon        AudioVoice = "charon"
	VoiceDespina       AudioVoice = "despina"
	VoiceEnceladus     AudioVoice = "enceladus"
	VoiceErinome       AudioVoice = "erinome"
	VoiceFenrir        AudioVoice = "fenrir"
	VoiceGacrux        AudioVoice = "gacrux"
	VoiceIapetus       AudioVoice = "iapetus"
	VoiceKore          AudioVoice = "kore"
	VoiceLaomedeia     AudioVoice = "laomedeia"
	VoiceLeda          AudioVoice = "leda"
	VoiceOrus          AudioVoice = "orus"
	VoicePuck          AudioVoice = "puck"
	VoicePulcherrima   AudioVoice = "pulcherrima"
	VoiceRasalgethi    AudioVoice = "rasalgethi"
	VoiceSadachbia     AudioVoice = "sadachbia"
	VoiceSadaltager    AudioVoice = "sadaltager"
	VoiceSchedar       AudioVoice = "schedar"
	VoiceSulafat       AudioVoice = "sulafat"
	VoiceUmbriel       AudioVoice = "umbriel"
	VoiceVindemiatrix  AudioVoice = "vindemiatrix"
	VoiceZephyr        AudioVoice = "zephyr"
	VoiceZubenelgenubi AudioVoice = "zubenelgenubi"
)

// CreateAudioParams configures a reusable voice preset from one of the built-in [AudioVoice] identities.
// The returned audio ID can be attached to characters or text-to-video params to control narration voice.
// This is a synchronous operation (no Create/Get polling needed, use Run directly).
type CreateAudioParams struct {
	AudioID          AudioVoice `json:"audio_id" help:"required; preset voice ID such as achernar, achird, or zephyr"`
	Name             string     `json:"name" help:"required; voice name, max 210 chars"`
	VoiceDescription string     `json:"voice_description,omitempty" help:"optional; voice characteristic description, max 20000 chars"`
	ExampleDialogue  string     `json:"example_dialogue,omitempty" help:"optional; example dialogue, max 120 chars"`
}

// Audio holds the ID and name of a created voice preset.
type Audio struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// CreateAudioResponse is the result of a synchronous create-audio call.
type CreateAudioResponse struct {
	core.TaskBillingFacts
	ID    string `json:"id"`
	Audio *Audio `json:"audio,omitempty"`
	Error string `json:"error,omitempty"`
}

// CreateCharacterParams configures a reusable character from a required portrait,
// an optional full-body reference image, and a description.
// The returned character ID can be passed to [TextToVideoParams].CharacterIDs for consistent identity across videos.
// Optionally attach AudioIDs to give the character a specific voice. Run follows an accepted Task to completion.
type CreateCharacterParams struct {
	Descriptions          string   `json:"descriptions" help:"required; character appearance, identity, style, clothing, or personality description"`
	ReferenceImageURL     string   `json:"reference_image_url" help:"required; portrait reference image URL, max 20MB"`
	BodyReferenceImageURL string   `json:"body_reference_image_url,omitempty" help:"optional; full-body reference image URL, max 20MB; requires the portrait reference image"`
	AudioIDs              []string `json:"audio_ids,omitempty" help:"optional; audio IDs from create-audio to guide voice traits"`
	CharacterName         string   `json:"character_name,omitempty" help:"optional; character name, max 210 chars"`
}

// Image holds a URL to a generated or reference image.
type Image struct {
	URL string `json:"url"`
}

// Character holds the ID, name, and ordered reference images for a created character.
// Images contains the portrait first and the optional full-body image second.
type Character struct {
	ID     string  `json:"id"`
	Name   string  `json:"name,omitempty"`
	Images []Image `json:"images,omitempty"`
}

// CreateCharacterResponse is the terminal result of a create-character call.
type CreateCharacterResponse struct {
	core.TaskBillingFacts
	ID        string     `json:"id"`
	Character *Character `json:"character,omitempty"`
	Error     string     `json:"error,omitempty"`
}

// TextToVideoModel selects a Gemini Omni text-to-video model.
type TextToVideoModel string

const (
	// ModelGeminiOmniFlash11 supports prompt, multimodal reference, and first/last-frame video generation.
	ModelGeminiOmniFlash11 TextToVideoModel = generatedTextToVideoModelGeminiOmniFlash11
	// ModelGeminiOmniFlashPreview is the prompt-only 720p Flash Preview model.
	ModelGeminiOmniFlashPreview TextToVideoModel = generatedTextToVideoModelGeminiOmniFlashPreview
	// ModelGeminiOmniTextToVideo is the multimodal Gemini Omni model.
	ModelGeminiOmniTextToVideo TextToVideoModel = generatedTextToVideoModelGeminiOmniTextToVideo
)

// VideoClip defines a trimmed segment of a source video for use in text-to-video generation.
// The trimmed segment (Start to Ends) must be within 10 seconds.
type VideoClip struct {
	URL   string  `json:"url" help:"required; public source video URL"`
	Start float64 `json:"start" help:"required; trim start time in seconds, must be >= 0"`
	Ends  float64 `json:"ends" help:"required; trim end time in seconds, must be greater than start and within 10s"`
}

// TextToVideoParams configures multimodal video generation with optional characters, audio voices,
// reference images, and source video clips. Pre-create characters via [CreateCharacter] and
// audio voices via [CreateAudio], then reference their IDs here.
type TextToVideoParams struct {
	Model              TextToVideoModel `json:"model,omitempty" help:"optional; model slug; defaults to gemini-omni-text-to-video"`
	Prompt             string           `json:"prompt" help:"required; video prompt, max 20000 chars"`
	DurationSeconds    int              `json:"duration_seconds,omitempty" help:"required except for gemini-omni-flash-preview"`
	CallbackURL        string           `json:"callback_url,omitempty" help:"optional; HTTPS callback URL"`
	ReferenceImageURLs []string         `json:"reference_image_urls,omitempty" help:"optional; reference image URLs, max 7"`
	AudioIDs           []string         `json:"audio_ids,omitempty" help:"optional; audio IDs from create-audio, max 3"`
	VideoList          []VideoClip      `json:"video_list,omitempty" help:"optional; source video clips, max 1; each clip uses 2 reference units"`
	CharacterIDs       []string         `json:"character_ids,omitempty" help:"optional; character IDs from create-character, max 3; dual-image characters use 2 reference units"`
	FirstFrameImageURL string           `json:"first_frame_image_url,omitempty" help:"optional for gemini-omni-flash-1-1; public first-frame image URL; cannot be combined with reference_image_urls, audio_ids, video_list, or character_ids"`
	LastFrameImageURL  string           `json:"last_frame_image_url,omitempty" help:"optional for gemini-omni-flash-1-1; public last-frame image URL; requires first_frame_image_url"`
	AspectRatio        string           `json:"aspect_ratio,omitempty" help:"optional; output aspect ratio"`
	OutputResolution   string           `json:"output_resolution,omitempty" help:"optional; output resolution; default 720p"`
	Seed               *int             `json:"seed,omitempty" help:"optional; integer in [0, 2147483647]"`
}

// TaskResponse carries the task ID, lifecycle status, and error for Gemini Omni async operations.
type TaskResponse struct {
	core.TaskBillingFacts
	ID     string `json:"id"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func (r TaskResponse) GetID() string     { return r.ID }
func (r TaskResponse) GetStatus() string { return r.Status }
func (r TaskResponse) GetError() string  { return r.Error }

// Video holds a URL to a generated video file.
type Video struct {
	URL string `json:"url"`
}

// TextToVideoResponse is the completed result of an async text-to-video task.
type TextToVideoResponse struct {
	TaskResponse
	Videos []Video `json:"videos,omitempty"`
}
