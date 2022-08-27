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
        ignoreKeepAlive: false,
        icon: 'ant-design:book-outlined',
      },
      component: () => import('/@/views/blog/sheets/list/index.vue'),
    },

    {
      path: 'write',
      name: 'SheetWrite',
      meta: {
        title: t('routes.blog.sheets.sheetWrite'),
        ignoreKeepAlive: false,
        icon: 'ant-design:appstore-add-outlined',
      },
      component: () => import('/@/views/blog/sheets/write/index.vue'),
    },

    {
      path: 'edit',
      name: 'SheetEdit',
      meta: {
        title: t('routes.blog.sheets.sheetEdit'),
        ignoreKeepAlive: false,
        icon: 'ant-design:edit-outlined',
      },
      component: () => import('/@/views/blog/sheets/edit/index.vue'),
    },

    {
      path: 'links',
      name: 'LinkList',
      meta: {
        title: t('routes.blog.sheets.linkList'),
        ignoreKeepAlive: false,
        icon: 'ant-design:link-outlined',
      },
      component: () => import('/@/views/blog/sheets/links/index.vue'),
    },

    {
      path: 'photos',
      name: 'PhotoList',
      meta: {
        title: t('routes.blog.sheets.photoList'),
        ignoreKeepAlive: false,
        icon: 'ant-design:picture-outlined',
      },
      component: () => import('/@/views/blog/sheets/photos/index.vue'),
    },

    {
      path: 'journals',
      name: 'JournalList',
      meta: {
        title: t('routes.blog.sheets.journalList'),
        ignoreKeepAlive: false,
        icon: 'ant-design:history-outlined',
      },
      component: () => import('/@/views/blog/sheets/journals/index.vue'),
    },
  ],
};

export default sheets;
