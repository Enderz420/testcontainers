export default defineEventHandler(async () => {
  const data = await $fetch("http://localhost:4000/api/v1/blogpost", {
    method: "GET",
  });
  return data;
});
