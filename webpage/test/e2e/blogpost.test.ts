import { mockNuxtImport, registerEndpoint } from "@nuxt/test-utils/runtime";
import { debug } from "debug";
import { readBody } from "h3";
import {
  DockerComposeEnvironment,
  StartedDockerComposeEnvironment,
  Wait,
} from "testcontainers";
import { afterAll, beforeAll, describe, expect, it } from "vitest";

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
  handler: () => $fetch("http://localhost:4000/api/v1/blogpost"),
});

registerEndpoint("/blogpost", {
  method: "POST",
  handler: async (event) => {
    const body = await readBody(event);
    return $fetch("http://localhost:4000/api/v1/blogpost", {
      method: "POST",
      body,
    });
  },
});

describe("test blogpost", async () => {
  //   const databaseString: string = "";

  //   let databaseClient: ConnectionPool;

  const baseUrl = "http://localhost:4000";

  let environment: StartedDockerComposeEnvironment;

  debug.enable("testcontainers*");

  const { getAllBlogposts, createBlogpost, deleteBlogpost } = useBlogpost();

  beforeAll(async () => {
    console.log("Setting up");
    try {
      environment = await new DockerComposeEnvironment(
        "../",
        "./deployment/docker-compose.testing.yaml",
      )
        .withWaitStrategy("migrate", Wait.forOneShotStartup())
        .up();

      // registerEndpoint("/blogpost", async (event) => {
      //   const data = await readBody(event);
      //   console.log("request ", data);

      //   const response = await $fetch<BlogpostResponse>(
      //     "http://localhost:4000/api/v1/blogpost",
      //     {
      //       method: "POST",
      //       body: data,
      //     },
      //   );
      //   console.log("returning");
      //   return response;
      // });
    } catch (error) {
      throw error;
    }

    // const body1: PostBlogpost = {
    //   Title: "Test",
    //   Content: "This is a test",
    //   CreatedBy: "TestUser",
    // };

    // const body2: PostBlogpost = {
    //   Title: "Test 2",
    //   Content: "This is another test",
    //   CreatedBy: "TestUser",
    // };

    // await Promise.all([
    //   $fetch(`${baseUrl}/api/v1/blogpost`, { method: "POST", body: body1 }),
    //   $fetch(`${baseUrl}/api/v1/blogpost`, { method: "POST", body: body2 }),
    // ]);
  }, 120_000);

  afterAll(async () => {
    console.log("Shutting down");
    await environment?.down({ removeVolumes: true });
  }, 120_000);

  it("should create a blogpost", async () => {
    console.log("Starting test");
    const body: PostBlogpost = {
      title: "test",
      content: "this is a test",
      created_by: "testuser",
    };
    const input = await createBlogpost(body);

    // const input = await $fetch<BlogpostResponse>(
    //   "http://localhost:4000/api/v1/blogpost",
    //   {
    //     method: "POST",
    //     body: body,
    //     headers: { "Content-Type": "application/json" },
    //   },
    // );

    console.log(input);

    console.log("aswserting");

    expect(input.data.title).toBe(body.title);
    expect(input.data.content).toBe(body.content);
    expect(input.data.created_by).toBe(body.created_by);
    console.log("test passed");
    console.log(input.data.id);
    await $fetch(`http://localhost:4000/api/v1/blogpost/${input.data.id}`, {
      method: "DELETE",
    });
  });

  // TODO: Fix server routes in nuxt context
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

    // const response = await getAllBlogposts();

    const response = await $fetch<BlogpostListResponse>(
      "http://localhost:4000/api/v1/blogpost",
      {
        method: "GET",
      },
    );

    console.log("Response: ", response);

    expect(response.metadata.length).toBe(2);
  });
});
