import TagService from '../service/TagService';

class TagController {
  private service: TagService = new TagService();
}

export const tagController = new TagController();

export const Routes = [];
