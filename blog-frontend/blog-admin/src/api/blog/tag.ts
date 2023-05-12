import { defHttp } from '/@/utils/http/axios';

const route = {
  Tags: '/tags',
  TagWithId: (id) => `/tags/${id}`,
};

// 获取列表
export const ListTag: ListTag = (params) => {
  return defHttp.get<ListTagResponse>({
    url: route.Tags,
    params,
  });
};

// 获取
export const GetTag: GetTag = (params) => {
  return defHttp.get<Tag>({
    url: route.TagWithId(params.id),
  });
};

// 创建
export const CreateTag: CreateTag = (params) => {
  return defHttp.post<Tag>(
    { url: route.Tags, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateTag: UpdateTag = (params) => {
  return defHttp.put<Tag>({
    url: route.TagWithId(params.id),
    data: params.tag,
  });
};

// 删除
export const DeleteTag: DeleteTag = (params) => {
  return defHttp.delete({
    url: route.TagWithId(params.id),
  });
};
