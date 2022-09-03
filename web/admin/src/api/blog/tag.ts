import { defHttp } from '/@/utils/http/axios';
import * as protocol from '/&/tag';

// 获取列表
export const ListTag: protocol.ListTag = (params) => {
  return defHttp.get<protocol.ListTagResponse>({
    url: protocol.route.Tags,
    params,
  });
};

// 获取
export const GetTag: protocol.GetTag = (params) => {
  return defHttp.get<protocol.Tag>({
    url: protocol.route.TagWithId(params.id),
  });
};

// 创建
export const CreateTag: protocol.CreateTag = (params) => {
  return defHttp.post<protocol.Tag>(
    { url: protocol.route.Tags, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateTag: protocol.UpdateTag = (params) => {
  return defHttp.put<protocol.Tag>({
    url: protocol.route.TagWithId(params.id),
    data: params.tag,
  });
};

// 删除
export const DeleteTag: protocol.DeleteTag = (params) => {
  return defHttp.delete({
    url: protocol.route.TagWithId(params.id),
  });
};
