import { Result } from '../utils/utils';
import { createPosts, createPost } from '../utils/mock';

export default class PostService {
  ListPost = async (ctx) => {
    ctx.body = Result.success(createPosts(20));
  };

  GetPost = async (ctx) => {
    ctx.body = Result.success(createPost());
  };

  CreatePost = async (ctx) => {
    ctx.body = Result.success(createPost());
  };

  UpdatePost = async (ctx) => {
    ctx.body = Result.success(createPost());
  };

  DeletePost = async (ctx) => {
    ctx.body = Result.success({});
  };
}
