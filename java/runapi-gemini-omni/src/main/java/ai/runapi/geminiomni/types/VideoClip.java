package ai.runapi.geminiomni.types;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Nested request item for typed parameter builders. */
public final class VideoClip {
  private final String url;
  private final Double start;
  private final Double ends;

  private VideoClip(Builder builder) {
    this.url = GeminiomniParamUtils.requireNonBlank(builder.url, "url");
    this.start = java.util.Objects.requireNonNull(builder.start, "start");
    this.ends = java.util.Objects.requireNonNull(builder.ends, "ends");
  }

  /** Creates a new VideoClip builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the media URL. */
  public String getUrl() {
    return url;
  }

  /** Returns the clip start time. */
  public Double getStart() {
    return start;
  }

  /** Returns the clip end time. */
  public Double getEnds() {
    return ends;
  }

  Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("url", GeminiomniParamUtils.wireValue(url));
    raw.put("start", GeminiomniParamUtils.wireValue(start));
    raw.put("ends", GeminiomniParamUtils.wireValue(ends));
    return GeminiomniParamUtils.compact(raw);
  }

  /** Builder for {@link VideoClip}. */
  public static final class Builder {
    private String url;
    private Double start;
    private Double ends;

    private Builder() {}

    /** Sets the media URL. */
    public Builder url(String value) {
      this.url = GeminiomniParamUtils.requireNonBlank(value, "url");
      return this;
    }

    /** Sets the clip start time. */
    public Builder start(double value) {
      this.start = value;
      return this;
    }

    /** Sets the clip end time. */
    public Builder ends(double value) {
      this.ends = value;
      return this;
    }

    /** Builds an immutable VideoClip. */
    public VideoClip build() {
      return new VideoClip(this);
    }
  }
}
