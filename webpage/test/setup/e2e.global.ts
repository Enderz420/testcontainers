import debug from "debug";
import {
  DockerComposeEnvironment,
  StartedDockerComposeEnvironment,
  Wait,
} from "testcontainers";
import { TestProject } from "vitest/node";

let environment: StartedDockerComposeEnvironment;

export async function setup(module: TestProject) {
  debug.enable("testcontainers*");

  environment = await new DockerComposeEnvironment(
    "../",
    "./deployment/docker-compose.testing.yaml",
  )
    .withWaitStrategy("migrate", Wait.forOneShotStartup())
    .up();

  const port = environment.getContainer("backend").getMappedPort(4000);
  module.provide("integrationsBaseUrl", `http://localhost:${port}`);
}

export async function teardown() {
  environment?.down({ removeVolumes: true });
}
