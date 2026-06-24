package ai.runapi.geminiomni.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for create audio operations. */
public final class CreateAudioModel extends GeminiomniValue {
  /** gemini-omni-audio model slug. */
  public static final CreateAudioModel GEMINI_OMNI_AUDIO = new CreateAudioModel("gemini-omni-audio");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public CreateAudioModel(String value) {
    super(value);
  }
}
