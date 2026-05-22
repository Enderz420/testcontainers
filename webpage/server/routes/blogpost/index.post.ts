export default defineEventHandler(async (event) => {
  const data = await readBody(event);
  const response = await $fetch("http://localhost:4000/api/v1/blogpost", {
    method: "POST",
    body: data,
  });
  return response;
});
