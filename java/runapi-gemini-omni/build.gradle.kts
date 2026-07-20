plugins {
  `java-library`
  `maven-publish`
}

extra["runapiSlug"] = "gemini-omni"

description = "RunAPI Gemini Omni Java SDK for Gemini Omni workflows."

java {
  withSourcesJar()
  withJavadocJar()
}

dependencies {
  api("ai.runapi:runapi-core:0.2.2")

  testImplementation(platform("org.junit:junit-bom:5.10.3"))
  testImplementation("org.junit.jupiter:junit-jupiter")
}

publishing {
  publications {
    create<MavenPublication>("mavenJava") {
      from(components["java"])
      artifactId = "runapi-gemini-omni"
      pom {
        name = "RunAPI Gemini Omni Java SDK"
        description = "RunAPI Gemini Omni Java SDK for Gemini Omni workflows."
        url = "https://runapi.ai/models/gemini-omni"
        licenses {
          license {
            name = "Apache License, Version 2.0"
            url = "https://www.apache.org/licenses/LICENSE-2.0"
          }
        }
        developers {
          developer {
            id = "runapi"
            name = "RunAPI"
            email = "contact@runapi.ai"
          }
        }
        scm {
          url = "https://github.com/runapi-ai/gemini-omni-sdk"
          connection = "scm:git:https://github.com/runapi-ai/gemini-omni-sdk.git"
          developerConnection = "scm:git:ssh://git@github.com/runapi-ai/gemini-omni-sdk.git"
        }
      }
    }
  }
}
