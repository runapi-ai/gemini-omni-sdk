import { BaseClient, type ClientOptions } from '@runapi.ai/core';
import { CreateAudio } from './resources/create-audio';
import { CreateCharacter } from './resources/create-character';
import { TextToVideo } from './resources/text-to-video';

/**
 * Gemini Omni multimodal generation client for voice presets, character creation,
 * and text-to-video with optional characters, audio voices, and reference media.
 *
 * @example
 * ```typescript
 * import { GeminiOmniClient } from '@runapi.ai/gemini-omni';
 * const client = new GeminiOmniClient({ apiKey: 'sk-...' });
 *
 * // Create a voice preset, then generate a video using it
 * const audio = await client.createAudio.run({
 *   audio_id: 'zephyr',
 *   name: 'Narrator',
 * });
 * const video = await client.textToVideo.run({
 *   prompt: 'A narrator walks through a futuristic city',
 *   duration_seconds: 6,
 *   audio_ids: [audio.audio!.id],
 * });
 * console.log(video.videos[0].url);
 * ```
 */
export class GeminiOmniClient extends BaseClient {
  /** Registers a reusable voice preset from a built-in voice identity (synchronous). */
  public readonly createAudio: CreateAudio;
  /** Builds a reusable character from a reference image and description (synchronous). */
  public readonly createCharacter: CreateCharacter;
  /** Generates video from a prompt with optional characters, voices, images, and clips (async with polling). */
  public readonly textToVideo: TextToVideo;

  constructor(options: ClientOptions = {}) {
    super(options);
    this.createAudio = new CreateAudio(this.http);
    this.createCharacter = new CreateCharacter(this.http);
    this.textToVideo = new TextToVideo(this.http);
  }
}
