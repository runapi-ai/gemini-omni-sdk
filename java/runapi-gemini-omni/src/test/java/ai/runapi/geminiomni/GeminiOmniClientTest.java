package ai.runapi.geminiomni;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;

import ai.runapi.core.RequestOptions;
import ai.runapi.core.errors.ValidationException;
import ai.runapi.core.http.HttpRequest;
import ai.runapi.core.http.HttpResponse;
import ai.runapi.core.http.HttpTransport;
import ai.runapi.core.http.JsonRequestBody;
import ai.runapi.core.json.Json;
import ai.runapi.geminiomni.types.CompletedTextToVideoResponse;
import ai.runapi.geminiomni.types.TextToVideoResponse;
import ai.runapi.geminiomni.types.CompletedTextToVideoResponse;
import ai.runapi.geminiomni.types.CreateAudioModel;
import ai.runapi.geminiomni.types.CreateAudioParams;
import ai.runapi.geminiomni.types.CreateAudioResponse;
import ai.runapi.geminiomni.types.CreateCharacterModel;
import ai.runapi.geminiomni.types.CreateCharacterParams;
import ai.runapi.geminiomni.types.CreateCharacterResponse;
import ai.runapi.geminiomni.types.TextToVideoModel;
import ai.runapi.geminiomni.types.TextToVideoParams;
import ai.runapi.geminiomni.types.TextToVideoResponse;
import ai.runapi.geminiomni.types.VideoClip;
import com.fasterxml.jackson.databind.JsonNode;
import java.io.ByteArrayOutputStream;
import java.time.Duration;
import java.util.Collections;
import org.junit.jupiter.api.Test;

class GeminiOmniClientTest {
  @Test
  void builderCreatesClientAndUniversalResources() {
    GeminiOmniClient client = GeminiOmniClient.builder().apiKey("sk-test").build();

    assertNotNull(client.textToVideo());
    assertNotNull(client.files());
    assertNotNull(client.account());
  }

  @Test
  void openValueClassesSerializeAsScalarStrings() throws Exception {
    String json = Json.mapper().writeValueAsString(new TextToVideoModel("gemini-omni-text-to-video"));

    assertEquals("\"gemini-omni-text-to-video\"", json);
    assertEquals(new TextToVideoModel("gemini-omni-text-to-video"), Json.mapper().readValue(json, TextToVideoModel.class));
  }

  @Test
  void createSendsExpectedRequestShape() throws Exception {
    CapturingTransport transport = new CapturingTransport("{\"id\":\"task_123\",\"status\":\"processing\"}");
    GeminiOmniClient client = GeminiOmniClient.builder().apiKey("sk-test").transport(transport).build();

    client.textToVideo().create(
        TextToVideoParams.builder()
            .prompt("A small red cube on a plain white table, studio product photo")
            .durationSeconds(4)
            .aspectRatio("16:9")
            .build()
    );

    assertEquals("POST", transport.request.getMethod().name());
    assertEquals("/api/v1/gemini_omni/text_to_video", transport.request.getPath());
    JsonNode body = bodyJson(transport.request);
    assertNotNull(body);
  }

  @Test
  void flashPreviewSendsModelWithoutDuration() throws Exception {
    CapturingTransport transport = new CapturingTransport("{\"id\":\"task_flash\",\"status\":\"processing\"}");
    GeminiOmniClient client = GeminiOmniClient.builder().apiKey("sk-test").transport(transport).build();

    client.textToVideo().create(
        TextToVideoParams.builder()
            .model(TextToVideoModel.GEMINI_OMNI_FLASH_PREVIEW)
            .prompt("A paper airplane flying through a sunlit studio")
            .aspectRatio("9:16")
            .outputResolution("720p")
            .build()
    );

    JsonNode body = bodyJson(transport.request);
    assertEquals("gemini-omni-flash-preview", body.get("model").asText());
    assertEquals(false, body.has("duration_seconds"));
  }

  @Test
  void getDecodesTaskResponseAndExtraFields() {
    CapturingTransport transport = new CapturingTransport("{\"id\":\"task_456\",\"status\":\"completed\",\"videos\":[{\"url\":\"https://file.runapi.ai/generated\"}],\"custom\":\"kept\"}");
    GeminiOmniClient client = GeminiOmniClient.builder().apiKey("sk-test").transport(transport).build();

    TextToVideoResponse response = client.textToVideo().get("task_456");

    assertEquals("GET", transport.request.getMethod().name());
    assertEquals("/api/v1/gemini_omni/text_to_video/task_456", transport.request.getPath());
    assertEquals("completed", response.getStatus().value());
    assertNotNull(response.getVideos());
    assertEquals("kept", response.extraFields().get("custom").asText());
  }

