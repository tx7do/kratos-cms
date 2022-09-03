import { defHttp } from '/@/utils/http/axios';
import * as protocol from '/&/comment';

// 获取列表
export const ListComment: protocol.ListComment = (params) => {
  return defHttp.get<protocol.ListCommentResponse>({
    url: protocol.route.Comments,
    params,
  });
};

// 获取
export const GetComment: protocol.GetComment = (params) => {
  return defHttp.get<protocol.Comment>({
    url: protocol.route.CommentWithId(params.id),
  });
};

// 创建
export const CreateComment: protocol.CreateComment = (params) => {
  return defHttp.post<protocol.Comment>(
    { url: protocol.route.Comments, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateComment: protocol.UpdateComment = (params) => {
  return defHttp.put<protocol.Comment>({
    url: protocol.route.CommentWithId(params.id),
    data: params.comment,
  });
};

// 删除
export const DeleteComment: protocol.DeleteComment = (params) => {
  return defHttp.delete({
    url: protocol.route.CommentWithId(params.id),
  });
};
