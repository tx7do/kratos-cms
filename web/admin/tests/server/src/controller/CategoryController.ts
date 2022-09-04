import CategoryService from '../service/CategoryService';
import * as protocol from '/&/category';

class CategoryController {
  private service: CategoryService = new CategoryService();

  ListCategory = async (ctx) => {
    await this.service.ListCategory(ctx);
  };

  GetCategory = async (ctx) => {
    await this.service.GetCategory(ctx);
  };

  CreateCategory = async (ctx) => {
    await this.service.CreateCategory(ctx);
  };

  UpdateCategory = async (ctx) => {
    await this.service.UpdateCategory(ctx);
  };

  DeleteCategory = async (ctx) => {
    await this.service.DeleteCategory(ctx);
  };
}

export const categoryController = new CategoryController();

export const Routes = [
  {
    path: protocol.route.Categories,
    method: 'get',
    action: categoryController.ListCategory,
  },
  {
    path: protocol.route.Categories,
    method: 'get',
    action: categoryController.GetCategory,
  },
  {
    path: protocol.route.Categories,
    method: 'post',
    action: categoryController.CreateCategory,
  },
  {
    path: protocol.route.Categories,
    method: 'put',
    action: categoryController.UpdateCategory,
  },
  {
    path: protocol.route.Categories,
    method: 'delete',
    action: categoryController.DeleteCategory,
  },
];
