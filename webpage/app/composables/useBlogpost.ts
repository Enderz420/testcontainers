export default async function useBlogpost() {
  type PostBlogpost = {
    Title: string;
    Content: string;
    CreatedBy: string;
  };

  const getAllBlogposts = async () => {
    const { data } = await useFetch("/blogpost", {
      method: "GET",
    });
    return data.value;
  };

  const createBlogpost = async (body: PostBlogpost) => {
    const { data } = await useFetch("/blogpost", {
      method: "POST",
      body: body,
    });

    return data.value;
  };
  return {
    getAllBlogposts,
    createBlogpost,
  };
}
