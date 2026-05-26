export default defineEventHandler(async (event) => {
  const query = await getQuery(event);

  const response = await $fetch(
    `http://localhost:4000/api/v1/blogpost/${query}`,
    {
      method: "DELETE",
    },
  );
  return response;
});
