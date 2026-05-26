import {
  buildFixture,
  createTestContext,
  exposeContextToEnv,
  loadFixture,
  startServer,
  stopServer,
} from "@nuxt/test-utils/e2e";
import debug from "debug";
import { fileURLToPath } from "node:url";
import {
  DockerComposeEnvironment,
  StartedDockerComposeEnvironment,
  Wait,
} from "testcontainers";

let environment: StartedDockerComposeEnvironment;

export async function setup() {
  debug.enable("testcontainers*");

  environment = await new DockerComposeEnvironment(
    "../",
    "./deployment/docker-compose.testing.yaml",
  )
    .withWaitStrategy("migrate", Wait.forOneShotStartup())
    .up();

  // createTestContext({
  //   rootDir: fileURLToPath(new URL("../../", import.meta.url)),
  //   server: true,
  //   build: true,
  //   runner: "vitest",
  // });

  // await loadFixture();
  // await buildFixture();
  // await startServer();
  // exposeContextToEnv();
}

export async function teardown() {
  await environment?.down({ removeVolumes: true });
}
