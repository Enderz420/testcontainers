export default defineEventHandler(async () => {
  const response = await $fetch("http://localhost:4000/api/v1/user", {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  });
  return response;
});
