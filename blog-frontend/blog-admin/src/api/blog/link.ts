import { defHttp } from '/@/utils/http/axios';

const route = {
  Links: '/links',
  LinkWithId: (id) => `/links/${id}`,
};

// 获取列表
export const ListLink: ListLink = (params) => {
  return defHttp.get<ListLinkResponse>({
    url: route.Links,
    params,
  });
};

// 获取
export const GetLink: GetLink = (params) => {
  return defHttp.get<Link>({
    url: route.LinkWithId(params.id),
  });
};

// 创建
export const CreateLink: CreateLink = (params) => {
  return defHttp.post<Link>(
    { url: route.Links, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateLink: UpdateLink = (params) => {
  return defHttp.put<Link>({
    url: route.LinkWithId(params.id),
    data: params.link,
  });
};

// 删除
export const DeleteLink: DeleteLink = (params) => {
  return defHttp.delete({
    url: route.LinkWithId(params.id),
  });
};
