import { $fetch, setup } from "@nuxt/test-utils/e2e";
import { fileURLToPath } from "node:url";
import { describe, expectTypeOf, it } from "vitest";
import {
  Blogpost,
  BlogpostResponse,
  PostBlogpost,
} from "../../shared/types/blogpost";

await setup({
  rootDir: fileURLToPath(new URL("../../", import.meta.url)),
  server: true,
  build: true,
});

describe("Blogpost Integration Test", { tags: "blogpost" }, async ({}) => {
  it("creates a blogpost", async ({ expect }) => {
    const body: PostBlogpost = {
      title: "Test",
      content: "This is a test",
      created_by: "GenericUser",
    };

    const input = await $fetch<BlogpostResponse>("/blogpost", {
      method: "POST",
      body: body,
    });

    expectTypeOf(input.data).toMatchObjectType<Blogpost>();

    expect(input.data.title).toBe(body.title);
    expect(input.data.content).toBe(body.content);
    expect(input.data.created_by).toBe(body.created_by);

    const inserted = await $fetch<BlogpostResponse>(
      `/blogpost/${input.data.id}`,
      {
        method: "GET",
      },
    );

    console.log(inserted);

    expectTypeOf(inserted.data).toMatchObjectType<Blogpost>();
    expect(inserted.data.id).toBe(input.data.id);
    expect(inserted.data.title).toBe(body.title);
    expect(inserted.data.content).toBe(body.content);
    expect(inserted.data.created_by).toBe(body.created_by);
    expect(inserted.data.created_at).toEqual(input.data.created_at);
    expect(inserted.data.updated_at).toEqual(input.data.updated_at);
  });
});
