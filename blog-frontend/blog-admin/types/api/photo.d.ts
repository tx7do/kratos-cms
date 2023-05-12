interface ListPhoto {
  (params: PagingRequest): Promise<ListPhotoResponse>;
}

interface GetPhoto {
  (params: GetPhotoRequest): Promise<Photo>;
}

interface CreatePhoto {
  (params: CreatePhotoRequest): Promise<Photo>;
}

interface UpdatePhoto {
  (params: UpdatePhotoRequest): Promise<Photo>;
}

interface DeletePhoto {
  (params: DeletePhotoRequest): Promise<{}>;
}

interface Photo {
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

interface ListPhotoResponse {
  items?: Photo[];
  total?: number;
}

interface GetPhotoRequest {
  id?: number;
}

interface CreatePhotoRequest {
  photo?: Photo;
  operatorId?: number;
}

interface UpdatePhotoRequest {
  id?: number;
  photo?: Photo;
  operatorId?: number;
}

interface DeletePhotoRequest {
  id?: number;
  operatorId?: number;
}

