import { defHttp } from '/@/utils/http/axios';

const route = {
  Login: '/login',
  Logout: '/logout',
  Register: '/register',
  GetMe: '/me',
};

// 用户登陆
export const Logon: Login = (params) => {
  return defHttp.post<LoginResponse>(
    {
      url: route.Login,
      params
    },
    {
      errorMessageMode: 'modal',
    },
  );
};

// 用户登出
export const Logout: Logout = (params) => {
  return defHttp.post<{}>(
    {
      url: route.Logout,
      params,
    },
    {
      errorMessageMode: 'none',
    },
  );
};

// 获取用户信息
export const GetMe: GetMe = () => {
  return defHttp.get<User>(
    {
      url: route.GetMe,
    },
    {
      errorMessageMode: 'none',
    },
  );
};
