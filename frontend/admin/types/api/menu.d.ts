interface ListMenu {
  (params: PagingRequest): Promise<ListMenuResponse>;
}

interface GetMenu {
  (params: GetMenuRequest): Promise<Menu>;
}

interface CreateMenu {
  (params: CreateMenuRequest): Promise<Menu>;
}

interface UpdateMenu {
  (params: UpdateMenuRequest): Promise<Menu>;
}

interface DeleteMenu {
  (params: DeleteMenuRequest): Promise<{}>;
}

interface Menu {
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

interface ListMenuResponse {
  items?: Menu[];
  total?: number;
}

interface GetMenuRequest {
  id?: number;
}

interface CreateMenuRequest {
  menu?: Menu;
  operatorId?: number;
}

interface UpdateMenuRequest {
  id?: number;
  menu?: Menu;
  operatorId?: number;
}

interface DeleteMenuRequest {
  id?: number;
  operatorId?: number;
}
