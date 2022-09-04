import { Random } from 'mockjs';
import { Result } from '../utils/utils';

export default class PhotoService {
  ListPhoto = async (ctx) => {
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

  GetPhoto = async (ctx) => {};

  CreatePhoto = async (ctx) => {};

  UpdatePhoto = async (ctx) => {};

  DeletePhoto = async (ctx) => {};
}
