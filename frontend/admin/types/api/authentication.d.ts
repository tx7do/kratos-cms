interface Login {
  (params: LoginRequest): Promise<LoginResponse>;
}

interface Logout {
  (params: LogoutRequest): Promise<{}>;
}

interface Register {
  (params: RegisterRequest): Promise<RegisterResponse>;
}

interface GetMe {
  (): Promise<User>;
}

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

interface LoginRequest {
  username?: string;
  password?: string;
}

interface LoginResponse {
  id?: number; //用户ID
  username?: string; //用户名
  token_type?: string; //令牌类型
  access_token?: string; //访问令牌
  refresh_token?: string; //刷新令牌
}

interface LogoutRequest {
  id?: number;
}

interface RegisterRequest {
  user?: User;
}

interface RegisterResponse {
  id?: number;
  userName?: string;
  token?: string;
}
