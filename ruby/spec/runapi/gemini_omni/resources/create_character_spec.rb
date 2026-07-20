# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::GeminiOmni::Resources::CreateCharacter do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:resource) { described_class.new(http) }
  let(:endpoint) { "/api/v1/gemini_omni/create_character" }

  it "POSTs to the correct endpoint" do
    params = {
      descriptions: "A silver-haired cyberpunk guide",
      reference_image_url: "https://file.runapi.ai/demo/character.png",
      audio_ids: ["audio-runapi-123"],
      character_name: "Jenny"
    }
    expect(http).to receive(:request).with(:post, endpoint, body: params)
      .and_return("id" => "character-runapi-123", "character" => {"id" => "character-runapi-123", "name" => "Jenny", "images" => [{"url" => "https://file.runapi.ai/gemini/jenny.png"}]})

    result = resource.run(**params)

    expect(result.id).to eq("character-runapi-123")
    expect(result.character.name).to eq("Jenny")
    expect(result.character.images.first.url).to eq("https://file.runapi.ai/gemini/jenny.png")
  end

  it "raises ValidationError when required fields are missing" do
    expect { resource.run(reference_image_url: "https://file.runapi.ai/demo/character.png") }
      .to raise_error(RunApi::Core::ValidationError, /descriptions is required/)
  end

  it "raises ValidationError when reference image is missing" do
    expect {
      resource.run(
        descriptions: "A silver-haired cyberpunk guide"
      )
    }.to raise_error(RunApi::Core::ValidationError, /reference_image_url is required/)
  end
end
