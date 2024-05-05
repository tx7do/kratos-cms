interface User {
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

interface CreateUserRequest {
  user?: User;
  operatorId?: number;
}

interface UpdateUserRequest {
  id?: number;
  user?: User;
  operatorId?: number;
}

interface DeleteUserRequest {
  id?: number;
  operatorId?: number;
}
