import { mountSuspended } from "@nuxt/test-utils/runtime";
import { describe, expect, it } from "vitest";
import Navbar from "../../app/components/Navbar.vue";

describe("component", async () => {
  it("mounts component and checks p tag for test", async () => {
    const component = await mountSuspended(Navbar);
    expect(component.find("p").text()).toBe("This is text");
  });
  it("mounts component and checks a tag for test", async () => {
    const component = await mountSuspended(Navbar);
    expect(component.find("a").exists()).toBe(true);
  });
});
