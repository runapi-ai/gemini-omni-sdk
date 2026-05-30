package geminiomni

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

type CreateAudioParams struct {
	AudioID          AudioVoice `json:"audio_id" help:"required; preset voice ID such as achernar, achird, or zephyr"`
	Name             string     `json:"name" help:"required; voice name, max 210 chars"`
	VoiceDescription string     `json:"voice_description,omitempty" help:"optional; voice characteristic description, max 20000 chars"`
	ExampleDialogue  string     `json:"example_dialogue,omitempty" help:"optional; example dialogue, max 120 chars"`
}

type Audio struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type CreateAudioResponse struct {
	ID    string `json:"id"`
	Audio *Audio `json:"audio,omitempty"`
	Error string `json:"error,omitempty"`
}

type CreateCharacterParams struct {
	Descriptions  string   `json:"descriptions" help:"required; character appearance, identity, style, clothing, or personality description"`
	ImageURLs     []string `json:"image_urls" help:"required; one public character reference image URL, max 20MB"`
	AudioIDs      []string `json:"audio_ids,omitempty" help:"optional; audio IDs from create-audio to guide voice traits"`
	CharacterName string   `json:"character_name,omitempty" help:"optional; character name, max 210 chars"`
}

type Character struct {
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}

type CreateCharacterResponse struct {
	ID        string     `json:"id"`
	Character *Character `json:"character,omitempty"`
	Error     string     `json:"error,omitempty"`
}

type VideoClip struct {
	URL   string  `json:"url" help:"required; public source video URL"`
	Start float64 `json:"start" help:"required; trim start time in seconds, must be >= 0"`
	Ends  float64 `json:"ends" help:"required; trim end time in seconds, must be greater than start and within 10s"`
}

type TextToVideoParams struct {
	Prompt       string      `json:"prompt" help:"required; video prompt, max 20000 chars"`
	Duration     string      `json:"duration" help:"required; duration"`
	CallbackURL  string      `json:"callback_url,omitempty" help:"optional; HTTPS callback URL"`
	ImageURLs    []string    `json:"image_urls,omitempty" help:"optional; public image reference URLs, max 7"`
	AudioIDs     []string    `json:"audio_ids,omitempty" help:"optional; audio IDs from create-audio, max 3"`
	VideoList    []VideoClip `json:"video_list,omitempty" help:"optional; source video clips, max 1; each clip uses 2 reference units"`
	CharacterIDs []string    `json:"character_ids,omitempty" help:"optional; character IDs from create-character, max 3"`
	AspectRatio  string      `json:"aspect_ratio,omitempty" help:"optional; output aspect ratio"`
	Resolution   string      `json:"resolution,omitempty" help:"optional; output resolution; default 720p"`
	Seed         *int        `json:"seed,omitempty" help:"optional; integer in [0, 2147483647]"`
}

type TaskResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func (r TaskResponse) GetID() string     { return r.ID }
func (r TaskResponse) GetStatus() string { return r.Status }
func (r TaskResponse) GetError() string  { return r.Error }

type Video struct {
	URL string `json:"url"`
}

type TextToVideoResponse struct {
	TaskResponse
	Videos []Video `json:"videos,omitempty"`
}
