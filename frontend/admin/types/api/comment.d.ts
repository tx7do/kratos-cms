interface ListComment {
  (params: PagingRequest): Promise<ListCommentResponse>;
}

interface GetComment {
  (params: GetCommentRequest): Promise<Comment>;
}

interface CreateComment {
  (params: CreateCommentRequest): Promise<Comment>;
}

interface UpdateComment {
  (params: UpdateCommentRequest): Promise<Comment>;
}

interface DeleteComment {
  (params: DeleteCommentRequest): Promise<{}>;
}

interface Comment {
  id?: number;
  author?: string;
  email?: string;
  ipAddress?: string;
  authorUrl?: string;
  gravatarMd5?: string;
  content?: string;
  status?: number;
  userAgent?: string;
  parentId?: number;
  isAdmin?: boolean;
  allowNotification?: boolean;
  avatar?: string;
  createTime?: string;
  updateTime?: string;
}

interface ListCommentResponse {
  items?: Comment[];
  total?: number;
}

interface GetCommentRequest {
  id?: number;
}

interface CreateCommentRequest {
  comment?: Comment;
  operatorId?: number;
}

interface UpdateCommentRequest {
  id?: number;
  comment?: Comment;
  operatorId?: number;
}

interface DeleteCommentRequest {
  id?: number;
  operatorId?: number;
}
