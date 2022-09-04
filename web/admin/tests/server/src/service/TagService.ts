import { Random } from 'mockjs';
import { Result } from '../utils/utils';

export default class TagService {
  ListTag = async (ctx) => {
    const tagList: any[] = [];
    for (let index = 0; index < 20; index++) {
      tagList.push({
        id: `${index}`,
        title: Random.name(),
        createTime: Random.datetime(),
      });
    }
    ctx.body = Result.success(tagList);
  };

  GetTag = async (ctx) => {};

  CreateTag = async (ctx) => {};

  UpdateTag = async (ctx) => {};

  DeleteTag = async (ctx) => {};
}
