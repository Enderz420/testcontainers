export function useBlogpost() {
  const getAllBlogposts = async (): Promise<BlogpostListResponse> => {
    const { data } = await useFetch<BlogpostListResponse>("/blogpost", {
      method: "GET",
    });
    if (!data.value) throw new Error("No data returned");

    return data.value;
  };

  const createBlogpost = async (
    body: PostBlogpost,
  ): Promise<BlogpostResponse> => {
    console.log("Inserting data");
    const { data, error } = await useFetch<BlogpostResponse>("/blogpost", {
      method: "POST",
      body: body,
    });
    console.log("error:", error.value);
    console.log("data:", data.value);
    if (error.value) throw error.value;
    if (!data.value) throw new Error("No data returned");
    return data.value;
  };

  const deleteBlogpost = async (id: string): Promise<string> => {
    const { status } = await useFetch(`/blogpost/${id}`, {
      method: "DELETE",
    });

    if (!status.value) {
      throw new Error(status.value);
    }

    return status.value;
  };
  return {
    deleteBlogpost,
    getAllBlogposts,
    createBlogpost,
  };
}
