import { defHttp } from '/@/utils/http/axios';
import { LoginParams, LoginResultModel, GetUserInfoModel } from './model/userModel';

import { ErrorMessageMode } from '/#/axios';

enum Api {
  Login = '/login',
  Logout = '/logout',
  GetUserInfo = '/me',
}

/**
 * @description: user login api
 */
export function loginApi(params: LoginParams, mode: ErrorMessageMode = 'modal') {
  return defHttp.post<LoginResultModel>(
    {
      url: Api.Login,
      params,
    },
    {
      errorMessageMode: mode,
    },
  );
}

/**
 * @description: getUserInfo
 */
export function getUserInfo() {
  return defHttp.get<GetUserInfoModel>({ url: Api.GetUserInfo }, { errorMessageMode: 'none' });
}

export function doLogout() {
  return defHttp.get({ url: Api.Logout });
}
