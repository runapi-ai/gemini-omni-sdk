import type { HttpClient, RequestOptions } from '@runapi.ai/core';
import { compactParams } from '@runapi.ai/core';
import type { CreateAudioParams, CreateAudioResponse } from '../types';

const ENDPOINT = '/api/v1/gemini_omni/create_audio';

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
    return this.http.request<CreateAudioResponse>('POST', ENDPOINT, {
      body: compactParams(params),
      ...options,
    });
  }
}
