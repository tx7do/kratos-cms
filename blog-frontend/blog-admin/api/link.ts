import * as pagination from './pagination';

export const route = {
  Links: '/links',
  LinkWithId: (id) => `/links/${id}`,
};

export interface ListLink {
  (params: pagination.PagingRequest): Promise<ListLinkResponse>;
}

export interface GetLink {
  (params: GetLinkRequest): Promise<Link>;
}

export interface CreateLink {
  (params: CreateLinkRequest): Promise<Link>;
}

export interface UpdateLink {
  (params: UpdateLinkRequest): Promise<Link>;
}

export interface DeleteLink {
  (params: DeleteLinkRequest): Promise<{}>;
}

export interface Link {
  id?: number;
  name?: string;
  url?: string;
  logo?: string;
  description?: string;
  team?: string;
  priority?: number;
  createTime?: string;
  updateTime?: string;
}

export interface ListLinkResponse {
  items?: Link[];
  total?: number;
}

export interface GetLinkRequest {
  id?: number;
}

export interface CreateLinkRequest {
  link?: Link;
  operatorId?: number;
}

export interface UpdateLinkRequest {
  id?: number;
  link?: Link;
  operatorId?: number;
}

export interface DeleteLinkRequest {
  id?: number;
  operatorId?: number;
}
