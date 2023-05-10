export const route = {
  Login: '/login',
  Logout: '/logout',
  Register: '/register',
  GetMe: '/me',
};

export interface Login {
  (params: LoginRequest): Promise<LoginResponse>;
}

export interface Logout {
  (params: LogoutRequest): Promise<{}>;
}

export interface Register {
  (params: RegisterRequest): Promise<RegisterResponse>;
}

export interface GetMe {
  (): Promise<User>;
}

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

export interface LoginRequest {
  userName?: string;
  password?: string;
}

export interface LoginResponse {
  id?: number;
  userName?: string;
  token?: string;
  refreshToken?: string;
}

export interface LogoutRequest {
  id?: number;
}

export interface RegisterRequest {
  user?: User;
}

export interface RegisterResponse {
  id?: number;
  userName?: string;
  token?: string;
}
