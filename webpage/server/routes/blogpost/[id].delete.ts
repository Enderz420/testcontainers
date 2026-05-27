export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");

  const config = useRuntimeConfig(event);

  const response = await $fetch(`${config.public.url}/api/v1/blogpost/${id}`, {
    method: "DELETE",
  });
  return response;
});
