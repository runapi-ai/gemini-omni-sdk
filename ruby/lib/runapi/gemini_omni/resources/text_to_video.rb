# frozen_string_literal: true

module RunApi
  module GeminiOmni
    module Resources
      # Generates video from a prompt with optional characters, audio voices,
      # reference images, and source video clips.
      # Async -- use +run+ for automatic polling or +create+/+get+ for manual control.
      class TextToVideo
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/gemini_omni/text_to_video"
        RESPONSE_CLASS = Types::TextToVideoResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedTextToVideoResponse
        MODEL = "gemini-omni-text-to-video"
        PROMPT_MAX_LENGTH = 20_000
        REFERENCE_IMAGE_URLS_MAX = 7
        AUDIO_IDS_MAX = 3
        VIDEO_LIST_MAX = 1
        CHARACTER_IDS_MAX = 3
        REFERENCE_UNITS_MAX = 7
        VIDEO_REFERENCE_UNITS = 2
        MAX_TRIM_SECONDS = 10

        def initialize(http)
          @http = http
        end

        def run(**params)
          task = create(**params)
          poll_until_complete { get(task.id) }
        end

        def create(**params)
          params = compact_params(params)
          validate_params!(params)
          request(:post, ENDPOINT, body: params)
        end

        def get(id)
          request(:get, "#{ENDPOINT}/#{id}")
        end

        private

        def validate_params!(params)
          validate_contract!(CONTRACT["text-to-video"], params.merge(model: MODEL))
          validate_length!(params, :prompt, PROMPT_MAX_LENGTH)
          validate_array!(params, :reference_image_urls, REFERENCE_IMAGE_URLS_MAX) if param(params, :reference_image_urls)
          validate_array!(params, :audio_ids, AUDIO_IDS_MAX) if param(params, :audio_ids)
          validate_array!(params, :video_list, VIDEO_LIST_MAX) if param(params, :video_list)
          validate_array!(params, :character_ids, CHARACTER_IDS_MAX) if param(params, :character_ids)
          validate_video_list!(param(params, :video_list)) if param(params, :video_list)
          validate_reference_units!(params)
          validate_seed!(params)
        end

        def validate_length!(params, key, max_length)
          value = param(params, key)
          return if value.nil? || value.to_s.length <= max_length

          raise Core::ValidationError, "#{key} must be at most #{max_length} characters"
        end

        def validate_array!(params, key, max_length)
          value = param(params, key)
          raise Core::ValidationError, "#{key} must be an array" unless value.is_a?(Array)
          return if value.length <= max_length

          raise Core::ValidationError, "#{key} accepts at most #{max_length} items"
        end

        def validate_video_list!(items)
          items.each_with_index do |item, index|
            url = param(item, :url)
            raise Core::ValidationError, "video_list[#{index}].url is required" if url.nil? || url.to_s.empty?

            start_time = numeric_param(item, :start)
            end_time = numeric_param(item, :ends)
            raise Core::ValidationError, "video_list[#{index}] start and ends must be numbers" unless start_time && end_time
            raise Core::ValidationError, "video_list[#{index}].start must be 0 or greater" if start_time.negative?
            raise Core::ValidationError, "video_list[#{index}].ends must be greater than start" unless end_time > start_time
            if (end_time - start_time) > MAX_TRIM_SECONDS
              raise Core::ValidationError, "video_list[#{index}] trim range must be #{MAX_TRIM_SECONDS} seconds or less"
            end
          end
        end

        def validate_reference_units!(params)
          units = Array(param(params, :reference_image_urls)).count +
            (Array(param(params, :video_list)).count * VIDEO_REFERENCE_UNITS) +
            Array(param(params, :character_ids)).count
          return if units <= REFERENCE_UNITS_MAX

          raise Core::ValidationError, "reference_image_urls + video_list*2 + character_ids must use #{REFERENCE_UNITS_MAX} reference units or fewer"
        end

        def validate_seed!(params)
          value = param(params, :seed)
          return if value.nil?

          seed = Integer(value, exception: false)
          return if seed && Types::SEED_RANGE.cover?(seed)

          raise Core::ValidationError, "seed must be an integer between #{Types::SEED_RANGE.min} and #{Types::SEED_RANGE.max}"
        end

        def numeric_param(params, key)
          Float(param(params, key), exception: false)
        end
      end
    end
  end
end
