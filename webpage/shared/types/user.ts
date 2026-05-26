import type { Metadata } from "./blogpost";

export type User = {
  id: string;
  username: string;
  email: string;
  created_at: Date;
  updated_at: Date;
};
export type PostUser = {
  username: string;
  email: string;
};
export type UserResponse = {
  data: User;
};

export type UserListResponse = {
  data: User[];
  metadata: Metadata;
};
