import CategoryService from '../service/CategoryService';

class CategoryController {
  private service: CategoryService = new CategoryService();
}

export const categoryController = new CategoryController();

export const Routes = [];
