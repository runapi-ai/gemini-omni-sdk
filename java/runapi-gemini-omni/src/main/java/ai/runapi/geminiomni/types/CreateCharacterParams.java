package ai.runapi.geminiomni.types;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Parameters for create character operations. */
public final class CreateCharacterParams {
  private final String descriptions;
  private final String referenceImageUrl;
  private final List<String> audioIds;
  private final String characterName;
  private final String model;

  private CreateCharacterParams(Builder builder) {
    this.descriptions = GeminiomniParamUtils.requireNonBlank(builder.descriptions, "descriptions");
    this.referenceImageUrl = GeminiomniParamUtils.requireNonBlank(builder.referenceImageUrl, "referenceImageUrl");
    this.audioIds = GeminiomniParamUtils.strings(builder.audioIds);
    this.characterName = builder.characterName;
    this.model = builder.model;
  }

  /** Creates a new CreateCharacterParams builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the RunAPI action key for this request. */
  public String action() {
    return "gemini-omni/create-character";
  }

  /** Converts these parameters to the JSON request body shape. */
  public Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("descriptions", GeminiomniParamUtils.wireValue(descriptions));
    raw.put("reference_image_url", GeminiomniParamUtils.wireValue(referenceImageUrl));
    raw.put("audio_ids", GeminiomniParamUtils.wireValue(audioIds));
    raw.put("character_name", GeminiomniParamUtils.wireValue(characterName));
    raw.put("model", GeminiomniParamUtils.wireValue(model));
    return GeminiomniParamUtils.compact(raw);
  }



  /** Builder for {@link CreateCharacterParams}. */
  public static final class Builder {
    private String descriptions;
    private String referenceImageUrl;
    private List<String> audioIds;
    private String characterName;
    private String model;

    private Builder() {}

    /** Sets the descriptions. */
    public Builder descriptions(String value) {
      this.descriptions = GeminiomniParamUtils.requireNonBlank(value, "descriptions");
      return this;
    }

    /** Sets the reference image URL. */
    public Builder referenceImageUrl(String value) {
      this.referenceImageUrl = GeminiomniParamUtils.requireNonBlank(value, "referenceImageUrl");
      return this;
    }

    /** Sets the audio IDs. */
    public Builder audioIds(List<String> value) {
      this.audioIds = value;
      return this;
    }

    /** Sets the character name. */
    public Builder characterName(String value) {
      this.characterName = GeminiomniParamUtils.requireNonBlank(value, "characterName");
      return this;
    }

    /** Sets the model slug using a typed model value. */
    public Builder model(CreateCharacterModel value) {
      this.model = java.util.Objects.requireNonNull(value, "model").value();
      return this;
    }

    /** Sets the model slug using a string value. */
    public Builder model(String value) {
      this.model = GeminiomniParamUtils.requireNonBlankTrim(value, "model");
      return this;
    }

    /** Builds immutable create character parameters. */
    public CreateCharacterParams build() {
      return new CreateCharacterParams(this);
    }
  }
}
