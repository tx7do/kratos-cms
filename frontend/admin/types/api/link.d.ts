interface ListLink {
  (params: PagingRequest): Promise<ListLinkResponse>;
}

interface GetLink {
  (params: GetLinkRequest): Promise<Link>;
}

interface CreateLink {
  (params: CreateLinkRequest): Promise<Link>;
}

interface UpdateLink {
  (params: UpdateLinkRequest): Promise<Link>;
}

interface DeleteLink {
  (params: DeleteLinkRequest): Promise<{}>;
}

interface Link {
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

interface ListLinkResponse {
  items?: Link[];
  total?: number;
}

interface GetLinkRequest {
  id?: number;
}

interface CreateLinkRequest {
  link?: Link;
  operatorId?: number;
}

interface UpdateLinkRequest {
  id?: number;
  link?: Link;
  operatorId?: number;
}

interface DeleteLinkRequest {
  id?: number;
  operatorId?: number;
}
