import type { HttpClient, HybridTaskOptions, ActionSchema } from '@runapi.ai/core';
import { compactParams, createHybridTask, validateParams } from '@runapi.ai/core';
import { contract } from '../contract_gen';
import type { CreateCharacterParams, CreateCharacterResponse } from '../types';

const ENDPOINT = '/api/v1/gemini_omni/create_character';

// Fixed endpoint model, injected only for contract validation (never sent on the wire).
const MODEL = 'gemini-omni-character';

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
    validateParams(contract['create-character'] as ActionSchema, { ...body, model: MODEL } as Record<string, unknown>);
    return (await createHybridTask<CreateCharacterResponse>(this.http, ENDPOINT, { body, ...options })).run();
  }
}
