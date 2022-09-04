import { Result } from '../utils/utils';
import { createCategories, createCategory } from '../utils/mock';
import * as pagination from '/&/pagination';

export default class CategoryService {
  ListCategory = async (ctx) => {
    const request: pagination.PagingRequest = ctx.query;
    ctx.body = Result.pageSuccess(
      request.page,
      request.pageSize,
      createCategories(request.pageSize),
    );
  };

  GetCategory = async (ctx) => {
    ctx.body = Result.success(createCategory());
  };

  CreateCategory = async (ctx) => {
    ctx.body = Result.success(createCategory());
  };

  UpdateCategory = async (ctx) => {
    ctx.body = Result.success(createCategory());
  };

  DeleteCategory = async (ctx) => {
    ctx.body = Result.success({});
  };
}
