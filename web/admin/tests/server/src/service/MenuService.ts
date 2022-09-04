import { Random } from 'mockjs';
import { Result } from '../utils/utils';

export default class MenuService {
  ListMenu = async (ctx) => {
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

  GetMenu = async (ctx) => {};

  CreateMenu = async (ctx) => {};

  UpdateMenu = async (ctx) => {};

  DeleteMenu = async (ctx) => {};
}
