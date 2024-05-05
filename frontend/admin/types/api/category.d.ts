interface ListCategory {
  (params: PagingRequest): Promise<ListCategoryResponse>;
}

interface GetCategory {
  (params: GetCategoryRequest): Promise<Category>;
}

interface CreateCategory {
  (params: CreateCategoryRequest): Promise<Category>;
}

interface UpdateCategory {
  (params: UpdateCategoryRequest): Promise<Category>;
}

interface DeleteCategory {
  (params: DeleteCategoryRequest): Promise<{}>;
}

interface Category {
  id?: number;
  parentId?: number;
  name?: string;
  slug?: string;
  description?: string;
  thumbnail?: string;
  password?: string;
  fullPath?: string;
  priority?: number;
  createTime?: string;
  updateTime?: string;
  postCount?: number;
  children?: Category[];
}

interface ListCategoryResponse {
  items?: Category[];
  total?: number;
}

interface GetCategoryRequest {
  id?: number;
}

interface CreateCategoryRequest {
  category?: Category;
  operatorId?: number;
}

interface UpdateCategoryRequest {
  id?: number;
  category?: Category;
  operatorId?: number;
}

interface DeleteCategoryRequest {
  id?: number;
  operatorId?: number;
}