  @Test
  void runPollsUntilCompletedAndKeepsExtraFields() {
    SequenceTransport transport = new SequenceTransport(
        "{\"id\":\"task_789\",\"status\":\"processing\"}",
        "{\"id\":\"task_789\",\"status\":\"completed\",\"videos\":[{\"url\":\"https://file.runapi.ai/generated\"}],\"custom\":\"kept\"}");
    GeminiOmniClient client = GeminiOmniClient.builder().apiKey("sk-test").transport(transport).build();

    CompletedTextToVideoResponse response = client.textToVideo().run(
        TextToVideoParams.builder()
            .prompt("A small red cube on a plain white table, studio product photo")
            .durationSeconds(4)
            .aspectRatio("16:9")
            .build(),
        RequestOptions.builder().pollingInterval(Duration.ofMillis(1)).pollingMaxWait(Duration.ofSeconds(1)).build());

    assertEquals("completed", response.getStatus().value());
    assertNotNull(response.getVideos());
    assertEquals("kept", response.extraFields().get("custom").asText());
    assertEquals(2, transport.calls);
  }

  @Test
  void runRejectsCompletedResponseMissingResultField() {
    SequenceTransport transport = new SequenceTransport(
        "{\"id\":\"task_missing\",\"status\":\"processing\"}",
        "{\"id\":\"task_missing\",\"status\":\"completed\"}");
    GeminiOmniClient client = GeminiOmniClient.builder().apiKey("sk-test").transport(transport).build();

    assertThrows(
        ValidationException.class,
        () -> client.textToVideo().run(
                TextToVideoParams.builder()
                    .prompt("A small red cube on a plain white table, studio product photo")
                    .durationSeconds(4)
                    .aspectRatio("16:9")
                    .build(),
            RequestOptions.builder().pollingInterval(Duration.ofMillis(1)).pollingMaxWait(Duration.ofSeconds(1)).build()));
  }

    @Test
    void coversCreateaudioResourceMethods() {
      CapturingTransport transport = new CapturingTransport("{\"id\":\"sync_create_audio\",\"audio\":{\"value\":\"sample\"}}");
      GeminiOmniClient client = GeminiOmniClient.builder().apiKey("sk-test").transport(transport).build();

      CreateAudioResponse response = client.createAudio().run(
              CreateAudioParams.builder()
                  .audioId("sample")
                  .name("sample")
                  .build()
      );
      assertNotNull(response);

      CapturingTransport transportWithOptions = new CapturingTransport("{\"id\":\"sync_create_audio_options\",\"audio\":{\"value\":\"sample\"}}");
      GeminiOmniClient clientWithOptions = GeminiOmniClient.builder().apiKey("sk-test").transport(transportWithOptions).build();
      assertNotNull(clientWithOptions.createAudio().run(
              CreateAudioParams.builder()
                  .audioId("sample")
                  .name("sample")
                  .build(),
          RequestOptions.none()));
    }

    @Test
    void coversCreatecharacterResourceMethods() {
      CapturingTransport transport = new CapturingTransport("{\"id\":\"sync_create_character\",\"character\":{\"value\":\"sample\"}}");
      GeminiOmniClient client = GeminiOmniClient.builder().apiKey("sk-test").transport(transport).build();

      CreateCharacterResponse response = client.createCharacter().run(
              CreateCharacterParams.builder()
                  .descriptions("sample")
                  .referenceImageUrl("https://cdn.runapi.ai/public/samples/image.jpg")
                  .build()
      );
      assertNotNull(response);

      CapturingTransport transportWithOptions = new CapturingTransport("{\"id\":\"sync_create_character_options\",\"character\":{\"value\":\"sample\"}}");
      GeminiOmniClient clientWithOptions = GeminiOmniClient.builder().apiKey("sk-test").transport(transportWithOptions).build();
      assertNotNull(clientWithOptions.createCharacter().run(
              CreateCharacterParams.builder()
                  .descriptions("sample")
                  .referenceImageUrl("https://cdn.runapi.ai/public/samples/image.jpg")
                  .build(),
          RequestOptions.none()));
    }

