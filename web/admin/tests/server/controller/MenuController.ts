import MenuService from '../service/MenuService';

class MenuController {
  private service: MenuService = new MenuService();
}

export const menuController = new MenuController();

export const Routes = [];
