import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const posts: AppRouteModule = {
  path: '/posts',
  name: 'Posts',
  component: LAYOUT,
  redirect: '/posts/list',
  meta: {
    orderNo: 2,
    icon: 'ant-design:read-outlined',
    title: t('routes.blog.posts.moduleName'),
  },

  children: [
    {
      path: 'list',
      name: 'PostList',
      meta: {
        title: t('routes.blog.posts.postList'),
        icon: 'ant-design:read-outlined',
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/blog/posts/list/index.vue'),
    },

    {
      path: 'write',
      name: 'PostWrite',
      meta: {
        title: t('routes.blog.posts.postWrite'),
        icon: 'ant-design:highlight-outlined',
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/blog/posts/edit/index.vue'),
    },

    {
      path: 'edit/:id',
      name: 'PostEdit',
      meta: {
        title: t('routes.blog.posts.postEdit'),
        icon: 'ant-design:edit-outlined',
        currentActiveMenu: '/posts/list',
        ignoreKeepAlive: true,
        hideMenu: true,
        showMenu: false,
      },
      component: () => import('/@/views/blog/posts/edit/index.vue'),
    },

    {
      path: 'categories',
      name: 'CategoryList',
      meta: {
        title: t('routes.blog.posts.categoryList'),
        icon: 'ant-design:unordered-list-outlined',
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/blog/posts/categories/index.vue'),
    },

    {
      path: 'tags',
      name: 'TagList',
      meta: {
        title: t('routes.blog.posts.tagList'),
        icon: 'ant-design:tags-outlined',
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/blog/posts/tags/index.vue'),
    },
  ],
};

export default posts;
