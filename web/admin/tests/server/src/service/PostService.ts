import { Result } from '../utils/utils';
import { createPosts, createPost } from '../utils/mock';
import * as pagination from '/&/pagination';

export default class PostService {
  ListPost = async (ctx) => {
    const request: pagination.PagingRequest = ctx.query;
    ctx.body = Result.pageSuccess(request.page, request.pageSize, createPosts(request.pageSize));
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
