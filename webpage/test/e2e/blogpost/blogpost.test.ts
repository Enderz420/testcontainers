import { mockNuxtImport, registerEndpoint } from "@nuxt/test-utils/runtime";
import { debug } from "debug";
import { readBody } from "h3";
import { describe, expect, inject, it } from "vitest";
import { ref } from "vue";
import { useBlogpost } from "../../../app/composables/useBlogpost";
import { PostBlogpost } from "../../../shared/types/blogpost";

const baseUrl = inject("e2eBaseUrl");

mockNuxtImport("useFetch", () => {
  return async (url: string, options?: any) => {
    const data = ref(null);
    const error = ref(null);
    const status = ref("idle");
    try {
      data.value = await $fetch(url, options);
      status.value = "success";
    } catch (e: any) {
      error.value = e;
      status.value = "error";
    }
    return { data, error, status };
  };
});

registerEndpoint("/blogpost", {
  method: "GET",
  handler: () => $fetch(`${baseUrl}/api/v1/blogpost`),
});

registerEndpoint("/blogpost", {
  method: "POST",
  handler: async (event) => {
    const body = await readBody(event);
    return $fetch(`${baseUrl}/api/v1/blogpost`, {
      method: "POST",
      body,
    });
  },
});

describe(
  "test blogpost",
  { tags: ["blogpost", "testcontainers"] },
  async () => {
    debug.enable("testcontainers*");

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

      expect(input.data.title).toBe(body.title);
      expect(input.data.content).toBe(body.content);
      expect(input.data.created_by).toBe(body.created_by);
      console.log("test passed");
      console.log(input.data.id);
      await $fetch(`${baseUrl}/api/v1/blogpost/${input.data.id}`, {
        method: "DELETE",
      });
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

      await Promise.all([
        $fetch(`${baseUrl}/api/v1/blogpost`, { method: "POST", body: body1 }),
        $fetch(`${baseUrl}/api/v1/blogpost`, { method: "POST", body: body2 }),
      ]);

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
  },
);
