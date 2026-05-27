import { defineVitestProject } from "@nuxt/test-utils/config";
import { defineConfig } from "vitest/config";

export default defineConfig({
  test: {
    // globalSetup: ["./test/setup/global.ts"],
    tags: [
      {
        name: "user",
        description: "Tests related to the user module.",
      },
      {
        name: "blogpost",
        description: "Tests related to the blogpost module.",
      },
      {
        name: "testcontainers",
        description:
          "Tests that use testcontainers to verify integrations with the backend service",
      },
    ],
    projects: [
      {
        test: {
          name: "unit",
          include: ["test/unit/*.{test,spec}.ts"],
          environment: "node",
        },
      },

      await defineVitestProject({
        test: {
          name: "nuxt",
          include: ["test/nuxt/*.{test,spec}.ts"],
          environment: "nuxt",
        },
      }),
      await defineVitestProject({
        test: {
          name: "e2e",
          include: ["test/e2e/**/*.{test,spec}.ts"],
          environment: "nuxt",
          globalSetup: ["./test/setup/e2e.global.ts"],
          hookTimeout: 120_000,
          testTimeout: 15_000,
        },
      }),
      {
        test: {
          name: "integrations",
          include: ["test/integrations/**/*.{test,spec}.ts"],
          environment: "node",
          globalSetup: ["./test/setup/integrations.global.ts"],
          hookTimeout: 120_000,
          testTimeout: 15_000,
        },
      },
    ],
  },
});
