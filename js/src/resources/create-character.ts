import type { HttpClient, RequestOptions } from '@runapi.ai/core';
import { compactParams } from '@runapi.ai/core';
import type { CreateCharacterParams, CreateCharacterResponse } from '../types';

const ENDPOINT = '/api/v1/gemini_omni/create_character';

export class CreateCharacter {
  constructor(private readonly http: HttpClient) {}

  async run(params: CreateCharacterParams, options?: RequestOptions): Promise<CreateCharacterResponse> {
    return this.http.request<CreateCharacterResponse>('POST', ENDPOINT, {
      body: compactParams(params),
      ...options,
    });
  }
}
