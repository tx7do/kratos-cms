import * as pagination from './pagination';

export const route = {
  Photos: '/photos',
  PhotoWithId: (id) => `/photos/${id}`,
};

export interface ListPhoto {
  (params: pagination.PagingRequest): Promise<ListPhotoResponse>;
}

export interface GetPhoto {
  (params: GetPhotoRequest): Promise<Photo>;
}

export interface CreatePhoto {
  (params: CreatePhotoRequest): Promise<Photo>;
}

export interface UpdatePhoto {
  (params: UpdatePhotoRequest): Promise<Photo>;
}

export interface DeletePhoto {
  (params: DeletePhotoRequest): Promise<{}>;
}

export interface Photo {
  id?: number;
  name?: string;
  thumbnail?: string;
  url?: string;
  team?: string;
  location?: string;
  description?: string;
  likes?: number;
  takeTime?: string;
  createTime?: string;
}

export interface ListPhotoResponse {
  items?: Photo[];
  total?: number;
}

export interface GetPhotoRequest {
  id?: number;
}

export interface CreatePhotoRequest {
  photo?: Photo;
  operatorId?: number;
}

export interface UpdatePhotoRequest {
  id?: number;
  photo?: Photo;
  operatorId?: number;
}

export interface DeletePhotoRequest {
  id?: number;
  operatorId?: number;
}

