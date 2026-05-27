/**
 * @name useBlogpost
 * @description useBlogpost Composable to enable CRUD related operations
 *
 * @returns CRUD operations related to blogpost
 */
export function useBlogpost() {
  /**
   * @name getAllBlogposts
   * @description Fetches all blogposts
   * @example ```
   *  const { getAllBlogposts } = useBlogpost();
   *  const response = await getAllBlogposts();
   *  console.log(response)
   * ```
   * @returns Promise<BlogpostListResponse>
   */
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

  /**
   * @name createBlogpost
   * @description Creates a blogpost
   * @param body: PostBlogpost
   * @returns
   * Promise<BlogpostResponse>
   */
  const createBlogpost = async (
    body: PostBlogpost,
  ): Promise<BlogpostResponse> => {
    console.log("Inserting data");
    const { data, error } = await useFetch<BlogpostResponse>("/blogpost", {
      method: "POST",
      body: body,
      headers: { "Content-Type": "application/json" },
    });
    console.log("error:", error.value);
    console.log("data:", data.value);
    if (error.value) throw error.value;
    if (!data.value) throw new Error("No data returned");
    return data.value;
  };
  /**
   * @name deleteBlogpost
   * @description Delets a blogpost
   * @param id: string
   * @returns Promise<BlogpostResponse>
   */
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
