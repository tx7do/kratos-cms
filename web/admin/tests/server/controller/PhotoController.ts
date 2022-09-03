import PostService from '../service/PostService';

class PhotoController {
  private service: PostService = new PostService();
}

export const photoController = new PhotoController();

export const Routes = [];
