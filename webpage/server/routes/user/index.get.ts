export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig(event);
  const response = await $fetch(`${config.public.url}/api/v1/user`, {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  });
  return response;
});
