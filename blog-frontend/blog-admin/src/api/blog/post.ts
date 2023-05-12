import { defHttp } from '/@/utils/http/axios';

const route = {
  Posts: '/posts',
  PostWithId: (id) => `/posts/${id}`,
};

// 获取列表
export const ListPost: ListPost = (params) => {
  return defHttp.get<ListPostResponse>({
    url: route.Posts,
    params,
  });
};

// 获取
export const GetPost: GetPost = (params) => {
  return defHttp.get<Post>({
    url: route.PostWithId(params.id),
  });
};

// 创建
export const CreatePost: CreatePost = (params) => {
  return defHttp.post<Post>(
    { url: route.Posts, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdatePost: UpdatePost = (params) => {
  return defHttp.put<Post>({
    url: route.PostWithId(params.id),
    data: params.post,
  });
};

// 删除
export const DeletePost: DeletePost = (params) => {
  return defHttp.delete({
    url: route.PostWithId(params.id),
  });
};
