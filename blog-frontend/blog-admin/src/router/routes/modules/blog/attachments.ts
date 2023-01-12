import type { AppRouteModule } from '/@/router/types';

import { t } from '/@/hooks/web/useI18n';
import { LAYOUT } from '/@/router/constant';

const attachments: AppRouteModule = {
  path: '/attachments',
  name: 'Attachments',
  component: LAYOUT,
  redirect: '/attachments/index',
  meta: {
    hideChildrenInMenu: true,
    orderNo: 4,
    icon: 'ant-design:cloud-upload-outlined',
    title: t('routes.blog.attachments.moduleName'),
  },
  children: [
    {
      path: 'index',
      name: 'AttachmentsPage',
      component: () => import('/@/views/blog/attachments/index.vue'),
      meta: {
        icon: 'ant-design:cloud-upload-outlined',
        title: t('routes.blog.attachments.moduleName'),
        hideMenu: true,
      },
    },
  ],
};

export default attachments;
