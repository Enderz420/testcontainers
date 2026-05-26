export function useBlogpost() {
  const getAllBlogposts = async (): Promise<BlogpostListResponse> => {
    const { data, error } = await useFetch<BlogpostListResponse>("/blogpost", {
      method: "GET",
    });

    console.log("error:", error.value);
    console.log("data:", data.value);

    if (error.value) throw error.value;
    if (!data.value) throw new Error("No data returned");

    return data.value;
  };

  const createBlogpost = async (
    body: PostBlogpost,
  ): Promise<BlogpostResponse> => {
    console.log("Inserting data");
    const { data, error } = await useFetch<BlogpostResponse>(
      "http://localhost:3000/blogpost",
      {
        method: "POST",
        body: body,
        headers: { "Content-Type": "application/json" },
      },
    );
    console.log("error:", error.value);
    console.log("data:", data.value);
    if (error.value) throw error.value;
    if (!data.value) throw new Error("No data returned");
    return data.value;
  };

  const deleteBlogpost = async (id: string): Promise<string> => {
    const { status, error } = await useFetch(`/blogpost/${id}`, {
      method: "DELETE",
    });

    console.log("error:", error.value);
    console.log("data:", status.value);
    if (error.value) throw error.value;

    return status.value;
  };
  return {
    deleteBlogpost,
    getAllBlogposts,
    createBlogpost,
  };
}
