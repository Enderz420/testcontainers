import debug from "debug";
import {
  DockerComposeEnvironment,
  StartedDockerComposeEnvironment,
  Wait,
} from "testcontainers";
import { TestProject } from "vitest/node";

let environment: StartedDockerComposeEnvironment;

export async function setup(module: TestProject) {
  // debug.enable("testcontainers*");

  environment = await new DockerComposeEnvironment(
    "../",
    "./deployment/docker-compose.testing.yaml",
  )
    .withWaitStrategy("migrate-1", Wait.forOneShotStartup())
    .up();

  const port = environment.getContainer("backend-1").getMappedPort(4000);
  process.env.NUXT_PUBLIC_URL = `http://localhost:${port}`;
  module.provide("integrationsBaseUrl", `http://localhost:${port}`);
}

export async function teardown() {
  await environment?.down({ removeVolumes: true });
}
