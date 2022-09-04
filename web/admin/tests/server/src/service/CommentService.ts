import { Random } from 'mockjs';
import { Result } from '../utils/utils';

export default class CommentService {
  ListComment = async (ctx) => {
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

  GetComment = async (ctx) => {};

  CreateComment = async (ctx) => {};

  UpdateComment = async (ctx) => {};

  DeleteComment = async (ctx) => {};
}
