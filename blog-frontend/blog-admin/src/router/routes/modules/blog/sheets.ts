import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const sheets: AppRouteModule = {
  path: '/sheets',
  name: 'Sheets',
  component: LAYOUT,
  redirect: '/sheets/list',
  meta: {
    orderNo: 3,
    icon: 'ant-design:book-outlined',
    title: t('routes.blog.sheets.moduleName'),
  },

  children: [
    {
      path: 'list',
      name: 'SheetList',
      meta: {
        title: t('routes.blog.sheets.sheetList'),
        icon: 'ant-design:book-outlined',
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/blog/sheets/list/index.vue'),
    },

    {
      path: 'write',
      name: 'SheetWrite',
      meta: {
        title: t('routes.blog.sheets.sheetWrite'),
        icon: 'ant-design:appstore-add-outlined',
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/blog/sheets/edit/index.vue'),
    },

    {
      path: 'edit/:id',
      name: 'SheetEdit',
      meta: {
        title: t('routes.blog.sheets.sheetEdit'),
        icon: 'ant-design:edit-outlined',
        currentActiveMenu: '/sheets/list',
        ignoreKeepAlive: true,
        hideMenu: true,
        showMenu: false,
      },
      component: () => import('/@/views/blog/sheets/edit/index.vue'),
    },

    {
      path: 'links',
      name: 'LinkList',
      meta: {
        title: t('routes.blog.sheets.linkList'),
        icon: 'ant-design:link-outlined',
        ignoreKeepAlive: false,
        hideMenu: true,
        showMenu: false,
      },
      component: () => import('/@/views/blog/sheets/links/index.vue'),
    },

    {
      path: 'photos',
      name: 'PhotoList',
      meta: {
        title: t('routes.blog.sheets.photoList'),
        icon: 'ant-design:picture-outlined',
        ignoreKeepAlive: false,
        hideMenu: true,
        showMenu: false,
      },
      component: () => import('/@/views/blog/sheets/photos/index.vue'),
    },

    {
      path: 'journals',
      name: 'JournalList',
      meta: {
        title: t('routes.blog.sheets.journalList'),
        icon: 'ant-design:history-outlined',
        ignoreKeepAlive: false,
        hideMenu: true,
        showMenu: false,
      },
      component: () => import('/@/views/blog/sheets/journals/index.vue'),
    },
  ],
};

export default sheets;
