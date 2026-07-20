# frozen_string_literal: true

module RunApi
  module GeminiOmni
    CONTRACT = {
      "create-audio" => {
        "models" => ["gemini-omni-audio"],
        "fields_by_model" => {
          "gemini-omni-audio" => {
            "audio_id" => {
              "required" => true
            },
            "name" => {
              "required" => true
            }
          }
        }
      },
      "create-character" => {
        "models" => ["gemini-omni-character"],
        "fields_by_model" => {
          "gemini-omni-character" => {
            "descriptions" => {
              "required" => true
            },
            "reference_image_url" => {
              "required" => true
            }
          }
        }
      },
      "text-to-video" => {
        "models" => ["gemini-omni-flash-preview", "gemini-omni-text-to-video"],
        "fields_by_model" => {
          "gemini-omni-flash-preview" => {
            "aspect_ratio" => {
              "enum" => ["16:9", "9:16"]
            },
            "duration_seconds" => {
              "type" => "integer"
            },
            "output_resolution" => {
              "enum" => ["720p"]
            },
            "prompt" => {
              "required" => true
            },
            "seed" => {
              "type" => "integer"
            }
          },
          "gemini-omni-text-to-video" => {
            "aspect_ratio" => {
              "enum" => ["16:9", "9:16"]
            },
            "audio_ids" => {
              "max_items" => 3
            },
            "character_ids" => {
              "max_items" => 3
            },
            "duration_seconds" => {
              "enum" => [4, 6, 8, 10],
              "required" => true,
              "type" => "integer"
            },
            "output_resolution" => {
              "enum" => ["720p", "1080p", "4k"]
            },
            "prompt" => {
              "required" => true
            },
            "reference_image_urls" => {
              "max_items" => 7
            },
            "seed" => {
              "type" => "integer"
            },
            "video_list" => {
              "max_items" => 1
            }
          }
        },
        "rules" => [{
          "when" => {
            "model" => "gemini-omni-flash-preview"
          },
          "forbidden" => ["reference_image_urls", "audio_ids", "video_list", "character_ids", "duration_seconds", "seed"]
        }]
      }
    }.freeze
  end
end
