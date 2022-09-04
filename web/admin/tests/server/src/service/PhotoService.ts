import { Result } from '../utils/utils';
import * as pagination from '/&/pagination';
import { createPhotos, createPhoto } from '../utils/mock';

export default class PhotoService {
  ListPhoto = async (ctx) => {
    const request: pagination.PagingRequest = ctx.query;
    ctx.body = Result.pageSuccess(request.page, request.pageSize, createPhotos(request.pageSize));
  };

  GetPhoto = async (ctx) => {
    ctx.body = Result.success(createPhoto());
  };

  CreatePhoto = async (ctx) => {
    ctx.body = Result.success(createPhoto());
  };

  UpdatePhoto = async (ctx) => {
    ctx.body = Result.success(createPhoto());
  };

  DeletePhoto = async (ctx) => {
    ctx.body = Result.success({});
  };
}
