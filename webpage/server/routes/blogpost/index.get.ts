import type { BlogpostResponse } from "~~/shared/types/blogpost";

export default defineEventHandler(async () => {
  const data = await $fetch<BlogpostResponse>(
    "http://localhost:4000/api/v1/blogpost",
    {
      method: "GET",
    },
  );
  return data;
});
