import type { AsyncTaskStatus } from '@runapi.ai/core';

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

export interface CreateAudioParams {
  audio_id: GeminiOmniAudioVoice | string;
  name: string;
  voice_description?: string;
  example_dialogue?: string;
}

export interface GeminiOmniAudio {
  id: string;
  name?: string;
  [key: string]: unknown;
}

export interface CreateAudioResponse {
  id: string;
  audio?: GeminiOmniAudio;
  error?: string;
  [key: string]: unknown;
}

export interface CreateCharacterParams {
  descriptions: string;
  reference_image_url: string;
  audio_ids?: string[];
  character_name?: string;
}

export interface ImageMetadata {
  url: string;
  [key: string]: unknown;
}

export interface GeminiOmniCharacter {
  id: string;
  name?: string;
  images?: ImageMetadata[];
  [key: string]: unknown;
}

export interface CreateCharacterResponse {
  id: string;
  character?: GeminiOmniCharacter;
  error?: string;
  [key: string]: unknown;
}

export type GeminiOmniTextToVideoDuration = 4 | 6 | 8 | 10;
export type GeminiOmniTextToVideoAspectRatio = '16:9' | '9:16';
export type GeminiOmniTextToVideoResolution = '720p' | '1080p' | '4k';

export interface GeminiOmniTextToVideoClip {
  url: string;
  start: number;
  ends: number;
}

export interface TextToVideoParams {
  prompt: string;
  duration_seconds: GeminiOmniTextToVideoDuration;
  callback_url?: string;
  reference_image_urls?: string[];
  audio_ids?: string[];
  video_list?: GeminiOmniTextToVideoClip[];
  character_ids?: string[];
  aspect_ratio?: GeminiOmniTextToVideoAspectRatio;
  output_resolution?: GeminiOmniTextToVideoResolution;
  seed?: number;
}

export interface TaskCreateResponse {
  id: string;
  status?: AsyncTaskStatus;
}

export interface VideoMetadata {
  url: string;
}

export interface TextToVideoResponse {
  id: string;
  status: AsyncTaskStatus;
  videos?: VideoMetadata[];
  error?: string;
  [key: string]: unknown;
}

export type CompletedTextToVideoResponse = TextToVideoResponse & {
  status: 'completed';
  videos: VideoMetadata[];
};
