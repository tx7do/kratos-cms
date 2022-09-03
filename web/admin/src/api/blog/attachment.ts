import { defHttp } from '/@/utils/http/axios';
import * as protocol from '/&/attachment';

// 获取列表
export const ListAttachment: protocol.ListAttachment = (params) => {
  return defHttp.get<protocol.ListAttachmentResponse>({
    url: protocol.route.Attachments,
    params,
  });
};

// 获取
export const GetAttachment: protocol.GetAttachment = (params) => {
  return defHttp.get<protocol.Attachment>({
    url: protocol.route.AttachmentWithId(params.id),
  });
};

// 创建
export const CreateAttachment: protocol.CreateAttachment = (params) => {
  return defHttp.post<protocol.Attachment>(
    { url: protocol.route.Attachments, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 更新
export const UpdateAttachment: protocol.UpdateAttachment = (params) => {
  return defHttp.put<protocol.Attachment>({
    url: protocol.route.AttachmentWithId(params.id),
    data: params.attachment,
  });
};

// 删除
export const DeleteAttachment: protocol.DeleteAttachment = (params) => {
  return defHttp.delete({
    url: protocol.route.AttachmentWithId(params.id),
  });
};
