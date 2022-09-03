import * as pagination from './pagination';

export const route = {
  Posts: '/posts',
  PostWithId: (id) => `/posts/${id}`,
};

export interface ListPost {
  (params: pagination.PagingRequest): Promise<ListPostResponse>;
}

export interface GetPost {
  (params: GetPostRequest): Promise<Post>;
}

export interface CreatePost {
  (params: CreatePostRequest): Promise<Post>;
}

export interface UpdatePost {
  (params: UpdatePostRequest): Promise<Post>;
}

export interface DeletePost {
  (params: DeletePostRequest): Promise<{}>;
}

export interface Post {
  id?: number;
  title?: string;
  status?: number;
  slug?: string;
  editorType?: number;
  metaKeywords?: string;
  metaDescription?: string;
  fullPath?: string;
  summary?: string;
  thumbnail?: string;
  password?: string;
  template?: string;
  content?: string;
  originalContent?: string;
  visits?: number;
  topPriority?: number;
  likes?: number;
  wordCount?: number;
  commentCount?: number;
  disallowComment?: boolean;
  inProgress?: boolean;
  createTime?: string;
  updateTime?: string;
  editTime?: string;
}

export interface ListPostResponse {
  items?: Post[];
  total?: number;
}

export interface GetPostRequest {
  id?: number;
}

export interface CreatePostRequest {
  post?: Post;
  operatorId?: number;
}

export interface UpdatePostRequest {
  id?: number;
  post?: Post;
  operatorId?: number;
}

export interface DeletePostRequest {
  id?: number;
  operatorId?: number;
}
