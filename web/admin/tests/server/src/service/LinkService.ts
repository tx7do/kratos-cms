import { Random } from 'mockjs';
import { Result } from '../utils/utils';

export default class LinkService {
  ListLink = async (ctx) => {
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

  GetLink = async (ctx) => {};

  CreateLink = async (ctx) => {};

  UpdateLink = async (ctx) => {};

  DeleteLink = async (ctx) => {};
}
