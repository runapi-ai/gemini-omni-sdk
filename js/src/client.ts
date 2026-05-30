import { createHttpClient, type ClientOptions } from '@runapi.ai/core';
import { CreateAudio } from './resources/create-audio';
import { CreateCharacter } from './resources/create-character';
import { TextToVideo } from './resources/text-to-video';

export class GeminiOmniClient {
  public readonly createAudio: CreateAudio;
  public readonly createCharacter: CreateCharacter;
  public readonly textToVideo: TextToVideo;

  constructor(options: ClientOptions = {}) {
    const http = createHttpClient(options);
    this.createAudio = new CreateAudio(http);
    this.createCharacter = new CreateCharacter(http);
    this.textToVideo = new TextToVideo(http);
  }
}
