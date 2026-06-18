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
        NAME_MAX_LENGTH = 210
        VOICE_DESCRIPTION_MAX_LENGTH = 20_000
        EXAMPLE_DIALOGUE_MAX_LENGTH = 120

        def initialize(http)
          @http = http
        end

        def run(**params)
          params = compact_params(params)
          validate_params!(params)
          request(:post, ENDPOINT, body: params)
        end

        private

        def validate_params!(params)
          validate_required!(params, :audio_id)
          validate_required!(params, :name)
          validate_length!(params, :name, NAME_MAX_LENGTH)
          validate_length!(params, :voice_description, VOICE_DESCRIPTION_MAX_LENGTH)
          validate_length!(params, :example_dialogue, EXAMPLE_DIALOGUE_MAX_LENGTH)
        end

        def validate_required!(params, key)
          value = param(params, key)
          return if value.is_a?(String) ? !value.empty? : !value.nil?

          raise Core::ValidationError, "#{key} is required"
        end

        def validate_length!(params, key, max_length)
          value = param(params, key)
          return if value.nil? || value.to_s.length <= max_length

          raise Core::ValidationError, "#{key} must be at most #{max_length} characters"
        end
      end
    end
  end
end
