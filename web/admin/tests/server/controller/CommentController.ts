import CommentService from '../service/CommentService';

class CommentController {
  private service: CommentService = new CommentService();
}

export const commentController = new CommentController();

export const Routes = [];
