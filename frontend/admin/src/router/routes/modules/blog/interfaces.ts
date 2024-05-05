import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const interfaces: AppRouteModule = {
  path: '/interfaces',
  name: 'Interface',
  component: LAYOUT,
  redirect: '/interfaces/themes',
  meta: {
    orderNo: 6,
    icon: 'ant-design:skin-outlined',
    title: t('routes.blog.interfaces.moduleName'),
  },

  children: [
    {
      path: 'themes',
      name: 'ThemeList',
      meta: {
        title: t('routes.blog.interfaces.themeList'),
        icon: 'ant-design:skin-outlined',
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/blog/interfaces/themes/list/index.vue'),
    },

    {
      path: 'themes/setting/:id',
      name: 'ThemeSetting',
      meta: {
        title: t('routes.blog.interfaces.themeSetting'),
        icon: 'ant-design:control-outlined',
        currentActiveMenu: '/themes/list',
        ignoreKeepAlive: false,
        hideMenu: true,
        showMenu: false,
      },
      component: () => import('/@/views/blog/interfaces/themes/setting/index.vue'),
    },

    {
      path: 'themes/edit/:id',
      name: 'ThemeEdit',
      meta: {
        title: t('routes.blog.interfaces.themeEdit'),
        icon: 'ant-design:edit-outlined',
        currentActiveMenu: '/themes/list',
        ignoreKeepAlive: false,
        hideMenu: true,
        showMenu: false,
      },
      component: () => import('/@/views/blog/interfaces/themes/edit/index.vue'),
    },

    {
      path: 'menus',
      name: 'MenuList',
      meta: {
        title: t('routes.blog.interfaces.menuList'),
        icon: 'ant-design:menu-outlined',
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/blog/interfaces/menus/index.vue'),
    },
  ],
};

export default interfaces;
