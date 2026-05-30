# frozen_string_literal: true

module RunApi
  module GeminiOmni
    class Client
      attr_reader :create_audio, :create_character, :text_to_video

      def initialize(api_key: nil, **options)
        @api_key = Core::Auth.resolve_api_key(api_key)

        client_options = Core::ClientOptions.new(api_key: @api_key, **options)
        http = client_options.http_client || Core::HttpClient.new(client_options)
        @create_audio = Resources::CreateAudio.new(http)
        @create_character = Resources::CreateCharacter.new(http)
        @text_to_video = Resources::TextToVideo.new(http)
      end
    end
  end
end
