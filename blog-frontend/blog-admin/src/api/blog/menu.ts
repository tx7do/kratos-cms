import { defHttp } from '/@/utils/http/axios';

const route = {
  Menus: '/menus',
  MenuWithId: (id) => `/menus/${id}`,
};

// 获取列表
export const ListMenu: ListMenu = (params) => {
  return defHttp.get<ListMenuResponse>({
    url: route.Menus,
    params,
  });
};

// 获取
export const GetMenu: GetMenu = (params) => {
  return defHttp.get<Menu>({
    url: route.MenuWithId(params.id),
  });
};

// 创建
export const CreateMenu: CreateMenu = (params) => {
  return defHttp.post<Menu>(
    { url: route.Menus, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateMenu: UpdateMenu = (params) => {
  return defHttp.put<Menu>({
    url: route.MenuWithId(params.id),
    data: params.menu,
  });
};

// 删除
export const DeleteMenu: DeleteMenu = (params) => {
  return defHttp.delete({
    url: route.MenuWithId(params.id),
  });
};
