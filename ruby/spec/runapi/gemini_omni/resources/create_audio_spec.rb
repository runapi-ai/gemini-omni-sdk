# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::GeminiOmni::Resources::CreateAudio do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:resource) { described_class.new(http) }
  let(:endpoint) { "/api/v1/gemini_omni/create_audio" }

  it "POSTs to the correct endpoint" do
    params = {
      audio_id: "achernar",
      name: "Acher Narrator",
      voice_description: "A calm, clear voice",
      example_dialogue: "Hello, I am achernar"
    }
    expect(http).to receive(:request).with(:post, endpoint, body: params)
      .and_return("id" => "audio-runapi-123", "audio" => {"id" => "audio-runapi-123", "name" => "Acher Narrator"})

    result = resource.run(**params)

    expect(result.id).to eq("audio-runapi-123")
    expect(result.audio.name).to eq("Acher Narrator")
  end

  it "raises ValidationError when required fields are missing" do
    expect { resource.run(name: "Acher Narrator") }
      .to raise_error(RunApi::Core::ValidationError, /audio_id is required/)
  end

  it "raises ValidationError for documented max lengths" do
    expect {
      resource.run(audio_id: "achernar", name: "Acher Narrator", example_dialogue: "x" * 121)
    }.to raise_error(RunApi::Core::ValidationError, /example_dialogue must be at most 120 characters/)
  end
end
