declare type Recordable<T = any> = Record<string, T>;

export class Result {
  static success<T = Recordable>(result: T) {
    return result;
  }

  static error(message = 'Request failed', { code = -1 } = {}) {
    return {
      code,
      reason: '',
      message,
      metadata: {},
    };
  }

  static pagination<T = any>(pageNo: number, pageSize: number, array: T[]): T[] {
    const offset = (pageNo - 1) * Number(pageSize);
    return offset + Number(pageSize) >= array.length
      ? array.slice(offset, array.length)
      : array.slice(offset, offset + Number(pageSize));
  }

  static pageSuccess<T = any>(page: number, per_page: number, list: T[], { message = 'ok' } = {}) {
    const pageData = Result.pagination(page, per_page, list);
    return {
      ...Result.success({
        items: pageData,
        total: list.length,
      }),
      message,
    };
  }
}

export interface requestParams {
  method: string;
  body: any;
  headers?: { authorization?: string };
  query: any;
}
