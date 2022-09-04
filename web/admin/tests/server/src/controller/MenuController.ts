import MenuService from '../service/MenuService';
import * as protocol from '../../../../api/menu';

class MenuController {
  private service: MenuService = new MenuService();

  ListMenu = async (ctx) => {
    await this.service.ListMenu(ctx);
  };

  GetMenu = async (ctx) => {
    await this.service.GetMenu(ctx);
  };

  CreateMenu = async (ctx) => {
    await this.service.CreateMenu(ctx);
  };

  UpdateMenu = async (ctx) => {
    await this.service.UpdateMenu(ctx);
  };

  DeleteMenu = async (ctx) => {
    await this.service.DeleteMenu(ctx);
  };
}

export const menuController = new MenuController();

export const Routes = [
  {
    path: protocol.route.Menus,
    method: 'get',
    action: menuController.ListMenu,
  },
  {
    path: protocol.route.Menus,
    method: 'get',
    action: menuController.GetMenu,
  },
  {
    path: protocol.route.Menus,
    method: 'post',
    action: menuController.CreateMenu,
  },
  {
    path: protocol.route.Menus,
    method: 'put',
    action: menuController.UpdateMenu,
  },
  {
    path: protocol.route.Menus,
    method: 'delete',
    action: menuController.DeleteMenu,
  },
];
