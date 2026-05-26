import {
  buildFixture,
  createTestContext,
  exposeContextToEnv,
  loadFixture,
  startServer,
  stopServer,
} from "@nuxt/test-utils/e2e";
import { fileURLToPath } from "node:url";
import {
  DockerComposeEnvironment,
  StartedDockerComposeEnvironment,
  Wait,
} from "testcontainers";

let environment: StartedDockerComposeEnvironment;

export async function setup() {
  environment = await new DockerComposeEnvironment(
    "../",
    "./deployment/docker-compose.testing.yaml",
  )
    .withWaitStrategy("migrate", Wait.forOneShotStartup())
    .up();

  createTestContext({
    rootDir: fileURLToPath(new URL("../../webpage", import.meta.url)),
    server: true,
    build: true,
    runner: "vitest",
  });

  await loadFixture();
  await buildFixture();
  await startServer();
  exposeContextToEnv(); // shares server URL via process.env
}

export async function teardown() {
  await stopServer();
  await environment?.down({ removeVolumes: true });
}
