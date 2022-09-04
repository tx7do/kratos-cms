import { Result } from '../utils/utils';
import { createTags, createTag } from '../utils/mock';

export default class TagService {
  ListTag = async (ctx) => {
    const total = 20;
    ctx.body = Result.success({
      items: createTags(total),
      total: total,
    });
  };

  GetTag = async (ctx) => {
    ctx.body = Result.success(createTag());
  };

  CreateTag = async (ctx) => {
    ctx.body = Result.success(createTag());
  };

  UpdateTag = async (ctx) => {
    ctx.body = Result.success(createTag());
  };

  DeleteTag = async (ctx) => {
    ctx.body = Result.success({});
  };
}
