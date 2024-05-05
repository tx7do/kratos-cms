interface ListTag {
  (params: PagingRequest): Promise<ListTagResponse>;
}

interface GetTag {
  (params: GetTagRequest): Promise<Tag>;
}

interface CreateTag {
  (params: CreateTagRequest): Promise<Tag>;
}

interface UpdateTag {
  (params: UpdateTagRequest): Promise<Tag>;
}

interface DeleteTag {
  (params: DeleteTagRequest): Promise<{}>;
}

interface Tag {
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

interface ListTagResponse {
  items?: Tag[];
  total?: number;
}

interface GetTagRequest {
  id?: number;
}

interface CreateTagRequest {
  tag?: Tag;
  operatorId?: number;
}

interface UpdateTagRequest {
  id?: number;
  tag?: Tag;
  operatorId?: number;
}

interface DeleteTagRequest {
  id?: number;
  operatorId?: number;
}
