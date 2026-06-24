package ai.runapi.geminiomni.resources;

import ai.runapi.core.ClientOptions;
import ai.runapi.core.RequestOptions;
import ai.runapi.core.http.HttpTransport;
import ai.runapi.geminiomni.types.CreateAudioParams;
import ai.runapi.geminiomni.types.CreateAudioResponse;

/** Create Audio operations. */
public final class CreateAudioResource extends GeminiomniResource {
  /** API endpoint path for create audio operations. */
  public static final String ENDPOINT = "/api/v1/gemini_omni/create_audio";

  /** Creates a resource bound to the supplied transport and client options. */
  public CreateAudioResource(HttpTransport transport, ClientOptions options) {
    super(transport, options, ENDPOINT);
  }

  /** Runs create audio and returns the response. */
  public CreateAudioResponse run(CreateAudioParams params) {
    return run(params, RequestOptions.none());
  }

  /** Runs create audio with per-request options and returns the response. */
  public CreateAudioResponse run(CreateAudioParams params, RequestOptions options) {
    return runSync(params.action(), params.toMap(), options, CreateAudioResponse.class);
  }
}
