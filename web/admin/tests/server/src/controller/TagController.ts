import TagService from '../service/TagService';
import * as protocol from '/&/tag';

class TagController {
  private service: TagService = new TagService();

  ListTag = async (ctx) => {
    await this.service.ListTag(ctx);
  };

  GetTag = async (ctx) => {
    await this.service.GetTag(ctx);
  };

  CreateTag = async (ctx) => {
    await this.service.CreateTag(ctx);
  };

  UpdateTag = async (ctx) => {
    await this.service.UpdateTag(ctx);
  };

  DeleteTag = async (ctx) => {
    await this.service.DeleteTag(ctx);
  };
}

export const tagController = new TagController();

export const Routes = [
  {
    path: protocol.route.Tags,
    method: 'get',
    action: tagController.ListTag,
  },
  {
    path: protocol.route.Tags,
    method: 'get',
    action: tagController.GetTag,
  },
  {
    path: protocol.route.Tags,
    method: 'post',
    action: tagController.CreateTag,
  },
  {
    path: protocol.route.Tags,
    method: 'put',
    action: tagController.UpdateTag,
  },
  {
    path: protocol.route.Tags,
    method: 'delete',
    action: tagController.DeleteTag,
  },
];
