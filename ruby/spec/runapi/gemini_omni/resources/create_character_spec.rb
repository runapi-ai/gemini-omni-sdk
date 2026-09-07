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
      body_reference_image_url: "https://file.runapi.ai/demo/character-body.png",
      audio_ids: ["audio-runapi-123"],
      character_name: "Jenny"
    }
    expect(http).to receive(:request) do |method, path, body:, options:|
      expect([method, path, body]).to eq([:post, endpoint, params])
      expect(options.headers.fetch("Idempotency-Key")).to match(/\A[0-9a-f-]{36}\z/)
    end.and_return("id" => "character-runapi-123", "character" => {"id" => "character-runapi-123", "name" => "Jenny", "images" => [{"url" => "https://file.runapi.ai/gemini/jenny.png"}, {"url" => "https://file.runapi.ai/gemini/jenny-body.png"}]}, "billing" => {"reservation" => {"amount_cents" => 10}, "settlement" => {"charged_amount_cents" => 10, "amount_micro_cents" => 10_000_000}, "refund" => nil})

    result = resource.run(**params)

    expect(result.id).to eq("character-runapi-123")
    expect(result.character.name).to eq("Jenny")
    expect(result.character.images.first.url).to eq("https://file.runapi.ai/gemini/jenny.png")
    expect(result.character.images.last.url).to eq("https://file.runapi.ai/gemini/jenny-body.png")
    expect(result.billing).to be_a(RunApi::Core::TaskBillingFacts)
    expect(result.billing.reservation.amount_cents).to eq(10)
  end

  it "follows an accepted task to its terminal character response" do
    client = RunApi::Core::HttpClient.new(
      RunApi::Core::ClientOptions.new(api_key: "test-key", base_url: "https://api.runapi.ai")
    )
    resource = described_class.new(client)
    location = "https://api.runapi.ai/api/v1/tasks/task-1/result"

    stub_request(:post, "https://api.runapi.ai#{endpoint}")
      .to_return(
        status: 202,
        body: '{"id":"task-1","status":"processing"}',
        headers: {"Content-Type" => "application/json", "Location" => location, "Retry-After" => "0"}
      )
    stub_request(:get, location)
      .to_return(
        status: 200,
        body: '{"id":"task-1","status":"completed","response":{"status":200,"content_type":"application/json","headers":{},"body":{"id":"character-1","character":{"id":"character-1","name":"Jenny","images":[]}}}}',
        headers: {"Content-Type" => "application/json"}
      )

    result = resource.run(
      descriptions: "A silver-haired guide",
      reference_image_url: "https://file.runapi.ai/demo/character.png"
    )

    expect(result).to be_a(RunApi::GeminiOmni::Types::CreateCharacterResponse)
    expect(result.id).to eq("character-1")
    expect(result.character.id).to eq("character-1")
    expect(a_request(:get, location)).to have_been_made.once
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
