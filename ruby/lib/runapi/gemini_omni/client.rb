# frozen_string_literal: true

module RunApi
  module GeminiOmni
    # Gemini Omni multimodal generation client for voice presets, character creation,
    # and text-to-video with optional characters, audio voices, and reference media.
    #
    # @example
    #   client = RunApi::GeminiOmni::Client.new(api_key: "sk-...")
    #
    #   audio = client.create_audio.run(audio_id: "zephyr", name: "Narrator")
    #   video = client.text_to_video.run(
    #     prompt: "A narrator walks through a futuristic city",
    #     duration_seconds: 6,
    #     audio_ids: [audio.audio.id]
    #   )
    #   puts video.videos.first.url
    class Client < RunApi::Core::Client
      # @return [Resources::CreateAudio] Registers reusable voice presets (synchronous).
      attr_reader :create_audio
      # @return [Resources::CreateCharacter] Builds reusable characters from reference images (synchronous).
      attr_reader :create_character
      # @return [Resources::TextToVideo] Generates video from prompts with optional characters and voices (async).
      attr_reader :text_to_video

      def initialize(api_key: nil, **options)
        super
        @create_audio = Resources::CreateAudio.new(http)
        @create_character = Resources::CreateCharacter.new(http)
        @text_to_video = Resources::TextToVideo.new(http)
      end
    end
  end
end
