package ai.runapi.geminiomni.types;

import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Parameters for text to video operations. */
public final class TextToVideoParams {
  private final String prompt;
  private final Integer durationSeconds;
  private final String callbackUrl;
  private final List<String> referenceImageUrls;
  private final List<String> audioIds;
  private final List<VideoClip> videoList;
  private final List<String> characterIds;
  private final String aspectRatio;
  private final String outputResolution;
  private final Integer seed;
  private final String model;

  private TextToVideoParams(Builder builder) {
    this.prompt = GeminiomniParamUtils.requireNonBlank(builder.prompt, "prompt");
    this.durationSeconds = java.util.Objects.requireNonNull(builder.durationSeconds, "durationSeconds");
    this.callbackUrl = builder.callbackUrl;
    this.referenceImageUrls = GeminiomniParamUtils.strings(builder.referenceImageUrls);
    this.audioIds = GeminiomniParamUtils.strings(builder.audioIds);
    this.videoList = GeminiomniParamUtils.list(builder.videoList, "videoList");
    this.characterIds = GeminiomniParamUtils.strings(builder.characterIds);
    this.aspectRatio = builder.aspectRatio;
    this.outputResolution = builder.outputResolution;
    this.seed = builder.seed;
    this.model = builder.model;
  }

  /** Creates a new TextToVideoParams builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the RunAPI action key for this request. */
  public String action() {
    return "gemini-omni/text-to-video";
  }

  /** Converts these parameters to the JSON request body shape. */
  public Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("prompt", GeminiomniParamUtils.wireValue(prompt));
    raw.put("duration_seconds", GeminiomniParamUtils.wireValue(durationSeconds));
    raw.put("callback_url", GeminiomniParamUtils.wireValue(callbackUrl));
    raw.put("reference_image_urls", GeminiomniParamUtils.wireValue(referenceImageUrls));
    raw.put("audio_ids", GeminiomniParamUtils.wireValue(audioIds));
    raw.put("video_list", videoListToMaps(videoList));
    raw.put("character_ids", GeminiomniParamUtils.wireValue(characterIds));
    raw.put("aspect_ratio", GeminiomniParamUtils.wireValue(aspectRatio));
    raw.put("output_resolution", GeminiomniParamUtils.wireValue(outputResolution));
    raw.put("seed", GeminiomniParamUtils.wireValue(seed));
    raw.put("model", GeminiomniParamUtils.wireValue(model));
    return GeminiomniParamUtils.compact(raw);
  }

  private static List<Map<String, Object>> videoListToMaps(List<VideoClip> values) {
    if (values == null) {
      return null;
    }
    List<Map<String, Object>> result = new ArrayList<Map<String, Object>>();
    for (VideoClip item : values) {
      result.add(item.toMap());
    }
    return java.util.Collections.unmodifiableList(result);
  }

  /** Builder for {@link TextToVideoParams}. */
  public static final class Builder {
    private String prompt;
    private Integer durationSeconds;
    private String callbackUrl;
    private List<String> referenceImageUrls;
    private List<String> audioIds;
    private List<VideoClip> videoList;
    private List<String> characterIds;
    private String aspectRatio;
    private String outputResolution;
    private Integer seed;
    private String model;

    private Builder() {}

    /** Sets the text prompt. */
    public Builder prompt(String value) {
      this.prompt = GeminiomniParamUtils.requireNonBlank(value, "prompt");
      return this;
    }

    /** Sets the duration in seconds. */
    public Builder durationSeconds(int value) {
      this.durationSeconds = value;
      return this;
    }

    /** Sets the webhook URL for task completion notifications. */
    public Builder callbackUrl(String value) {
      this.callbackUrl = GeminiomniParamUtils.requireNonBlank(value, "callbackUrl");
      return this;
    }

    /** Sets the reference image URLs. */
    public Builder referenceImageUrls(List<String> value) {
      this.referenceImageUrls = value;
      return this;
    }

    /** Sets the audio IDs. */
    public Builder audioIds(List<String> value) {
      this.audioIds = value;
      return this;
    }

    /** Sets the video list. */
    public Builder videoList(List<VideoClip> value) {
      this.videoList = value;
      return this;
    }

    /** Sets the character IDs. */
    public Builder characterIds(List<String> value) {
      this.characterIds = value;
      return this;
    }

    /** Sets the output aspect ratio. */
    public Builder aspectRatio(String value) {
      this.aspectRatio = GeminiomniParamUtils.requireNonBlank(value, "aspectRatio");
      return this;
    }

    /** Sets the output resolution. */
    public Builder outputResolution(String value) {
      this.outputResolution = GeminiomniParamUtils.requireNonBlank(value, "outputResolution");
      return this;
    }

    /** Sets the random seed. */
    public Builder seed(int value) {
      this.seed = value;
      return this;
    }

    /** Sets the model slug using a typed model value. */
    public Builder model(TextToVideoModel value) {
      this.model = java.util.Objects.requireNonNull(value, "model").value();
      return this;
    }

    /** Sets the model slug using a string value. */
    public Builder model(String value) {
      this.model = GeminiomniParamUtils.requireNonBlankTrim(value, "model");
      return this;
    }

    /** Builds immutable text to video parameters. */
    public TextToVideoParams build() {
      return new TextToVideoParams(this);
    }
  }
}
