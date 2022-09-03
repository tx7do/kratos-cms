import { defHttp } from '/@/utils/http/axios';
import * as protocol from '/&/menu';

// 获取列表
export const ListMenu: protocol.ListMenu = (params) => {
  return defHttp.get<protocol.ListMenuResponse>({
    url: protocol.route.Menus,
    params,
  });
};

// 获取
export const GetMenu: protocol.GetMenu = (params) => {
  return defHttp.get<protocol.Menu>({
    url: protocol.route.MenuWithId(params.id),
  });
};

// 创建
export const CreateMenu: protocol.CreateMenu = (params) => {
  return defHttp.post<protocol.Menu>(
    { url: protocol.route.Menus, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateMenu: protocol.UpdateMenu = (params) => {
  return defHttp.put<protocol.Menu>({
    url: protocol.route.MenuWithId(params.id),
    data: params.menu,
  });
};

// 删除
export const DeleteMenu: protocol.DeleteMenu = (params) => {
  return defHttp.delete({
    url: protocol.route.MenuWithId(params.id),
  });
};
