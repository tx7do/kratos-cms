export enum SortOrder {
  ASCENDING = 'ascend',
  DESCENDING = 'descend',
}

export interface PagingRequest {
  page?: number; // 页码
  pageSize?: number; // 每页数量
  query?: { [key: string]: string }; // 查询条件
  orderBy?: { [key: string]: string }; // 排序
  nopaging?: boolean; // 不分页
}
