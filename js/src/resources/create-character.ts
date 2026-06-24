import type { HttpClient, RequestOptions, ActionSchema } from '@runapi.ai/core';
import { compactParams, validateParams } from '@runapi.ai/core';
import { contract } from '../contract_gen';
import type { CreateCharacterParams, CreateCharacterResponse } from '../types';

const ENDPOINT = '/api/v1/gemini_omni/create_character';

// Fixed endpoint model, injected only for contract validation (never sent on the wire).
const MODEL = 'gemini-omni-character';

/**
 * Builds a reusable character from a reference image and description.
 * Attach audio IDs to give the character a specific voice.
 * This is a synchronous operation -- only `run()` is available (no create/get polling).
 */
export class CreateCharacter {
  constructor(private readonly http: HttpClient) {}

  /**
   * Build a reusable character (synchronous).
   * @param params Character creation parameters.
   * @param options Per-request overrides.
   * @returns The created character.
   */
  async run(params: CreateCharacterParams, options?: RequestOptions): Promise<CreateCharacterResponse> {
    const body = compactParams(params);
    validateParams(contract['create-character'] as ActionSchema, { ...body, model: MODEL } as Record<string, unknown>);
    return this.http.request<CreateCharacterResponse>('POST', ENDPOINT, {
      body,
      ...options,
    });
  }
}
