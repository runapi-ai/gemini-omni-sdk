package ai.runapi.geminiomni.resources;

import ai.runapi.core.ClientOptions;
import ai.runapi.core.RequestOptions;
import ai.runapi.core.http.HttpTransport;
import ai.runapi.geminiomni.types.CreateCharacterParams;
import ai.runapi.geminiomni.types.CreateCharacterResponse;

/** Create Character operations. */
public final class CreateCharacterResource extends GeminiomniResource {
  /** API endpoint path for create character operations. */
  public static final String ENDPOINT = "/api/v1/gemini_omni/create_character";

  /** Creates a resource bound to the supplied transport and client options. */
  public CreateCharacterResource(HttpTransport transport, ClientOptions options) {
    super(transport, options, ENDPOINT);
  }

  /** Runs create character and returns the response. */
  public CreateCharacterResponse run(CreateCharacterParams params) {
    return run(params, RequestOptions.none());
  }

  /** Runs create character with per-request options and returns the response. */
  public CreateCharacterResponse run(CreateCharacterParams params, RequestOptions options) {
    return runSync(params.action(), params.toMap(), options, CreateCharacterResponse.class);
  }
}
