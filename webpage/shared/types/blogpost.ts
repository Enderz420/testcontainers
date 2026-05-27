export type Blogpost = {
  id: string;
  title: string;
  content: string;
  created_by: string;
  created_at: Date;
  updated_at: Date;
};

export type BlogpostListResponse = {
  results: Blogpost[];
  metadata: Metadata;
};

export type BlogpostResponse = {
  results: Blogpost;
  metadata: Metadata;
};

export type Metadata = {
  last_seen: string;
  length: number;
};

export type PostBlogpost = {
  title: string;
  content: string;
  created_by: string;
};
