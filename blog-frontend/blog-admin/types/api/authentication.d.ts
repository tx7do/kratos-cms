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
  userName?: string;
  password?: string;
}

interface LoginResponse {
  id?: number;
  userName?: string;
  token?: string;
  refreshToken?: string;
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
