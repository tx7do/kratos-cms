import UserService from '../service/UserService';
import * as protocol from '../../../api/user';

class UserController {
  private service: UserService = new UserService();

  Login = async (ctx) => {
    ctx.body = await this.service.Login();
  };

  Logout = async (ctx) => {
    ctx.body = await this.service.Logout();
  };

  GetMe = async (ctx) => {
    ctx.body = await this.service.GetMe();
  };
}

export const userController = new UserController();

export const Routes = [
  {
    path: protocol.route.Login,
    method: 'post',
    action: userController.Login,
  },
  {
    path: protocol.route.Logout,
    method: 'post',
    action: userController.Logout,
  },
  {
    path: protocol.route.GetMe,
    method: 'get',
    action: userController.GetMe,
  },
];
