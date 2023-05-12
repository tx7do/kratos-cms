import { defHttp } from '/@/utils/http/axios';

const route = {
  Categories: '/categories',
  CategoryWithId: (id) => `/categories/${id}`,
};

// 获取列表
export const ListCategory: ListCategory = (params) => {
  return defHttp.get<ListCategoryResponse>({
    url: route.Categories,
    params,
  });
};

// 获取
export const GetCategory: GetCategory = (params) => {
  return defHttp.get<Category>({
    url: route.CategoryWithId(params.id),
  });
};

// 创建
export const CreateCategory: CreateCategory = (params) => {
  return defHttp.post<Category>(
    { url: route.Categories, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateCategory: UpdateCategory = (params) => {
  return defHttp.put<Category>({
    url: route.CategoryWithId(params.id),
    data: params.category,
  });
};

// 删除
export const DeleteCategory: DeleteCategory = (params) => {
  return defHttp.delete({
    url: route.CategoryWithId(params.id),
  });
};
