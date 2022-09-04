import PhotoService from '../service/PhotoService';
import * as protocol from '../../../../api/photo';

class PhotoController {
  private service: PhotoService = new PhotoService();

  ListPhoto = async (ctx) => {
    await this.service.ListPhoto(ctx);
  };

  GetPhoto = async (ctx) => {
    await this.service.GetPhoto(ctx);
  };

  CreatePhoto = async (ctx) => {
    await this.service.CreatePhoto(ctx);
  };

  UpdatePhoto = async (ctx) => {
    await this.service.UpdatePhoto(ctx);
  };

  DeletePhoto = async (ctx) => {
    await this.service.DeletePhoto(ctx);
  };
}

export const photoController = new PhotoController();

export const Routes = [
  {
    path: protocol.route.Photos,
    method: 'get',
    action: photoController.ListPhoto,
  },
  {
    path: protocol.route.Photos,
    method: 'get',
    action: photoController.GetPhoto,
  },
  {
    path: protocol.route.Photos,
    method: 'post',
    action: photoController.CreatePhoto,
  },
  {
    path: protocol.route.Photos,
    method: 'put',
    action: photoController.UpdatePhoto,
  },
  {
    path: protocol.route.Photos,
    method: 'delete',
    action: photoController.DeletePhoto,
  },
];
