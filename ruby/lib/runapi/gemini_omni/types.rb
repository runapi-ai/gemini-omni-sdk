# frozen_string_literal: true

module RunApi
  module GeminiOmni
    # Type definitions and constants for the Gemini Omni API.
    module Types
      # The 30 preset voice identities, each with distinct pitch, cadence, and personality.
      AUDIO_VOICES = %w[
        achernar achird algenib algieba alnilam aoede autonoe callirrhoe charon
        despina enceladus erinome fenrir gacrux iapetus kore laomedeia leda orus
        puck pulcherrima rasalgethi sadachbia sadaltager schedar sulafat umbriel
        vindemiatrix zephyr zubenelgenubi
      ].freeze

      # A created voice preset with its server-assigned ID.
      class Audio < RunApi::Core::BaseModel
        required :id, String
        optional :name, String
      end

      # Result of a synchronous create-audio call.
      class CreateAudioResponse < RunApi::Core::BaseModel
        required :id, String
        optional :audio, -> { Audio }
        optional :error, String
      end

      # URL to a generated or reference image.
      class Image < RunApi::Core::BaseModel
        optional :url, String
      end

      # A created character with its ID, name, and reference images.
      class Character < RunApi::Core::BaseModel
        required :id, String
        optional :name, String
        optional :images, [-> { Image }]
      end

      # Result of a synchronous create-character call.
      class CreateCharacterResponse < RunApi::Core::BaseModel
        required :id, String
        optional :character, -> { Character }
        optional :error, String
      end

      # Valid seed range for reproducible generation.
      SEED_RANGE = (0..2_147_483_647)

      # URL to a generated video file.
      class Video < RunApi::Core::BaseModel
        optional :url, String
      end

      # Async text-to-video task result with lifecycle status and generated videos.
      class TextToVideoResponse < RunApi::Core::TaskResponse
        optional :videos, [-> { Video }]
      end

      # Narrowed response returned by +run+ once polling confirms completion.
      # Videos are guaranteed present.
      class CompletedTextToVideoResponse < TextToVideoResponse
        required :videos, [-> { Video }]
      end
    end
  end
end
