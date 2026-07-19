# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::GeminiOmni::Client do
  let(:client) { described_class.new(api_key: "test-key") }

  it "exposes create_audio resource" do
    expect(client.create_audio).to be_a(RunApi::GeminiOmni::Resources::CreateAudio)
  end

  it "exposes create_character resource" do
    expect(client.create_character).to be_a(RunApi::GeminiOmni::Resources::CreateCharacter)
  end

  it "exposes text_to_video resource" do
    expect(client.text_to_video).to be_a(RunApi::GeminiOmni::Resources::TextToVideo)
  end
end
