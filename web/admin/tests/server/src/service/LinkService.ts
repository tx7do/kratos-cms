import { Result } from '../utils/utils';
import * as pagination from '/&/pagination';
import { createLinks, createLink } from '../utils/mock';

export default class LinkService {
  ListLink = async (ctx) => {
    const request: pagination.PagingRequest = ctx.query;
    ctx.body = Result.pageSuccess(request.page, request.pageSize, createLinks(request.pageSize));
  };

  GetLink = async (ctx) => {
    ctx.body = Result.success(createLink());
  };

  CreateLink = async (ctx) => {
    ctx.body = Result.success(createLink());
  };

  UpdateLink = async (ctx) => {
    ctx.body = Result.success(createLink());
  };

  DeleteLink = async (ctx) => {
    ctx.body = Result.success({});
  };
}
