export type Blogpost = {
  id: string;
  title: string;
  content: string;
  createdBy: string;
  createdAt: Date;
  updatedAt: Date;
};

export type BlogpostListResponse = {
  data: Blogpost[];
  metadata: Metadata;
};

export type BlogpostResponse = {
  data: Blogpost;
  metadata: Metadata;
};

export type Metadata = {
  LastSeen: string;
  Length: number;
};

export type PostBlogpost = {
  title: string;
  content: string;
  created_by: string;
};
