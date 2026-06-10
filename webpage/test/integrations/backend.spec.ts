import { $fetch } from "ofetch";
import { describe, expect, inject, it } from "vitest";

const baseUrl = inject("integrationsBaseUrl");

describe("verify connection to backend", () => {
  it("verifies that the backend is running", async () => {
    const response = await $fetch(`${baseUrl}/api/v1/healthcheck`, {
      method: "GET",
    });

    expect(response).exist;
  });
});
