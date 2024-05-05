import { defHttp } from '/@/utils/http/axios';

const route = {
  Photos: '/photos',
  PhotoWithId: (id) => `/photos/${id}`,
};

// 获取列表
export const ListPhoto: ListPhoto = (params) => {
  return defHttp.get<ListPhotoResponse>({
    url: route.Photos,
    params,
  });
};

// 获取
export const GetPhoto: GetPhoto = (params) => {
  return defHttp.get<Photo>({
    url: route.PhotoWithId(params.id),
  });
};

// 创建
export const CreatePhoto: CreatePhoto = (params) => {
  return defHttp.post<Photo>(
    { url: route.Photos, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdatePhoto: UpdatePhoto = (params) => {
  return defHttp.put<Photo>({
    url: route.PhotoWithId(params.id),
    data: params.photo,
  });
};

// 删除
export const DeletePhoto: DeletePhoto = (params) => {
  return defHttp.delete({
    url: route.PhotoWithId(params.id),
  });
};
