package ai.runapi.geminiomni.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for create character operations. */
public final class CreateCharacterModel extends GeminiomniValue {
  /** gemini-omni-character model slug. */
  public static final CreateCharacterModel GEMINI_OMNI_CHARACTER = new CreateCharacterModel("gemini-omni-character");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public CreateCharacterModel(String value) {
    super(value);
  }
}
