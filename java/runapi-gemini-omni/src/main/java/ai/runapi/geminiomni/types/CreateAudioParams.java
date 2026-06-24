package ai.runapi.geminiomni.types;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Parameters for create audio operations. */
public final class CreateAudioParams {
  private final String audioId;
  private final String name;
  private final String voiceDescription;
  private final String exampleDialogue;
  private final String model;

  private CreateAudioParams(Builder builder) {
    this.audioId = GeminiomniParamUtils.requireNonBlank(builder.audioId, "audioId");
    this.name = GeminiomniParamUtils.requireNonBlank(builder.name, "name");
    this.voiceDescription = builder.voiceDescription;
    this.exampleDialogue = builder.exampleDialogue;
    this.model = builder.model;
  }

  /** Creates a new CreateAudioParams builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the RunAPI action key for this request. */
  public String action() {
    return "gemini-omni/create-audio";
  }

  /** Converts these parameters to the JSON request body shape. */
  public Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("audio_id", GeminiomniParamUtils.wireValue(audioId));
    raw.put("name", GeminiomniParamUtils.wireValue(name));
    raw.put("voice_description", GeminiomniParamUtils.wireValue(voiceDescription));
    raw.put("example_dialogue", GeminiomniParamUtils.wireValue(exampleDialogue));
    raw.put("model", GeminiomniParamUtils.wireValue(model));
    return GeminiomniParamUtils.compact(raw);
  }



  /** Builder for {@link CreateAudioParams}. */
  public static final class Builder {
    private String audioId;
    private String name;
    private String voiceDescription;
    private String exampleDialogue;
    private String model;

    private Builder() {}

    /** Sets the audio ID. */
    public Builder audioId(String value) {
      this.audioId = GeminiomniParamUtils.requireNonBlank(value, "audioId");
      return this;
    }

    /** Sets the item name. */
    public Builder name(String value) {
      this.name = GeminiomniParamUtils.requireNonBlank(value, "name");
      return this;
    }

    /** Sets the voice description. */
    public Builder voiceDescription(String value) {
      this.voiceDescription = GeminiomniParamUtils.requireNonBlank(value, "voiceDescription");
      return this;
    }

    /** Sets the example dialogue. */
    public Builder exampleDialogue(String value) {
      this.exampleDialogue = GeminiomniParamUtils.requireNonBlank(value, "exampleDialogue");
      return this;
    }

    /** Sets the model slug using a typed model value. */
    public Builder model(CreateAudioModel value) {
      this.model = java.util.Objects.requireNonNull(value, "model").value();
      return this;
    }

    /** Sets the model slug using a string value. */
    public Builder model(String value) {
      this.model = GeminiomniParamUtils.requireNonBlankTrim(value, "model");
      return this;
    }

    /** Builds immutable create audio parameters. */
    public CreateAudioParams build() {
      return new CreateAudioParams(this);
    }
  }
}
