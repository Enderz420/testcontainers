import { describe, expect, test } from "vitest";
import { increment } from "~/composables/useIncrement";

describe("increment", () => {
  test("it increments", () => {
    expect(increment(0, 10)).toBe(1);
  });
});
