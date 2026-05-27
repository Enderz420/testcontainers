import { $fetch, setup } from "@nuxt/test-utils/e2e";
import { fileURLToPath } from "node:url";
import { describe, expectTypeOf, it } from "vitest";
import {
  Blogpost,
  BlogpostListResponse,
  BlogpostResponse,
  PostBlogpost,
} from "../../shared/types/blogpost";

await setup({
  rootDir: fileURLToPath(new URL("../../", import.meta.url)),
  server: true,
  build: true,
});

/**
 * Blogpost Integration tests
 * @module-tag blogpost
 */
describe("Blogpost Integration Test", async ({ afterEach }) => {
  let createdId: string[] = [];

  afterEach(async () => {
    if (createdId.length > 0) {
      await Promise.all(
        createdId.map(async (id) => {
          await $fetch(`/blogpost/${id}`, { method: "DELETE" });
        }),
      );
      createdId = [];
    }
  });

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

    createdId.push(input.results.id);

    expectTypeOf(input.results).toMatchObjectType<Blogpost>();

    expect(input.results.title).toBe(body.title);
    expect(input.results.content).toBe(body.content);
    expect(input.results.created_by).toBe(body.created_by);

    const inserted = await $fetch<BlogpostResponse>(
      `/blogpost/${input.results.id}`,
      {
        method: "GET",
      },
    );

    console.log(inserted);

    expectTypeOf(inserted.results).toMatchObjectType<Blogpost>();
    expect(inserted.results.id).toBe(input.results.id);
    expect(inserted.results.title).toBe(body.title);
    expect(inserted.results.content).toBe(body.content);
    expect(inserted.results.created_by).toBe(body.created_by);
    expect(inserted.results.created_at).toEqual(input.results.created_at);
    expect(inserted.results.updated_at).toEqual(input.results.updated_at);
  });

  /**
   * Creates two blogposts and verifies that they are created by the backend and
   * fetchable by the frontend.
   *
   */
  it("gets multiple blogposts", async ({ expect, annotate }) => {
    const body1: PostBlogpost = {
      title: "Test1",
      content: "This is a test",
      created_by: "GenericUser1",
    };
    const body2: PostBlogpost = {
      title: "Test2",
      content: "This is another test",
      created_by: "GenericUser2",
    };

    const [test1, test2] = await Promise.all([
      $fetch<BlogpostResponse>(`/blogpost`, { method: "POST", body: body1 }),
      $fetch<BlogpostResponse>(`/blogpost`, { method: "POST", body: body2 }),
    ]);

    expectTypeOf(test1).toMatchObjectType<BlogpostResponse>();
    expectTypeOf(test2).toMatchObjectType<BlogpostResponse>();
    expect(test1.results.title).toBe(body1.title);
    expect(test2.results.title).toBe(body2.title);

    createdId.push(test1.results.id, test2.results.id);

    const response = await $fetch<BlogpostListResponse>("/blogpost", {
      method: "GET",
    });
    console.log(response);

    expectTypeOf(response).toMatchObjectType<BlogpostListResponse>();

    expect(response.metadata.length).toBe(2);
    expect(response.results).toEqual(
      expect.arrayContaining([
        expect.objectContaining({
          title: "Test2",
          content: "This is another test",
          created_by: "GenericUser2",
        }),
      ]),
    );
  });
});
