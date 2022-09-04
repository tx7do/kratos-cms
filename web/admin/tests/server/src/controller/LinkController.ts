import LinkService from '../service/LinkService';
import * as protocol from '/&/link';

class LinkController {
  private service: LinkService = new LinkService();

  ListLink = async (ctx) => {
    await this.service.ListLink(ctx);
  };

  GetLink = async (ctx) => {
    await this.service.GetLink(ctx);
  };

  CreateLink = async (ctx) => {
    await this.service.CreateLink(ctx);
  };

  UpdateLink = async (ctx) => {
    await this.service.UpdateLink(ctx);
  };

  DeleteLink = async (ctx) => {
    await this.service.DeleteLink(ctx);
  };
}

export const linkController = new LinkController();

export const Routes = [
  {
    path: protocol.route.Links,
    method: 'get',
    action: linkController.ListLink,
  },
  {
    path: protocol.route.Links,
    method: 'get',
    action: linkController.GetLink,
  },
  {
    path: protocol.route.Links,
    method: 'post',
    action: linkController.CreateLink,
  },
  {
    path: protocol.route.Links,
    method: 'put',
    action: linkController.UpdateLink,
  },
  {
    path: protocol.route.Links,
    method: 'delete',
    action: linkController.DeleteLink,
  },
];
