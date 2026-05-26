import { $fetch, setup } from "@nuxt/test-utils/e2e";
import { debug } from "debug";
import { fileURLToPath } from "node:url";
import {
  DockerComposeEnvironment,
  StartedDockerComposeEnvironment,
  Wait,
} from "testcontainers";
import { afterAll, beforeAll, describe, it } from "vitest";
import { PostUser } from "../../shared/types/user";

// const baseUrl = "http://localhost:4000";

debug.enable("testcontainers*");
console.log("context:", process.env.NUXT_TEST_CONTEXT);
// let environment: StartedDockerComposeEnvironment;

// beforeAll(async () => {
//   console.log("Setting up");
//   try {
//     environment = await new DockerComposeEnvironment(
//       "../",
//       "./deployment/docker-compose.testing.yaml",
//     )
//       .withWaitStrategy("migrate", Wait.forOneShotStartup())
//       .up();
//   } catch (error) {
//     throw error;
//   }
// }, 120_000);

// afterAll(async () => {
//   console.log("Shutting down");
//   await environment?.down({ removeVolumes: true });
// }, 120_000);

await setup({
  rootDir: fileURLToPath(new URL("../../", import.meta.url)),
  server: true,
  build: true,
});

describe("user test", async () => {
  it("should create a user", async () => {
    const body: PostUser = {
      username: "tester",
      email: "test@test.com",
    };
    const input = await $fetch("/user", {
      method: "POST",
      body: body,
    });

    console.log(input);
  });
});
