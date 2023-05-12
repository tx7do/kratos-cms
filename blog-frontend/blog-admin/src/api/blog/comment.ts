import { defHttp } from '/@/utils/http/axios';

const route = {
  Comments: '/comments',
  CommentWithId: (id) => `/comments/${id}`,
};

// 获取列表
export const ListComment: ListComment = (params) => {
  return defHttp.get<ListCommentResponse>({
    url: route.Comments,
    params,
  });
};

// 获取
export const GetComment: GetComment = (params) => {
  return defHttp.get<Comment>({
    url: route.CommentWithId(params.id),
  });
};

// 创建
export const CreateComment: CreateComment = (params) => {
  return defHttp.post<Comment>(
    { url: route.Comments, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateComment: UpdateComment = (params) => {
  return defHttp.put<Comment>({
    url: route.CommentWithId(params.id),
    data: params.comment,
  });
};

// 删除
export const DeleteComment: DeleteComment = (params) => {
  return defHttp.delete({
    url: route.CommentWithId(params.id),
  });
};
