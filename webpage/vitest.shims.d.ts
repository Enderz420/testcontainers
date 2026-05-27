import "vitest";

declare module "vitest" {
  interface TestTags {
    tags: "user" | "blogpost" | "testcontainers";
  }
  export interface ProvidedContext {
    integrationsBaseUrl: string;
    e2eBaseUrl: string;
  }
}
