import type { AsyncTaskStatus } from '@runapi.ai/core';

/**
 * One of 30 preset voice identities for audio creation.
 * Each voice has a distinct pitch, cadence, and personality.
 * Pass the chosen value to {@link CreateAudioParams.audio_id}.
 */
export type GeminiOmniAudioVoice =
  | 'achernar'
  | 'achird'
  | 'algenib'
  | 'algieba'
  | 'alnilam'
  | 'aoede'
  | 'autonoe'
  | 'callirrhoe'
  | 'charon'
  | 'despina'
  | 'enceladus'
  | 'erinome'
  | 'fenrir'
  | 'gacrux'
  | 'iapetus'
  | 'kore'
  | 'laomedeia'
  | 'leda'
  | 'orus'
  | 'puck'
  | 'pulcherrima'
  | 'rasalgethi'
  | 'sadachbia'
  | 'sadaltager'
  | 'schedar'
  | 'sulafat'
  | 'umbriel'
  | 'vindemiatrix'
  | 'zephyr'
  | 'zubenelgenubi';

/**
 * Parameters for registering a reusable voice preset.
 * The returned audio ID can be attached to characters or text-to-video params
 * to control narration voice.
 */
export interface CreateAudioParams {
  /** Preset voice identity. See {@link GeminiOmniAudioVoice} for the full list. */
  audio_id: GeminiOmniAudioVoice | string;
  /** Display name for the voice preset, max 210 characters. */
  name: string;
  /** Describes vocal characteristics (pitch, tone, accent), max 20 000 characters. */
  voice_description?: string;
  /** Example dialogue line demonstrating the voice style, max 120 characters. */
  example_dialogue?: string;
}

/** A created voice preset with its server-assigned ID. */
export interface GeminiOmniAudio {
  id: string;
  name?: string;
  [key: string]: unknown;
}

/** Result of a synchronous create-audio call. */
export interface CreateAudioResponse {
  id: string;
  /** The created voice preset; present on success. */
  audio?: GeminiOmniAudio;
  error?: string;
  [key: string]: unknown;
}

/**
 * Parameters for building a reusable character from a reference image and description.
 * The returned character ID can be passed to {@link TextToVideoParams.character_ids}
 * for consistent identity across videos.
 */
export interface CreateCharacterParams {
  /** Appearance, identity, style, clothing, or personality description. */
  descriptions: string;
  /** Character reference image URL, max 20 MB. */
  reference_image_url: string;
  /** Audio IDs from create-audio to give the character a specific voice. */
  audio_ids?: string[];
  /** Character display name, max 210 characters. */
  character_name?: string;
}

/** URL to a generated or reference image. */
export interface ImageMetadata {
  url: string;
  [key: string]: unknown;
}

/** A created character with its reference images. */
export interface GeminiOmniCharacter {
  id: string;
  name?: string;
  /** Reference images associated with this character. */
  images?: ImageMetadata[];
  [key: string]: unknown;
}

/** Result of a synchronous create-character call. */
export interface CreateCharacterResponse {
  id: string;
  /** The created character; present on success. */
  character?: GeminiOmniCharacter;
  error?: string;
  [key: string]: unknown;
}

/** Video duration in seconds. Longer durations consume more credits. */
export type GeminiOmniTextToVideoDuration = 4 | 6 | 8 | 10;
/** Output aspect ratio -- landscape (16:9) or portrait (9:16). */
export type GeminiOmniTextToVideoAspectRatio = '16:9' | '9:16';
/** Output resolution -- higher resolutions produce sharper video at higher cost. */
export type GeminiOmniTextToVideoResolution = '720p' | '1080p' | '4k';

/**
 * A trimmed segment of a source video for use in text-to-video generation.
 * The trimmed segment (start to ends) must be within 10 seconds.
 * Each clip consumes 2 reference units toward the 7-unit limit.
 */
export interface GeminiOmniTextToVideoClip {
  /** Public source video URL. */
  url: string;
  /** Trim start time in seconds (>= 0). */
  start: number;
  /** Trim end time in seconds (must be > start, within 10 s of start). */
  ends: number;
}

/**
 * Parameters for multimodal video generation.
 * Pre-create characters via `createCharacter` and audio voices via `createAudio`,
 * then reference their IDs here.
 *
 * Reference units are shared: images (1 each) + video clips (2 each) + characters (1 each)
 * must total 7 or fewer.
 */
export interface TextToVideoParams {
  /** Video generation prompt, max 20 000 characters. */
  prompt: string;
  /** Output video duration in seconds. */
  duration_seconds: GeminiOmniTextToVideoDuration;
  /** HTTPS callback URL for task completion notification. */
  callback_url?: string;
  /** Reference image URLs, max 7. Each consumes 1 reference unit. */
  reference_image_urls?: string[];
  /** Audio IDs from create-audio for narration voices, max 3. */
  audio_ids?: string[];
  /** Source video clips for motion reference, max 1. Each consumes 2 reference units. */
  video_list?: GeminiOmniTextToVideoClip[];
  /** Character IDs from create-character for consistent identity, max 3. Each consumes 1 reference unit. */
  character_ids?: string[];
  aspect_ratio?: GeminiOmniTextToVideoAspectRatio;
  /** Defaults to 720p. */
  output_resolution?: GeminiOmniTextToVideoResolution;
  /** Reproducibility seed in [0, 2 147 483 647]. */
  seed?: number;
}

/** Acknowledgement returned by `create()` before the task starts processing. */
export interface TaskCreateResponse {
  id: string;
  status?: AsyncTaskStatus;
}

/** URL to a generated video file. */
export interface VideoMetadata {
  url: string;
}

/** Async text-to-video task result with lifecycle status. */
export interface TextToVideoResponse {
  id: string;
  status: AsyncTaskStatus;
  /** Generated video files; populated once the task completes. */
  videos?: VideoMetadata[];
  error?: string;
  [key: string]: unknown;
}

/** Narrowed response returned by `run()` once polling confirms completion. Videos are guaranteed present. */
export type CompletedTextToVideoResponse = TextToVideoResponse & {
  status: 'completed';
  videos: VideoMetadata[];
};
