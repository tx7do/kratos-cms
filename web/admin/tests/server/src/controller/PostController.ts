import PostService from '../service/PostService';
import * as protocol from '/&/post';

class PostController {
  private service: PostService = new PostService();

  ListPost = async (ctx) => {
    await this.service.ListPost(ctx);
  };

  GetPost = async (ctx) => {
    await this.service.GetPost(ctx);
  };

  CreatePost = async (ctx) => {
    await this.service.CreatePost(ctx);
  };

  UpdatePost = async (ctx) => {
    await this.service.UpdatePost(ctx);
  };

  DeletePost = async (ctx) => {
    await this.service.DeletePost(ctx);
  };
}

export const postController = new PostController();

export const Routes = [
  {
    path: protocol.route.Posts,
    method: 'get',
    action: postController.ListPost,
  },
  {
    path: protocol.route.Posts,
    method: 'get',
    action: postController.GetPost,
  },
  {
    path: protocol.route.Posts,
    method: 'post',
    action: postController.CreatePost,
  },
  {
    path: protocol.route.Posts,
    method: 'put',
    action: postController.UpdatePost,
  },
  {
    path: protocol.route.Posts,
    method: 'delete',
    action: postController.DeletePost,
  },
];