    @Test
    void coversTexttovideoResourceMethods() {
      CapturingTransport createTransport = new CapturingTransport("{\"id\":\"task_text_to_video\",\"status\":\"processing\"}");
      GeminiOmniClient createClient = GeminiOmniClient.builder().apiKey("sk-test").transport(createTransport).build();
      assertNotNull(createClient.textToVideo().create(
              TextToVideoParams.builder()
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .durationSeconds(4)
                  .aspectRatio("16:9")
                  .build()
      ));

      CapturingTransport createWithOptionsTransport = new CapturingTransport("{\"id\":\"task_text_to_video_options\",\"status\":\"processing\"}");
      GeminiOmniClient createWithOptionsClient = GeminiOmniClient.builder().apiKey("sk-test").transport(createWithOptionsTransport).build();
      assertNotNull(createWithOptionsClient.textToVideo().create(
              TextToVideoParams.builder()
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .durationSeconds(4)
                  .aspectRatio("16:9")
                  .build(),
          RequestOptions.none()));

      CapturingTransport getTransport = new CapturingTransport("{\"id\":\"task_text_to_video\",\"status\":\"completed\",\"videos\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GeminiOmniClient getClient = GeminiOmniClient.builder().apiKey("sk-test").transport(getTransport).build();
      assertNotNull(getClient.textToVideo().get("task_text_to_video"));

      CapturingTransport getWithOptionsTransport = new CapturingTransport("{\"id\":\"task_text_to_video_options\",\"status\":\"completed\",\"videos\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GeminiOmniClient getWithOptionsClient = GeminiOmniClient.builder().apiKey("sk-test").transport(getWithOptionsTransport).build();
      assertNotNull(getWithOptionsClient.textToVideo().get("task_text_to_video_options", RequestOptions.none()));

      SequenceTransport runTransport = new SequenceTransport(
          "{\"id\":\"task_text_to_video_run\",\"status\":\"processing\"}",
          "{\"id\":\"task_text_to_video_run\",\"status\":\"completed\",\"videos\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GeminiOmniClient runClient = GeminiOmniClient.builder().apiKey("sk-test").transport(runTransport).build();
      CompletedTextToVideoResponse runResponse = runClient.textToVideo().run(
              TextToVideoParams.builder()
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .durationSeconds(4)
                  .aspectRatio("16:9")
                  .build(),
          RequestOptions.builder().pollingInterval(Duration.ofMillis(1)).pollingMaxWait(Duration.ofSeconds(1)).build());
      assertNotNull(runResponse);

      SequenceTransport runWithOptionsTransport = new SequenceTransport(
          "{\"id\":\"task_text_to_video_run_options\",\"status\":\"processing\"}",
          "{\"id\":\"task_text_to_video_run_options\",\"status\":\"completed\",\"videos\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GeminiOmniClient runWithOptionsClient = GeminiOmniClient.builder().apiKey("sk-test").transport(runWithOptionsTransport).build();
      assertNotNull(runWithOptionsClient.textToVideo().run(
              TextToVideoParams.builder()
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .durationSeconds(4)
                  .aspectRatio("16:9")
                  .build(),
          RequestOptions.builder().pollingInterval(Duration.ofMillis(1)).pollingMaxWait(Duration.ofSeconds(1)).build()));
    }

  private static JsonNode bodyJson(HttpRequest request) throws Exception {
    JsonRequestBody body = (JsonRequestBody) request.getBody();
    ByteArrayOutputStream out = new ByteArrayOutputStream();
    body.writeTo(out);
    return Json.mapper().readTree(out.toByteArray());
  }

  private static final class CapturingTransport implements HttpTransport {
    private final String body;
    private HttpRequest request;

    private CapturingTransport(String body) {
      this.body = body;
    }

    public HttpResponse send(HttpRequest request) {
      this.request = request;
      return new HttpResponse(200, body, Collections.<String, java.util.List<String>>emptyMap());
    }

    public void close() {}
  }

  private static final class SequenceTransport implements HttpTransport {
    private final String[] responses;
    private int calls;

    private SequenceTransport(String... responses) {
      this.responses = responses;
    }

    public HttpResponse send(HttpRequest request) {
      String response = responses[Math.min(calls, responses.length - 1)];
      calls++;
      return new HttpResponse(200, response, Collections.<String, java.util.List<String>>emptyMap());
    }

    public void close() {}
  }
}
