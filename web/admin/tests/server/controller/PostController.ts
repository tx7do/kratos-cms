import PostService from '../service/PostService';

class PostController {
  private service: PostService = new PostService();
}

export const postController = new PostController();

export const Routes = [];
