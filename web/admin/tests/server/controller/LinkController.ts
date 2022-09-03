import LinkService from '../service/LinkService';

class LinkController {
  private service: LinkService = new LinkService();
}

export const linkController = new LinkController();

export const Routes = [];
