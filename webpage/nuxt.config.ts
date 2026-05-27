// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },
  modules: ["@nuxt/ui"],
  runtimeConfig: {
    public: {
      url: "http://localhost:4000", // Default URL. Would typically be set by either an env file or the test suite
    },
  },
  hooks: {
    "pages:extend": (pages) => {
      const filteredPages = pages
        .filter((page) => {
          if (page.file === undefined) return true;
          if (page.file.endsWith(".ts")) return false;

          return true;
        })
        .filter((page) => {
          if (page.file === undefined) return true;
          if (page.file.includes("/components/")) return false;

          return true;
        });
      pages.length = 0;
      pages.push(...filteredPages);
    },
  },

  typescript: {
    typeCheck: true,
    tsConfig: {
      compilerOptions: {
        types: ["vitest/globals"],
      },
      include: ["/vitest.shims.d.ts"],
    },
  },
});
