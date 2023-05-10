import { defHttp } from '/@/utils/http/axios';
import * as protocol from '/&/authentication';

// 用户登陆
export const Logon: protocol.Login = (params) => {
  return defHttp.post<protocol.LoginResponse>(
    { url: protocol.route.Login, params },
    {
      errorMessageMode: 'modal',
    },
  );
};

// 用户登出
export const Logout: protocol.Logout = (params) => {
  return defHttp.post<{}>(
    { url: protocol.route.Logout, params },
    {
      errorMessageMode: 'none',
    },
  );
};

// 获取用户信息
export const GetMe: protocol.GetMe = () => {
  return defHttp.get<protocol.User>(
    { url: protocol.route.GetMe },
    {
      errorMessageMode: 'none',
    },
  );
};
