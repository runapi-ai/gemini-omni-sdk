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
      audio_ids: ["audio-runapi-123"],
      character_ids: ["character-runapi-123"]
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

  it "POSTs Flash Preview with an explicit model and no duration" do
    params = {
      model: "gemini-omni-flash-preview",
      prompt: "A paper airplane flying through a sunlit studio",
      aspect_ratio: "9:16",
      output_resolution: "720p"
    }
    expect(http).to receive(:request).with(:post, endpoint, body: params)
      .and_return("id" => "task-flash-123", "status" => "processing")

    result = resource.create(**params)

    expect(result.id).to eq("task-flash-123")
  end

  it "POSTs Flash 1.1 with first and last frames at 360p" do
    params = {
      model: "gemini-omni-flash-1-1",
      prompt: "A paper airplane crosses from dawn into dusk",
      duration_seconds: 6,
      first_frame_image_url: "https://cdn.runapi.ai/public/samples/first-frame.jpg",
      last_frame_image_url: "https://cdn.runapi.ai/public/samples/last-frame.jpg",
      aspect_ratio: "16:9",
      output_resolution: "360p"
    }
    expect(http).to receive(:request).with(:post, endpoint, body: params)
      .and_return("id" => "task-flash-1-1-123", "status" => "processing")

    result = resource.create(**params)

    expect(result.id).to eq("task-flash-1-1-123")
  end
end
