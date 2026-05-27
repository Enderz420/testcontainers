export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig(event);
  const data = await $fetch<BlogpostListResponse>(
    `${config.public.url}/api/v1/blogpost`,
    {
      method: "GET",
    },
  );
  return data;
});
