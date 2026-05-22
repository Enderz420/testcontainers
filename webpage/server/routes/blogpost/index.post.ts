export default defineEventHandler(async (event) => {
  const data = await readBody(event);
  console.log("request ", data);

  const response = await $fetch<BlogpostResponse>(
    "http://localhost:4000/api/v1/blogpost",
    {
      method: "POST",
      body: data,
    },
  );
  console.log("returning");
  return response;
});
