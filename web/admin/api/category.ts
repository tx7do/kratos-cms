import * as pagination from './pagination';

export const route = {
  Categories: '/categories',
  CategoryWithId: (id) => `/categories/${id}`,
};

export interface ListCategory {
  (params: pagination.PagingRequest): Promise<ListCategoryResponse>;
}

export interface GetCategory {
  (params: GetCategoryRequest): Promise<Category>;
}

export interface CreateCategory {
  (params: CreateCategoryRequest): Promise<Category>;
}

export interface UpdateCategory {
  (params: UpdateCategoryRequest): Promise<Category>;
}

export interface DeleteCategory {
  (params: DeleteCategoryRequest): Promise<{}>;
}

export interface Category {
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
}

export interface ListCategoryResponse {
  items?: Category[];
  total?: number;
}

export interface GetCategoryRequest {
  id?: number;
}

export interface CreateCategoryRequest {
  category?: Category;
  operatorId?: number;
}

export interface UpdateCategoryRequest {
  id?: number;
  category?: Category;
  operatorId?: number;
}

export interface DeleteCategoryRequest {
  id?: number;
  operatorId?: number;
}
