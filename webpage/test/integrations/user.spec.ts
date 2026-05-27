import { $fetch, setup } from "@nuxt/test-utils/e2e";
import { fileURLToPath } from "node:url";
import { describe, expectTypeOf, it } from "vitest";
import { PostUser, User, UserResponse } from "../../shared/types/user";

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
    const input = await $fetch<UserResponse>("/user", {
      method: "POST",
      body: body,
    });

    console.log(input);

    expectTypeOf(input.data).toMatchObjectType<User>();
  });
});
