import { Result } from '../utils/utils';
import { createAttachment, createAttachments } from '../utils/mock';

export default class AttachmentService {
  ListAttachment = async (ctx) => {
    const total = 20;
    ctx.body = Result.success({
      items: createAttachments(total),
      total: total,
    });
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
