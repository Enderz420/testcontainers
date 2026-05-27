import { defineVitestProject } from "@nuxt/test-utils/config";
import { defineConfig } from "vitest/config";

export default defineConfig({
  test: {
    globalSetup: ["./test/setup/global.ts"],
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
        },
      }),
      {
        test: {
          name: "integrations",
          include: ["test/integrations/**/*.{test,spec}.ts"],
          environment: "node",
          // globalSetup: ["./test/setup/global.ts"],
        },
      },
    ],
  },
});
