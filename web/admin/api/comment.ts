import * as pagination from './pagination';

export const route = {
  Comments: '/comments',
  CommentWithId: (id) => `/comments/${id}`,
};

export interface ListComment {
  (params: pagination.PagingRequest): Promise<ListCommentResponse>;
}

export interface GetComment {
  (params: GetCommentRequest): Promise<Comment>;
}

export interface CreateComment {
  (params: CreateCommentRequest): Promise<Comment>;
}

export interface UpdateComment {
  (params: UpdateCommentRequest): Promise<Comment>;
}

export interface DeleteComment {
  (params: DeleteCommentRequest): Promise<{}>;
}

export interface Comment {
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

export interface ListCommentResponse {
  items?: Comment[];
  total?: number;
}

export interface GetCommentRequest {
  id?: number;
}

export interface CreateCommentRequest {
  comment?: Comment;
  operatorId?: number;
}

export interface UpdateCommentRequest {
  id?: number;
  comment?: Comment;
  operatorId?: number;
}

export interface DeleteCommentRequest {
  id?: number;
  operatorId?: number;
}
