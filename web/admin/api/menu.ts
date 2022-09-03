import * as pagination from './pagination';

export const route = {
  Menus: '/menus',
  MenuWithId: (id) => `/menus/${id}`,
};

export interface ListMenu {
  (params: pagination.PagingRequest): Promise<ListMenuResponse>;
}

export interface GetMenu {
  (params: GetMenuRequest): Promise<Menu>;
}

export interface CreateMenu {
  (params: CreateMenuRequest): Promise<Menu>;
}

export interface UpdateMenu {
  (params: UpdateMenuRequest): Promise<Menu>;
}

export interface DeleteMenu {
  (params: DeleteMenuRequest): Promise<{}>;
}

export interface Menu {
  id?: number;
  name?: string;
  url?: string;
  priority?: number;
  target?: string;
  icon?: string;
  parentId?: number;
  team?: string;
  createTime?: string;
  updateTime?: string;
}

export interface ListMenuResponse {
  items?: Menu[];
  total?: number;
}

export interface GetMenuRequest {
  id?: number;
}

export interface CreateMenuRequest {
  menu?: Menu;
  operatorId?: number;
}

export interface UpdateMenuRequest {
  id?: number;
  menu?: Menu;
  operatorId?: number;
}

export interface DeleteMenuRequest {
  id?: number;
  operatorId?: number;
}
