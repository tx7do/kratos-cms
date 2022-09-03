import PostService from '../service/PostService';
import * as protocol from "../../../api/post";
import {userController} from "./UserController";

class PostController {
  private service: PostService = new PostService();

  ListPost = async (ctx) => {
    ctx.body = await this.service.ListPost(ctx);
  };

  GetPost = async (ctx) => {
    ctx.body = await this.service.GetPost(ctx);
  };

  CreatePost = async (ctx) => {
    ctx.body = await this.service.CreatePost(ctx);
  };

  UpdatePost = async (ctx) => {
    ctx.body = await this.service.UpdatePost(ctx);
  };

  DeletePost = async (ctx) => {
    ctx.body = await this.service.DeletePost(ctx);
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
