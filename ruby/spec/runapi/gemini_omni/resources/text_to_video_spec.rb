# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::GeminiOmni::Resources::TextToVideo do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:resource) { described_class.new(http) }
  let(:endpoint) { "/api/v1/gemini_omni/text_to_video" }

  it "POSTs to the correct endpoint" do
    params = {
      prompt: "Create a neon city tracking shot",
      duration_seconds: 8,
      aspect_ratio: "16:9",
      output_resolution: "1080p",
      reference_image_urls: ["https://file.runapi.ai/demo/scene.png"],
      audio_ids: ["audio-demo-123"],
      character_ids: ["character-demo-123"]
    }
    expect(http).to receive(:request).with(:post, endpoint, body: params)
      .and_return("id" => "task-local-123", "status" => "processing")

    result = resource.create(**params)

    expect(result.id).to eq("task-local-123")
    expect(result.status).to eq("processing")
  end

  it "GETs task status" do
    expect(http).to receive(:request).with(:get, "#{endpoint}/task-local-123")
      .and_return("id" => "task-local-123", "status" => "completed", "videos" => [{"url" => "https://tempfile.runapi.ai/gemini/output.mp4"}])

    result = resource.get("task-local-123")

    expect(result.status).to eq("completed")
    expect(result.videos.first.url).to eq("https://tempfile.runapi.ai/gemini/output.mp4")
  end

  it "raises ValidationError when required fields are missing" do
    expect { resource.create(prompt: "Create a neon city tracking shot") }
      .to raise_error(RunApi::Core::ValidationError, /duration_seconds is required/)
  end

  it "raises ValidationError for invalid reference quota" do
    expect {
      resource.create(
        prompt: "Create a neon city tracking shot",
        duration_seconds: 8,
        reference_image_urls: ["a", "b", "c"],
        character_ids: ["c1", "c2", "c3"],
        video_list: [{url: "https://cdn.runapi.ai/public/samples/source.mp4", start: 0, ends: 10}]
      )
    }.to raise_error(RunApi::Core::ValidationError, /reference units/)
  end
end
