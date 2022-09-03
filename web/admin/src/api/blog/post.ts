import { defHttp } from '/@/utils/http/axios';
import * as protocol from '/&/post';

// 获取列表
export const ListPost: protocol.ListPost = (params) => {
  return defHttp.get<protocol.ListPostResponse>({
    url: protocol.route.Posts,
    params,
  });
};

// 获取
export const GetPost: protocol.GetPost = (params) => {
  return defHttp.get<protocol.Post>({
    url: protocol.route.PostWithId(params.id),
  });
};

// 创建
export const CreatePost: protocol.CreatePost = (params) => {
  return defHttp.post<protocol.Post>(
    { url: protocol.route.Posts, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdatePost: protocol.UpdatePost = (params) => {
  return defHttp.put<protocol.Post>({
    url: protocol.route.PostWithId(params.id),
    data: params.post,
  });
};

// 删除
export const DeletePost: protocol.DeletePost = (params) => {
  return defHttp.delete({
    url: protocol.route.PostWithId(params.id),
  });
};
