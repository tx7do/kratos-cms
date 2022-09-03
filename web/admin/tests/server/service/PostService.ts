import { Random } from 'mockjs';
import { Result } from '../utils';
import * as protocol from '../../../api/post';


export default class PostService {
  ListPost = async (ctx) => {
    const postList: any[] = [];
    for (let index = 0; index < 20; index++) {
      postList.push({
        id: `${index}`,
        title: Random.name(),
        createTime: Random.datetime(),
      });
    }
    return postList;
    return Result.success(postList);
  };

  GetPost = async (ctx) => {

  };

  CreatePost = async (ctx) => {

  };

  UpdatePost = async (ctx) => {

  };

  DeletePost = async (ctx) => {

  };
}
