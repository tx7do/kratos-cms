import { Random } from 'mockjs';
import { Result } from '../utils/utils';

export default class CategoryService {
  ListCategory = async (ctx) => {
    const photoList: any[] = [];
    for (let index = 0; index < 20; index++) {
      photoList.push({
        id: `${index}`,
        title: Random.name(),
        createTime: Random.datetime(),
      });
    }
    ctx.body = Result.success(photoList);
  };

  GetCategory = async (ctx) => {};

  CreateCategory = async (ctx) => {};

  UpdateCategory = async (ctx) => {};

  DeleteCategory = async (ctx) => {};
}
