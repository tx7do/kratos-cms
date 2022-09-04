import { Result } from '../utils/utils';
import { createCategories, createCategory } from '../utils/mock';

export default class CategoryService {
  ListCategory = async (ctx) => {
    const total = 20;
    ctx.body = Result.success({
      items: createCategories(total),
      total: total,
    });
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
