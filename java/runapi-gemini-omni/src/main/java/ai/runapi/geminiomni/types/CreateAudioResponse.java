package ai.runapi.geminiomni.types;

import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.JsonNode;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.Map;

/** Response for create audio operations. */
public class CreateAudioResponse {
  @JsonProperty("id")
  private String id;

  @JsonProperty("error")
  private String error;

  @JsonProperty("audio")
  private Map<String, Object> audio;

  private final Map<String, JsonNode> extraFields = new LinkedHashMap<String, JsonNode>();

  /** Returns the response ID. */
  public String getId() {
    return id;
  }

  /** Returns the error message, if the request failed. */
  public String getError() {
    return error;
  }

  /** Returns the audio generation toggle, when present. */
  public Map<String, Object> getAudio() {
    return audio;
  }

  /** Returns unrecognized response fields preserved from the API response. */
  @JsonAnyGetter
  public Map<String, JsonNode> extraFields() {
    return Collections.unmodifiableMap(extraFields);
  }

  @JsonAnySetter
  void putExtraField(String name, JsonNode value) {
    extraFields.put(name, value);
  }
}
