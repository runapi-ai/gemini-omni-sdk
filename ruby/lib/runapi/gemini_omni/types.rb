# frozen_string_literal: true

module RunApi
  module GeminiOmni
    module Types
      AUDIO_VOICES = %w[
        achernar achird algenib algieba alnilam aoede autonoe callirrhoe charon
        despina enceladus erinome fenrir gacrux iapetus kore laomedeia leda orus
        puck pulcherrima rasalgethi sadachbia sadaltager schedar sulafat umbriel
        vindemiatrix zephyr zubenelgenubi
      ].freeze

      class Audio < RunApi::Core::BaseModel
        required :id, String
        optional :name, String
      end

      class CreateAudioResponse < RunApi::Core::BaseModel
        required :id, String
        optional :audio, -> { Audio }
        optional :error, String
      end

      class Character < RunApi::Core::BaseModel
        required :id, String
        optional :name, String
        optional :image_url, String
      end

      class CreateCharacterResponse < RunApi::Core::BaseModel
        required :id, String
        optional :character, -> { Character }
        optional :error, String
      end

      DURATIONS = %w[4 6 8 10].freeze
      ASPECT_RATIOS = %w[16:9 9:16].freeze
      RESOLUTIONS = %w[720p 1080p 4k].freeze
      SEED_RANGE = (0..2_147_483_647)

      class Video < RunApi::Core::BaseModel
        optional :url, String
      end

      class TextToVideoResponse < RunApi::Core::TaskResponse
        optional :videos, [-> { Video }]
      end

      class CompletedTextToVideoResponse < TextToVideoResponse
        required :videos, [-> { Video }]
      end
    end
  end
end
