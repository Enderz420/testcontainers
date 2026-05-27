export default defineEventHandler(async (event) => {
  const query = await getQuery(event);
  const config = useRuntimeConfig(event);

  const response = await $fetch(
    `${config.public.url}/api/v1/blogpost/${query}`,
    {
      method: "DELETE",
    },
  );
  return response;
});
