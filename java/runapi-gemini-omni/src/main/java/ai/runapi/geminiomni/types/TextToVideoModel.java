package ai.runapi.geminiomni.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for text to video operations. */
public final class TextToVideoModel extends GeminiomniValue {
  /** gemini-omni-flash-1-1 model slug. */
  public static final TextToVideoModel GEMINI_OMNI_FLASH_1_1 =
      new TextToVideoModel(GeneratedModels.TEXT_TO_VIDEO_GEMINI_OMNI_FLASH_1_1);
  /** gemini-omni-flash-preview model slug. */
  public static final TextToVideoModel GEMINI_OMNI_FLASH_PREVIEW =
      new TextToVideoModel(GeneratedModels.TEXT_TO_VIDEO_GEMINI_OMNI_FLASH_PREVIEW);
  /** gemini-omni-text-to-video model slug. */
  public static final TextToVideoModel GEMINI_OMNI_TEXT_TO_VIDEO =
      new TextToVideoModel(GeneratedModels.TEXT_TO_VIDEO_GEMINI_OMNI_TEXT_TO_VIDEO);

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public TextToVideoModel(String value) {
    super(value);
  }
}
