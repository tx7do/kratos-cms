import { Result } from '../utils/utils';
import * as pagination from '/&/pagination';
import { createMenus, createMenu } from '../utils/mock';

export default class MenuService {
  ListMenu = async (ctx) => {
    const request: pagination.PagingRequest = ctx.query;
    ctx.body = Result.pageSuccess(request.page, request.pageSize, createMenus(request.pageSize));
  };

  GetMenu = async (ctx) => {
    ctx.body = Result.success(createMenu());
  };

  CreateMenu = async (ctx) => {
    ctx.body = Result.success(createMenu());
  };

  UpdateMenu = async (ctx) => {
    ctx.body = Result.success(createMenu());
  };

  DeleteMenu = async (ctx) => {
    ctx.body = Result.success({});
  };
}
