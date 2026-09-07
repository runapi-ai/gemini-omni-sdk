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
            "example_dialogue" => {
              "max" => 120,
              "length" => true
            },
            "name" => {
              "required" => true,
              "max" => 210,
              "length" => true
            },
            "voice_description" => {
              "max" => 20000,
              "length" => true
            }
          }
        }
      },
      "create-character" => {
        "models" => ["gemini-omni-character"],
        "fields_by_model" => {
          "gemini-omni-character" => {
            "character_name" => {
              "max" => 210,
              "length" => true
            },
            "descriptions" => {
              "required" => true,
              "max" => 20000,
              "length" => true
            },
            "reference_image_url" => {
              "required" => true
            }
          }
        }
      },
      "text-to-video" => {
        "models" => ["gemini-omni-flash-1-1", "gemini-omni-flash-preview", "gemini-omni-text-to-video"],
        "fields_by_model" => {
          "gemini-omni-flash-1-1" => {
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
              "enum" => ["360p", "720p", "1080p", "4k"]
            },
            "prompt" => {
              "required" => true,
              "max" => 20000,
              "length" => true
            },
            "reference_image_urls" => {
              "max_items" => 7
            },
            "seed" => {
              "min" => 0,
              "max" => 2147483647,
              "type" => "integer"
            },
            "video_list" => {
              "max_items" => 1
            }
          },
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
              "required" => true,
              "max" => 20000,
              "length" => true
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
              "required" => true,
              "max" => 20000,
              "length" => true
            },
            "reference_image_urls" => {
              "max_items" => 7
            },
            "seed" => {
              "min" => 0,
              "max" => 2147483647,
              "type" => "integer"
            },
            "video_list" => {
              "max_items" => 1
            }
          }
        },
        "rules" => [{
          "when" => {
            "model" => "gemini-omni-flash-1-1",
            "first_frame_image_url" => {
              "present" => true
            }
          },
          "forbidden" => ["reference_image_urls", "audio_ids", "video_list", "character_ids"]
        }, {
          "when" => {
            "model" => "gemini-omni-flash-1-1",
            "last_frame_image_url" => {
              "present" => true
            }
          },
          "required" => ["first_frame_image_url"]
        }, {
          "when" => {
            "model" => "gemini-omni-flash-preview"
          },
          "forbidden" => ["reference_image_urls", "audio_ids", "video_list", "character_ids", "first_frame_image_url", "last_frame_image_url", "duration_seconds", "seed"]
        }, {
          "when" => {
            "model" => "gemini-omni-text-to-video"
          },
          "forbidden" => ["first_frame_image_url", "last_frame_image_url"]
        }]
      }
    }.freeze
  end
end
