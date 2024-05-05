import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const comments: AppRouteModule = {
  path: '/comments',
  name: 'Comments',
  component: LAYOUT,
  redirect: '/comments/index',
  meta: {
    orderNo: 5,
    hideChildrenInMenu: true,
    icon: 'ant-design:comment-outlined',
    title: t('routes.blog.comments.moduleName'),
  },
  children: [
    {
      path: 'index',
      name: 'CommentsPage',
      component: () => import('/@/views/blog/comments/index.vue'),
      meta: {
        icon: 'ant-design:comment-outlined',
        title: t('routes.blog.comments.moduleName'),
        hideMenu: true,
      },
    },
  ],
};

export default comments;
