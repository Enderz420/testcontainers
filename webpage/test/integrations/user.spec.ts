import { $fetch, setup } from "@nuxt/test-utils/e2e";
import { debug } from "debug";
import { fileURLToPath } from "node:url";
import { describe, it } from "vitest";
import { PostUser } from "../../shared/types/user";

debug.enable("testcontainers*");
console.log("context:", process.env.NUXT_TEST_CONTEXT);

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
