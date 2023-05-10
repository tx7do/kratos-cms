export const route = {};

export interface User {
  id?: number;
  userName?: string;
  nickName?: string;
  email?: string;
  avatar?: string;
  description?: string;
  password?: string;
  createTime?: string;
  updateTime?: string;
}

export interface CreateUserRequest {
  user?: User;
  operatorId?: number;
}

export interface UpdateUserRequest {
  id?: number;
  user?: User;
  operatorId?: number;
}

export interface DeleteUserRequest {
  id?: number;
  operatorId?: number;
}
