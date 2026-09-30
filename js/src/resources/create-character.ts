import type { HttpClient, HybridTaskOptions } from '@runapi.ai/core';
import { compactParams, createHybridTask } from '@runapi.ai/core';
import type { CreateCharacterParams, CreateCharacterResponse } from '../types';

const ENDPOINT = '/api/v1/gemini_omni/create_character';

/**
 * Builds a reusable character from a portrait, optional full-body reference, and description.
 * Attach audio IDs to give the character a specific voice.
 * `run()` waits for the terminal character result when the request is accepted as a Task.
 */
export class CreateCharacter {
  constructor(private readonly http: HttpClient) {}

  /**
   * Build a reusable character and wait for its terminal result.
   * @param params Character creation parameters.
   * @param options Per-request overrides.
   * @returns The created character.
   */
  async run(params: CreateCharacterParams, options?: HybridTaskOptions): Promise<CreateCharacterResponse> {
    const body = compactParams(params);
    return (await createHybridTask<CreateCharacterResponse>(this.http, ENDPOINT, { body, ...options })).run();
  }
}
