import { Result } from '../utils/utils';
import { createAttachment, createAttachments } from '../utils/mock';
import * as pagination from '/&/pagination';

export default class AttachmentService {
  ListAttachment = async (ctx) => {
    const request: pagination.PagingRequest = ctx.query;
    ctx.body = Result.pageSuccess(
      request.page,
      request.pageSize,
      createAttachments(request.pageSize),
    );
  };

  GetAttachment = async (ctx) => {
    ctx.body = Result.success(createAttachment());
  };

  CreateAttachment = async (ctx) => {
    ctx.body = Result.success(createAttachment());
  };

  UpdateAttachment = async (ctx) => {
    ctx.body = Result.success(createAttachment());
  };

  DeleteAttachment = async (ctx) => {
    ctx.body = Result.success({});
  };
}
