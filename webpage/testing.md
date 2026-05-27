blogpost.test.ts

```typescript
const databaseString: string = "";

let databaseClient: ConnectionPool;

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
  } catch (error) {
    throw error;
  }

  const body1: PostBlogpost = {
    Title: "Test",
    Content: "This is a test",
    CreatedBy: "TestUser",
  };

  const body2: PostBlogpost = {
    Title: "Test 2",
    Content: "This is another test",
    CreatedBy: "TestUser",
  };

  await Promise.all([
    $fetch(`${baseUrl}/api/v1/blogpost`, { method: "POST", body: body1 }),
    $fetch(`${baseUrl}/api/v1/blogpost`, { method: "POST", body: body2 }),
  ]);
}, 120_000);

afterAll(async () => {
  console.log("Shutting down");
  await environment?.down({ removeVolumes: true });
}, 120_000);
```

user.spec.ts

```typescript
const baseUrl = "http://localhost:4000";

debug.enable("testcontainers*");
console.log("context:", process.env.NUXT_TEST_CONTEXT);
let environment: StartedDockerComposeEnvironment;

beforeAll(async () => {
  console.log("Setting up");
  try {
    environment = await new DockerComposeEnvironment(
      "../",
      "./deployment/docker-compose.testing.yaml",
    )
      .withWaitStrategy("migrate", Wait.forOneShotStartup())
      .up();
  } catch (error) {
    throw error;
  }
}, 120_000);

afterAll(async () => {
  console.log("Shutting down");
  await environment?.down({ removeVolumes: true });
}, 120_000);
```
