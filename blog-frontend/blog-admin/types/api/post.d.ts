interface ListPost {
  (params: PagingRequest): Promise<ListPostResponse>;
}

interface GetPost {
  (params: GetPostRequest): Promise<Post>;
}

interface CreatePost {
  (params: CreatePostRequest): Promise<Post>;
}

interface UpdatePost {
  (params: UpdatePostRequest): Promise<Post>;
}

interface DeletePost {
  (params: DeletePostRequest): Promise<{}>;
}

interface Post {
  id?: number;
  title?: string;
  status?: string;
  slug?: string;
  editorType?: string;
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
  tags?: Tag[];
  categories?: Category[];
}

interface ListPostResponse {
  items?: Post[];
  total?: number;
}

interface GetPostRequest {
  id?: number;
}

interface CreatePostRequest {
  post?: Post;
  operatorId?: number;
}

interface UpdatePostRequest {
  id?: number;
  post?: Post;
  operatorId?: number;
}

interface DeletePostRequest {
  id?: number;
  operatorId?: number;
}
