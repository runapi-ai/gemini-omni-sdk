import type { HttpClient, RequestOptions, ActionSchema } from '@runapi.ai/core';
import { compactParams, validateParams } from '@runapi.ai/core';
import { contract } from '../contract_gen';
import type { CreateAudioParams, CreateAudioResponse } from '../types';

const ENDPOINT = '/api/v1/gemini_omni/create_audio';

// Fixed endpoint model, injected only for contract validation (never sent on the wire).
const MODEL = 'gemini-omni-audio';

/**
 * Registers a reusable voice preset from a built-in voice identity.
 * This is a synchronous operation -- only `run()` is available (no create/get polling).
 */
export class CreateAudio {
  constructor(private readonly http: HttpClient) {}

  /**
   * Register a reusable voice preset (synchronous).
   * @param params Voice preset parameters.
   * @param options Per-request overrides.
   * @returns The created audio voice preset.
   */
  async run(params: CreateAudioParams, options?: RequestOptions): Promise<CreateAudioResponse> {
    const body = compactParams(params);
    validateParams(contract['create-audio'] as ActionSchema, { ...body, model: MODEL } as Record<string, unknown>);
    return this.http.request<CreateAudioResponse>('POST', ENDPOINT, {
      body,
      ...options,
    });
  }
}
