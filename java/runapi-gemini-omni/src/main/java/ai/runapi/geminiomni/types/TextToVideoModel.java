package ai.runapi.geminiomni.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for text to video operations. */
public final class TextToVideoModel extends GeminiomniValue {
  /** gemini-omni-text-to-video model slug. */
  public static final TextToVideoModel GEMINI_OMNI_TEXT_TO_VIDEO = new TextToVideoModel("gemini-omni-text-to-video");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public TextToVideoModel(String value) {
    super(value);
  }
}
