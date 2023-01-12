import { defHttp } from '/@/utils/http/axios';
import * as protocol from '/&/category';

// 获取列表
export const ListCategory: protocol.ListCategory = (params) => {
  return defHttp.get<protocol.ListCategoryResponse>({
    url: protocol.route.Categories,
    params,
  });
};

// 获取
export const GetCategory: protocol.GetCategory = (params) => {
  return defHttp.get<protocol.Category>({
    url: protocol.route.CategoryWithId(params.id),
  });
};

// 创建
export const CreateCategory: protocol.CreateCategory = (params) => {
  return defHttp.post<protocol.Category>(
    { url: protocol.route.Categories, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateCategory: protocol.UpdateCategory = (params) => {
  return defHttp.put<protocol.Category>({
    url: protocol.route.CategoryWithId(params.id),
    data: params.category,
  });
};

// 删除
export const DeleteCategory: protocol.DeleteCategory = (params) => {
  return defHttp.delete({
    url: protocol.route.CategoryWithId(params.id),
  });
};
