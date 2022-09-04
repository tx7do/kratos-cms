import CommentService from '../service/CommentService';
import * as protocol from '../../../../api/comment';

class CommentController {
  private service: CommentService = new CommentService();

  ListComment = async (ctx) => {
    await this.service.ListComment(ctx);
  };

  GetComment = async (ctx) => {
    await this.service.GetComment(ctx);
  };

  CreateComment = async (ctx) => {
    await this.service.CreateComment(ctx);
  };

  UpdateComment = async (ctx) => {
    await this.service.UpdateComment(ctx);
  };

  DeleteComment = async (ctx) => {
    await this.service.DeleteComment(ctx);
  };
}

export const commentController = new CommentController();

export const Routes = [
  {
    path: protocol.route.Comments,
    method: 'get',
    action: commentController.ListComment,
  },
  {
    path: protocol.route.Comments,
    method: 'get',
    action: commentController.GetComment,
  },
  {
    path: protocol.route.Comments,
    method: 'post',
    action: commentController.CreateComment,
  },
  {
    path: protocol.route.Comments,
    method: 'put',
    action: commentController.UpdateComment,
  },
  {
    path: protocol.route.Comments,
    method: 'delete',
    action: commentController.DeleteComment,
  },
];
