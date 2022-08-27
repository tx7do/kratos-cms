/**
 * @description: Login interfaces parameters
 */
export interface LoginParams {
  userName: string;
  password: string;
}

export interface RoleInfo {
  roleName: string;
  value: string;
}

/**
 * @description: Login interfaces return value
 */
export interface LoginResultModel {
  userId: string | number;
  token: string;
  role: RoleInfo;
}

/**
 * @description: Get user information return value
 */
export interface GetUserInfoModel {
  roles: RoleInfo[];
  id: string | number;
  userName: string;
  nickName: string;
  avatar: string;
  email: string;
  description?: string;
}
