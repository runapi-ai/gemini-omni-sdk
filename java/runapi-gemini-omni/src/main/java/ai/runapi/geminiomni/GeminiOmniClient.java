package ai.runapi.geminiomni;

import ai.runapi.core.BaseClient;
import ai.runapi.core.ClientOptions;
import ai.runapi.core.http.HttpTransport;
import java.net.URI;
import ai.runapi.geminiomni.resources.CreateAudioResource;
import ai.runapi.geminiomni.resources.CreateCharacterResource;
import ai.runapi.geminiomni.resources.TextToVideoResource;

/** GeminiOmni model-family Java SDK client. */
public final class GeminiOmniClient extends BaseClient {
  private final CreateAudioResource createAudio;
  private final CreateCharacterResource createCharacter;
  private final TextToVideoResource textToVideo;

  private GeminiOmniClient(ClientOptions options) {
    super(options);
    this.createAudio = new CreateAudioResource(transport(), options());
    this.createCharacter = new CreateCharacterResource(transport(), options());
    this.textToVideo = new TextToVideoResource(transport(), options());
  }

  /** Creates a new GeminiOmniClient builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Create Audio operations. */
  public CreateAudioResource createAudio() {
    return createAudio;
  }

  /** Create Character operations. */
  public CreateCharacterResource createCharacter() {
    return createCharacter;
  }

  /** Text To Video operations. */
  public TextToVideoResource textToVideo() {
    return textToVideo;
  }

  /** Builder for {@link GeminiOmniClient}. */
  public static final class Builder extends BaseClient.Builder<Builder> {
    private Builder() {}

    /** Sets the API key. If omitted, the SDK reads {@code RUNAPI_API_KEY}. */
    @Override
    public Builder apiKey(String value) {
      return super.apiKey(value);
    }

    /** Sets the RunAPI base URL. If omitted, the SDK reads {@code RUNAPI_BASE_URL}. */
    @Override
    public Builder baseUrl(String value) {
      return super.baseUrl(value);
    }

    /** Sets the RunAPI base URL from a URI. */
    @Override
    public Builder baseUrl(URI value) {
      return super.baseUrl(value);
    }

    /** Sets a custom HTTP transport. User-provided transports are not closed by SDK clients. */
    @Override
    public Builder transport(HttpTransport value) {
      return super.transport(value);
    }

    /** Builds an immutable GeminiOmniClient. */
    @Override
    public GeminiOmniClient build() {
      return new GeminiOmniClient(options.build());
    }
  }
}
