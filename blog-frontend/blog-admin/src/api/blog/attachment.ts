import { defHttp } from '/@/utils/http/axios';

const route = {
  Attachments: '/attachments',
  AttachmentWithId: (id) => `/attachments/${id}`,
};

// 获取列表
export const ListAttachment: ListAttachment = (params) => {
  return defHttp.get<ListAttachmentResponse>({
    url: route.Attachments,
    params,
  });
};

// 获取
export const GetAttachment: GetAttachment = (params) => {
  return defHttp.get<Attachment>({
    url: route.AttachmentWithId(params.id),
  });
};

// 创建
export const CreateAttachment: CreateAttachment = (params) => {
  return defHttp.post<Attachment>(
    {url: route.Attachments, params},
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateAttachment: UpdateAttachment = (params) => {
  return defHttp.put<Attachment>({
    url: route.AttachmentWithId(params.id),
    data: params.attachment,
  });
};

// 删除
export const DeleteAttachment: DeleteAttachment = (params) => {
  return defHttp.delete({
    url: route.AttachmentWithId(params.id),
  });
};
