import { Random } from 'mockjs';
import { Result } from '../utils/utils';
import * as protocol from '../../../../api/post';

export default class PostService {
  ListPost = async (ctx) => {
    const postList: any[] = [];
    for (let index = 0; index < 20; index++) {
      postList.push({
        id: `${index}`,
        title: Random.ctitle(),
        visits: Random.integer(0, 101000),
        commentCount: Random.integer(0, 10100),
        createTime: Random.datetime(),
        editorType: 'MARKDOWN',
        status: 'PUBLISHED',
      });
    }
    ctx.body = Result.success(postList);
  };

  GetPost = async (ctx) => {};

  CreatePost = async (ctx) => {};

  UpdatePost = async (ctx) => {};

  DeletePost = async (ctx) => {};
}
