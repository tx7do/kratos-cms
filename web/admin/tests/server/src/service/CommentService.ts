import { Result } from '../utils/utils';
import * as pagination from '/&/pagination';
import { createComments, createComment } from '../utils/mock';

export default class CommentService {
  ListComment = async (ctx) => {
    const request: pagination.PagingRequest = ctx.query;
    ctx.body = Result.pageSuccess(request.page, request.pageSize, createComments(request.pageSize));
  };

  GetComment = async (ctx) => {
    ctx.body = Result.success(createComment());
  };

  CreateComment = async (ctx) => {
    ctx.body = Result.success(createComment());
  };

  UpdateComment = async (ctx) => {
    ctx.body = Result.success(createComment());
  };

  DeleteComment = async (ctx) => {
    ctx.body = Result.success({});
  };
}
