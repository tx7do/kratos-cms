import { Result } from '../utils/utils';
import { createTags, createTag } from '../utils/mock';
import * as pagination from '/&/pagination';

export default class TagService {
  ListTag = async (ctx) => {
    const request: pagination.PagingRequest = ctx.query;
    ctx.body = Result.pageSuccess(request.page, request.pageSize, createTags(request.pageSize));
  };

  GetTag = async (ctx) => {
    ctx.body = Result.success(createTag());
  };

  CreateTag = async (ctx) => {
    ctx.body = Result.success(createTag());
  };

  UpdateTag = async (ctx) => {
    ctx.body = Result.success(createTag());
  };

  DeleteTag = async (ctx) => {
    ctx.body = Result.success({});
  };
}
