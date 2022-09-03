import AttachmentService from '../service/AttachmentService';

class AttachmentController {
  private service: AttachmentService = new AttachmentService();
}

export const attachmentController = new AttachmentController();

export const Routes = [];
