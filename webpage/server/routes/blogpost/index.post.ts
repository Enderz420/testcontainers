export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig(event);
  const data = await readBody(event);
  console.log("request ", data);

  const response = await $fetch<BlogpostResponse>(
    `${config.public.url}/api/v1/blogpost`,
    {
      method: "POST",
      body: data,
    },
  );
  console.log("returning");
  return response;
});
