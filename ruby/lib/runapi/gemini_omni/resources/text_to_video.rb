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
        DEFAULT_MODEL = "gemini-omni-text-to-video"

        def initialize(http)
          @http = http
        end

        def run(options: nil, **params)
          task = create(options: options, **params)
          poll_until_complete { get(task.id, options: options) }
        end

        def create(options: nil, **params)
          params = compact_params(params)
          request(:post, ENDPOINT, body: params, options: options)
        end

        def get(id, options: nil)
          request(:get, "#{ENDPOINT}/#{id}", options: options)
        end
      end
    end
  end
end
