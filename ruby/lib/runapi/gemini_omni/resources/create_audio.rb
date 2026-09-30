# frozen_string_literal: true

module RunApi
  module GeminiOmni
    module Resources
      # Registers a reusable voice preset from a built-in voice identity.
      # Synchronous -- only +run+ is available (no create/get polling).
      class CreateAudio
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/gemini_omni/create_audio"
        RESPONSE_CLASS = Types::CreateAudioResponse
        MODEL = "gemini-omni-audio"

        def initialize(http)
          @http = http
        end

        def run(options: nil, **params)
          params = compact_params(params)
          request(:post, ENDPOINT, body: params, options: options)
        end
      end
    end
  end
end
