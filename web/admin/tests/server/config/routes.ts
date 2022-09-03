import { Routes as AttachmentRoutes } from '../controller/AttachmentController';
import { Routes as CategoryRoutes } from '../controller/CategoryController';
import { Routes as CommentRoutes } from '../controller/CommentController';
import { Routes as LinkRoutes } from '../controller/LinkController';
import { Routes as MenuRoutes } from '../controller/MenuController';
import { Routes as PhotoRoutes } from '../controller/PhotoController';
import { Routes as PostRoutes } from '../controller/PostController';
import { Routes as SystemRoutes } from '../controller/SystemController';
import { Routes as TagRoutes } from '../controller/TagController';
import { Routes as UserRoutes } from '../controller/UserController';

// export const AppRoutes = SystemRoutes;

export const AppRoutes = [
  ...AttachmentRoutes,
  ...CategoryRoutes,
  ...CommentRoutes,
  ...LinkRoutes,
  ...MenuRoutes,
  ...PhotoRoutes,
  ...PostRoutes,
  ...SystemRoutes,
  ...TagRoutes,
  ...UserRoutes,
];
