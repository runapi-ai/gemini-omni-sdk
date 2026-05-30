import { beforeEach, describe, expect, it, vi } from 'vitest';
import type { HttpClient } from '@runapi.ai/core';
import { CreateAudio } from '../../src/resources/create-audio';
import { CreateCharacter } from '../../src/resources/create-character';
import { TextToVideo } from '../../src/resources/text-to-video';

describe('Gemini Omni resources', () => {
  const mockHttp: HttpClient = {
    request: vi.fn(),
  };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('creates audio voices with direct service params', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({
      id: 'audio-kie-123',
      audio: { id: 'audio-kie-123', name: 'Acher Narrator' },
    });
    const createAudio = new CreateAudio(mockHttp);

    const result = await createAudio.run({
      audio_id: 'achernar',
      name: 'Acher Narrator',
      voice_description: 'A calm, clear voice',
      example_dialogue: 'Hello, I am achernar',
    });

    expect(mockHttp.request).toHaveBeenCalledWith('POST', '/api/v1/gemini_omni/create_audio', {
      body: {
        audio_id: 'achernar',
        name: 'Acher Narrator',
        voice_description: 'A calm, clear voice',
        example_dialogue: 'Hello, I am achernar',
      },
    });
    expect(result.id).toBe('audio-kie-123');
  });

  it('creates characters with direct service params', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({
      id: 'character-kie-123',
      character: { id: 'character-kie-123', name: 'Jenny', image_url: 'https://file.runapi.ai/gemini/jenny.png' },
    });
    const createCharacter = new CreateCharacter(mockHttp);

    const result = await createCharacter.run({
      descriptions: 'A silver-haired cyberpunk guide',
      image_urls: ['https://file.runapi.ai/demo/character.png'],
      audio_ids: ['audio-kie-123'],
      character_name: 'Jenny',
    });

    expect(mockHttp.request).toHaveBeenCalledWith('POST', '/api/v1/gemini_omni/create_character', {
      body: {
        descriptions: 'A silver-haired cyberpunk guide',
        image_urls: ['https://file.runapi.ai/demo/character.png'],
        audio_ids: ['audio-kie-123'],
        character_name: 'Jenny',
      },
    });
    expect(result.id).toBe('character-kie-123');
  });

  it('creates text-to-video tasks with direct service params', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({
      id: 'task-local-123',
      status: 'processing',
    });
    const textToVideo = new TextToVideo(mockHttp);

    const result = await textToVideo.create({
      prompt: 'Create a neon city tracking shot',
      duration: '8',
      aspect_ratio: '16:9',
      resolution: '1080p',
      image_urls: ['https://file.runapi.ai/demo/scene.png'],
      audio_ids: ['audio-kie-123'],
      character_ids: ['character-kie-123'],
    });

    expect(mockHttp.request).toHaveBeenCalledWith('POST', '/api/v1/gemini_omni/text_to_video', {
      body: {
        prompt: 'Create a neon city tracking shot',
        duration: '8',
        aspect_ratio: '16:9',
        resolution: '1080p',
        image_urls: ['https://file.runapi.ai/demo/scene.png'],
        audio_ids: ['audio-kie-123'],
        character_ids: ['character-kie-123'],
      },
    });
    expect(result.id).toBe('task-local-123');
  });

  it('gets text-to-video task status', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({
      id: 'task-local-123',
      status: 'completed',
      videos: [{ url: 'https://tempfile.runapi.ai/gemini/output.mp4' }],
    });
    const textToVideo = new TextToVideo(mockHttp);

    const result = await textToVideo.get('task-local-123');

    expect(mockHttp.request).toHaveBeenCalledWith('GET', '/api/v1/gemini_omni/text_to_video/task-local-123', {});
    expect(result.videos?.[0]?.url).toBe('https://tempfile.runapi.ai/gemini/output.mp4');
  });
});
