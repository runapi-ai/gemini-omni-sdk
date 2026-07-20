# frozen_string_literal: true

Dir.chdir(__dir__) do

  Gem::Specification.new do |spec|
    spec.name = "runapi-gemini-omni"
    spec.version = "0.3.1"
    spec.metadata["runapi_slug"] = "gemini-omni"
    spec.authors = ["RunAPI"]
    spec.email = ["contact@runapi.ai"]

    spec.summary = "Gemini Omni Ruby SDK for RunAPI"
    spec.description = "The Gemini Omni Ruby SDK is the language-specific package for Gemini Omni on RunAPI. Use this package for voice resources, character resources, and multimodal video generation workflows when your application needs request bodies, task status lookup, and consistent RunAPI errors in Ruby."
    spec.homepage = "https://runapi.ai/models/gemini-omni"
    spec.license = "Apache-2.0"
    spec.required_ruby_version = ">= 3.1.0"
    spec.metadata["homepage_uri"] = "https://runapi.ai/models/gemini-omni"
    spec.metadata["documentation_uri"] = "https://github.com/runapi-ai/gemini-omni-sdk/blob/main/ruby/README.md"
    spec.metadata["source_code_uri"] = "https://github.com/runapi-ai/gemini-omni-sdk"
    spec.metadata["bug_tracker_uri"] = "https://github.com/runapi-ai/gemini-omni-sdk/issues"
    spec.metadata["changelog_uri"] = "https://github.com/runapi-ai/gemini-omni-sdk/blob/main/CHANGELOG.md"


    spec.files = Dir.glob("lib/**/*") + %w[LICENSE README.md]
    spec.extra_rdoc_files = ["README.md"]
        spec.require_paths = ["lib"]

    spec.add_dependency "runapi-core", "~> 0.2.14"
  end
end
