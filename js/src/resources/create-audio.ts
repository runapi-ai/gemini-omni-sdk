import type { HttpClient, RequestOptions } from '@runapi.ai/core';
import { compactParams } from '@runapi.ai/core';
import type { CreateAudioParams, CreateAudioResponse } from '../types';

const ENDPOINT = '/api/v1/gemini_omni/create_audio';

export class CreateAudio {
  constructor(private readonly http: HttpClient) {}

  async run(params: CreateAudioParams, options?: RequestOptions): Promise<CreateAudioResponse> {
    return this.http.request<CreateAudioResponse>('POST', ENDPOINT, {
      body: compactParams(params),
      ...options,
    });
  }
}
