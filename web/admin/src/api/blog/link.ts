import { defHttp } from '/@/utils/http/axios';
import * as protocol from '/&/link';

// 获取列表
export const ListLink: protocol.ListLink = (params) => {
  return defHttp.get<protocol.ListLinkResponse>({
    url: protocol.route.Links,
    params,
  });
};

// 获取
export const GetLink: protocol.GetLink = (params) => {
  return defHttp.get<protocol.Link>({
    url: protocol.route.LinkWithId(params.id),
  });
};

// 创建
export const CreateLink: protocol.CreateLink = (params) => {
  return defHttp.post<protocol.Link>(
    { url: protocol.route.Links, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateLink: protocol.UpdateLink = (params) => {
  return defHttp.put<protocol.Link>({
    url: protocol.route.LinkWithId(params.id),
    data: params.link,
  });
};

// 删除
export const DeleteLink: protocol.DeleteLink = (params) => {
  return defHttp.delete({
    url: protocol.route.LinkWithId(params.id),
  });
};
