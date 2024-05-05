interface ListAttachment {
  (params: PagingRequest): Promise<ListAttachmentResponse>;
}

interface GetAttachment {
  (params: GetAttachmentRequest): Promise<Attachment>;
}

interface CreateAttachment {
  (params: CreateAttachmentRequest): Promise<Attachment>;
}

interface UpdateAttachment {
  (params: UpdateAttachmentRequest): Promise<Attachment>;
}

interface DeleteAttachment {
  (params: DeleteAttachmentRequest): Promise<{}>;
}

interface Attachment {
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

interface ListAttachmentResponse {
  items?: Attachment[];
  total?: number;
}

interface GetAttachmentRequest {
  id?: number;
}

interface CreateAttachmentRequest {
  attachment?: Attachment;
  operatorId?: number;
}

interface UpdateAttachmentRequest {
  id?: number;
  attachment?: Attachment;
  operatorId?: number;
}

interface DeleteAttachmentRequest {
  id?: number;
  operatorId?: number;
}
