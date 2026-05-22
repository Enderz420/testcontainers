import { debug } from "debug";
import {
  DockerComposeEnvironment,
  StartedDockerComposeEnvironment,
  Wait,
} from "testcontainers";
import { afterAll, beforeAll, describe, expect, it } from "vitest";
// import {  setup, } from '@nuxt/test-utils/e2e'

// await setup({
//    server
// })

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

      const backendContainer = environment.getContainer("backend-1");
      const stream = await backendContainer.logs();
      stream.on("data", (chunk) => process.stdout.write(`[backend] ${chunk}`));
      stream.on("err", (chunk) => process.stderr.write(`[backend] ${chunk}`));
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

    // TODO: Fix server routes in nuxt context
    // const input = await createBlogpost(body);
    const input = await $fetch<BlogpostResponse>(
      "http://localhost:4000/api/v1/blogpost",
      {
        method: "POST",
        body: body,
        headers: { "Content-Type": "application/json" },
      },
    );

    console.log(input);

    console.log("aswserting");

    expect(input.data.title).toBe(body.title);
    expect(input.data.content).toBe(body.content);
    expect(input.data.createdBy).toBe(body.created_by);
    console.log("test passed");
    await deleteBlogpost(input.data.id);
  });

  // it("gets all blogposts", async () => {
  //   const response = await getAllBlogposts();
  //   expect(response.metadata.Length).toBe(2);
  //   expect(response.data.length).toBe(2);
  // });
});
