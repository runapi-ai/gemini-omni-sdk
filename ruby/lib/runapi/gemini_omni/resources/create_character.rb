# frozen_string_literal: true

module RunApi
  module GeminiOmni
    module Resources
      class CreateCharacter
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/gemini_omni/create_character"
        RESPONSE_CLASS = Types::CreateCharacterResponse
        DESCRIPTIONS_MAX_LENGTH = 20_000
        CHARACTER_NAME_MAX_LENGTH = 210
        IMAGE_URLS_MAX = 1

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
          validate_required!(params, :descriptions)
          validate_required!(params, :image_urls)
          validate_array!(params, :image_urls)
          validate_array!(params, :audio_ids) if param(params, :audio_ids)
          validate_image_count!(params)
          validate_length!(params, :descriptions, DESCRIPTIONS_MAX_LENGTH)
          validate_length!(params, :character_name, CHARACTER_NAME_MAX_LENGTH)
        end

        def validate_required!(params, key)
          value = param(params, key)
          present = if value.is_a?(Array)
            value.any?
          else
            value.is_a?(String) ? !value.empty? : !value.nil?
          end
          return if present

          raise Core::ValidationError, "#{key} is required"
        end

        def validate_array!(params, key)
          return if param(params, key).is_a?(Array)

          raise Core::ValidationError, "#{key} must be an array"
        end

        def validate_image_count!(params)
          return if param(params, :image_urls).length <= IMAGE_URLS_MAX

          raise Core::ValidationError, "image_urls accepts at most #{IMAGE_URLS_MAX} image"
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
