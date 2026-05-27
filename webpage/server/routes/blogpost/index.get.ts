import type { BlogpostResponse } from "~~/shared/types/blogpost";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig(event);
  const data = await $fetch<BlogpostResponse>(
    `${config.public.url}/api/v1/blogpost`,
    {
      method: "GET",
    },
  );
  return data;
});
