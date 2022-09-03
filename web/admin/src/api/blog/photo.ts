import { defHttp } from '/@/utils/http/axios';
import * as protocol from '/&/photo';

// 获取列表
export const ListPhoto: protocol.ListPhoto = (params) => {
  return defHttp.get<protocol.ListPhotoResponse>({
    url: protocol.route.Photos,
    params,
  });
};

// 获取
export const GetPhoto: protocol.GetPhoto = (params) => {
  return defHttp.get<protocol.Photo>({
    url: protocol.route.PhotoWithId(params.id),
  });
};

// 创建
export const CreatePhoto: protocol.CreatePhoto = (params) => {
  return defHttp.post<protocol.Photo>(
    { url: protocol.route.Photos, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdatePhoto: protocol.UpdatePhoto = (params) => {
  return defHttp.put<protocol.Photo>({
    url: protocol.route.PhotoWithId(params.id),
    data: params.photo,
  });
};

// 删除
export const DeletePhoto: protocol.DeletePhoto = (params) => {
  return defHttp.delete({
    url: protocol.route.PhotoWithId(params.id),
  });
};
