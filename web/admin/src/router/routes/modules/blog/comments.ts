import type { AppRouteModule } from '/@/router/types';

import {t} from '/@/hooks/web/useI18n';

const comments: AppRouteModule = {
  path: '/comments',
  name: 'Comments',
  component: () => import('/@/views/blog/comments/index.vue'),
  meta: {
    orderNo: 5,
    icon: 'ant-design:comment-outlined',
    title: t('routes.blog.comments.moduleName'),
  },
};

export default comments;
