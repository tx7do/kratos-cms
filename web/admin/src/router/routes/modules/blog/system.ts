import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const system: AppRouteModule = {
  path: '/system',
  name: 'System',
  component: LAYOUT,
  redirect: '/system/options',
  meta: {
    orderNo: 2000,
    icon: 'ant-design:setting-outline',
    title: t('routes.blog.system.moduleName'),
  },
  children: [
    {
      path: 'developer/options',
      name: 'DeveloperOptions',
      meta: {
        title: t('routes.blog.system.developerOptions'),
        ignoreKeepAlive: false,
        icon: 'ant-design:experiment-outlined',
      },
      component: () => import('/@/views/blog/system/developer/index.vue'),
    },

    {
      path: 'options',
      name: 'SystemOptions',
      meta: {
        title: t('routes.blog.system.systemOptions'),
        ignoreKeepAlive: false,
        icon: 'ant-design:sliders-outlined',
      },
      component: () => import('/@/views/blog/system/options/index.vue'),
    },

    {
      path: 'tools',
      name: 'ToolList',
      meta: {
        title: t('routes.blog.system.toolList'),
        ignoreKeepAlive: false,
        icon: 'ant-design:tool-outlined',
      },
      component: () => import('/@/views/blog/system/tools/index.vue'),
    },

    {
      path: 'actionlogs',
      name: 'SystemActionLogs',
      meta: {
        title: t('routes.blog.system.actionLogs'),
        ignoreKeepAlive: false,
        icon: 'ant-design:database-outlined',
      },
      component: () => import('/@/views/blog/system/actionlogs/index.vue'),
    },

    {
      path: 'about',
      name: 'AboutPage',
      component: () => import('/@/views/blog/system/about/index.vue'),
      meta: {
        title: t('routes.dashboard.about'),
        icon: 'ant-design:question-circle-outlined',
        hideMenu: false,
      },
    },
  ],
};

export default system;
