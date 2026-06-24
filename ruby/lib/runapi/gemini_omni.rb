# frozen_string_literal: true

require "runapi/core"
require_relative "gemini_omni/types"
require_relative "gemini_omni/contract_gen"
require_relative "gemini_omni/resources/create_audio"
require_relative "gemini_omni/resources/create_character"
require_relative "gemini_omni/resources/text_to_video"
require_relative "gemini_omni/client"

module RunApi
  module GeminiOmni
    AuthenticationError = RunApi::Core::AuthenticationError
    RateLimitError = RunApi::Core::RateLimitError
    InsufficientCreditsError = RunApi::Core::InsufficientCreditsError
    NotFoundError = RunApi::Core::NotFoundError
    ValidationError = RunApi::Core::ValidationError
  end
end
