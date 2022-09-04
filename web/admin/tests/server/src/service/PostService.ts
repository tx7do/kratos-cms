import { Result } from '../utils/utils';
import { createPosts, createPost } from '../utils/mock';

export default class PostService {
  ListPost = async (ctx) => {
    const total = 20;
    ctx.body = Result.success({
      items: createPosts(total),
      total: total,
    });
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
