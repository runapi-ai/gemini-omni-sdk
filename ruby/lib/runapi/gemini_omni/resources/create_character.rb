# frozen_string_literal: true

module RunApi
  module GeminiOmni
    module Resources
      # Builds a reusable character from a portrait, optional full-body reference, and description.
      # Attach audio IDs to give the character a specific voice.
      # +run+ returns the terminal character result, following an accepted Task when needed.
      class CreateCharacter
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/gemini_omni/create_character"
        RESPONSE_CLASS = Types::CreateCharacterResponse
        MODEL = "gemini-omni-character"

        def initialize(http)
          @http = http
        end

        def run(options: nil, **params)
          params = compact_params(params)
          run_hybrid(ENDPOINT, body: params, options: options, response_class: RESPONSE_CLASS)
        end
      end
    end
  end
end
