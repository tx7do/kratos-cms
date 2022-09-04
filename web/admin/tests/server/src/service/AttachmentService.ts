import { Random } from 'mockjs';
import { Result } from '../utils/utils';

export default class AttachmentService {
  ListAttachment = async (ctx) => {
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

  GetAttachment = async (ctx) => {};

  CreateAttachment = async (ctx) => {};

  UpdateAttachment = async (ctx) => {};

  DeleteAttachment = async (ctx) => {};
}
