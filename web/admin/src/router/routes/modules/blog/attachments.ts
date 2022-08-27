import type { AppRouteModule } from '/@/router/types';

import { t } from '/@/hooks/web/useI18n';

const attachments: AppRouteModule = {
  path: '/attachments',
  name: 'Attachments',
  component: () => import('/@/views/blog/attachments/index.vue'),
  meta: {
    orderNo: 4,
    icon: 'ant-design:cloud-upload-outlined',
    title: t('routes.blog.attachments.moduleName'),
  },
};

export default attachments;
