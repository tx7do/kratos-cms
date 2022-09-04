import * as pagination from './pagination';

export const route = {
  Attachments: '/attachments',
  AttachmentWithId: (id) => `/attachments/${id}`,
};

export interface ListAttachment {
  (params: pagination.PagingRequest): Promise<ListAttachmentResponse>;
}

export interface GetAttachment {
  (params: GetAttachmentRequest): Promise<Attachment>;
}

export interface CreateAttachment {
  (params: CreateAttachmentRequest): Promise<Attachment>;
}

export interface UpdateAttachment {
  (params: UpdateAttachmentRequest): Promise<Attachment>;
}

export interface DeleteAttachment {
  (params: DeleteAttachmentRequest): Promise<{}>;
}

export interface Attachment {
  id?: number;
  name?: string;
  path?: string;
  fileKey?: string;
  thumbPath?: string;
  mediaType?: string;
  suffix?: string;
  width?: number;
  height?: number;
  size?: number;
  type?: string;
  createTime?: string;
}

export interface ListAttachmentResponse {
  items?: Attachment[];
  total?: number;
}

export interface GetAttachmentRequest {
  id?: number;
}

export interface CreateAttachmentRequest {
  attachment?: Attachment;
  operatorId?: number;
}

export interface UpdateAttachmentRequest {
  id?: number;
  attachment?: Attachment;
  operatorId?: number;
}

export interface DeleteAttachmentRequest {
  id?: number;
  operatorId?: number;
}
