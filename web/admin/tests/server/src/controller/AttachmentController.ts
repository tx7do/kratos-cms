import AttachmentService from '../service/AttachmentService';
import * as protocol from '/&/attachment';

class AttachmentController {
  private service: AttachmentService = new AttachmentService();

  ListAttachment = async (ctx) => {
    await this.service.ListAttachment(ctx);
  };

  GetAttachment = async (ctx) => {
    await this.service.GetAttachment(ctx);
  };

  CreateAttachment = async (ctx) => {
    await this.service.CreateAttachment(ctx);
  };

  UpdateAttachment = async (ctx) => {
    await this.service.UpdateAttachment(ctx);
  };

  DeleteAttachment = async (ctx) => {
    await this.service.DeleteAttachment(ctx);
  };
}

export const attachmentController = new AttachmentController();

export const Routes = [
  {
    path: protocol.route.Attachments,
    method: 'get',
    action: attachmentController.ListAttachment,
  },
  {
    path: protocol.route.Attachments,
    method: 'get',
    action: attachmentController.GetAttachment,
  },
  {
    path: protocol.route.Attachments,
    method: 'post',
    action: attachmentController.CreateAttachment,
  },
  {
    path: protocol.route.Attachments,
    method: 'put',
    action: attachmentController.UpdateAttachment,
  },
  {
    path: protocol.route.Attachments,
    method: 'delete',
    action: attachmentController.DeleteAttachment,
  },
];
