import { getRouterParam, readBody } from "@nuxt/nitro-server/h3";
import { registerEndpoint } from "@nuxt/test-utils/runtime";
import { afterEach, describe, expect, it, test } from "vitest";
import { useBlogpost } from "../../app/composables/useBlogpost";
import { increment } from "../../app/composables/useIncrement";
import {
  Blogpost,
  BlogpostListResponse,
  BlogpostResponse,
  PostBlogpost,
} from "../../shared/types/blogpost";

const blogposts: Blogpost[] = [];

registerEndpoint("/blogpost", {
  method: "GET",
  handler: async (): Promise<BlogpostListResponse> => {
    return {
      results: blogposts,
      metadata: {
        last_seen: new Date().toISOString(),
        length: blogposts.length,
      },
    };
  },
});

registerEndpoint("/blogpost", {
  method: "POST",
  handler: async (event): Promise<BlogpostResponse> => {
    const body = await readBody<PostBlogpost>(event);
    const newPost: Blogpost = {
      id: crypto.randomUUID(),
      title: body.title,
      content: body.content,
      created_by: body.created_by,
      created_at: new Date(),
      updated_at: new Date(),
    };
    blogposts.push(newPost);
    return { results: newPost };
  },
});

registerEndpoint("/blogpost/:id", {
  method: "DELETE",
  handler: async (event): Promise<string> => {
    const id = getRouterParam(event, "id");
    const index = blogposts.findIndex((p) => p.id === id);
    if (index !== -1) blogposts.splice(index, 1);
    return "deleted";
  },
});

describe("increment", () => {
  test("it increments", () => {
    expect(increment(0, 10)).toBe(1);
  });
});

describe("test blogpost", async () => {
  afterEach(async () => {
    blogposts.length = 0;
  });
  const { getAllBlogposts, createBlogpost, deleteBlogpost } = useBlogpost();

  it("should create a blogpost", async () => {
    console.log("Starting test");
    const body: PostBlogpost = {
      title: "test",
      content: "this is a test",
      created_by: "testuser",
    };
    const input = await createBlogpost(body);

    console.log(input);

    console.log("aswserting");

    expect(input.results.title).toBe(body.title);
    expect(input.results.content).toBe(body.content);
    expect(input.results.created_by).toBe(body.created_by);
    console.log("test passed");
    console.log(input.results.id);
    // await deleteBlogpost(input.results.id);
  });

  it("gets all blogposts", async () => {
    const body1: PostBlogpost = {
      title: "Test 1",
      content: "This is a test",
      created_by: "TestUser",
    };

    const body2: PostBlogpost = {
      title: "Test 2",
      content: "This is another test",
      created_by: "TestUser2",
    };

    await Promise.all([createBlogpost(body1), createBlogpost(body2)]);

    const response = await getAllBlogposts();

    // const response = await $fetch<BlogpostListResponse>(
    //   "http://localhost:4000/api/v1/blogpost",
    //   {
    //     method: "GET",
    //   },
    // );

    console.log("Response: ", response);

    expect(response.metadata.length).toBe(2);
  });
});
