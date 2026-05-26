export default defineEventHandler(async (event) => {
  const body = await readBody(event);
  const response = await $fetch("http://localhost:4000/api/v1/user", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: body,
  });
  return response;
});
