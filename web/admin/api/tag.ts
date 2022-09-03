import * as pagination from './pagination';

export const route = {
  Tags: '/tags',
  TagWithId: (id) => `/tags/${id}`,
};

export interface ListTag {
  (params: pagination.PagingRequest): Promise<ListTagResponse>;
}

export interface GetTag {
  (params: GetTagRequest): Promise<Tag>;
}

export interface CreateTag {
  (params: CreateTagRequest): Promise<Tag>;
}

export interface UpdateTag {
  (params: UpdateTagRequest): Promise<Tag>;
}

export interface DeleteTag {
  (params: DeleteTagRequest): Promise<{}>;
}

export interface Tag {
  id?: number;
  name?: string;
  slug?: string;
  slugName?: string;
  color?: string;
  thumbnail?: string;
  createTime?: string;
  updateTime?: string;
  postCount?: number;
  fullPath?: string;
}

export interface ListTagResponse {
  items?: Tag[];
  total?: number;
}

export interface GetTagRequest {
  id?: number;
}

export interface CreateTagRequest {
  tag?: Tag;
  operatorId?: number;
}

export interface UpdateTagRequest {
  id?: number;
  tag?: Tag;
  operatorId?: number;
}

export interface DeleteTagRequest {
  id?: number;
  operatorId?: number;
}
