export default defineEventHandler(async (event) => {
  const body = await readBody(event);
  const config = useRuntimeConfig(event);
  const response = await $fetch(`${config.public.url}/api/v1/user`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: body,
  });
  return response;
});
