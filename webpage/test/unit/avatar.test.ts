import { render, screen } from "@testing-library/vue";
import { describe, expect, it } from "vitest";
import avatar from "../../app/components/avatar.vue";

describe("Test avatar", () => {
  it("Renders and sets props for the avatar", () => {
    render(avatar, {
      props: { name: "Test" },
    });
    expect(screen.getByText("Test")).toBeDefined();
  });
});
