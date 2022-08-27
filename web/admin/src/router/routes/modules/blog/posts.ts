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
        ignoreKeepAlive: false,
        icon: 'ant-design:read-outlined',
      },
      component: () => import('/@/views/blog/posts/list/index.vue'),
    },

    {
      path: 'write',
      name: 'PostWrite',
      meta: {
        title: t('routes.blog.posts.postWrite'),
        ignoreKeepAlive: false,
        icon: 'ant-design:highlight-outlined',
      },
      component: () => import('/@/views/blog/posts/write/index.vue'),
    },

    {
      path: 'edit',
      name: 'PostEdit',
      meta: {
        title: t('routes.blog.posts.postEdit'),
        ignoreKeepAlive: false,
        icon: 'ant-design:edit-outlined',
      },
      component: () => import('/@/views/blog/posts/edit/index.vue'),
    },

    {
      path: 'categories',
      name: 'CategoryList',
      meta: {
        title: t('routes.blog.posts.categoryList'),
        ignoreKeepAlive: false,
        icon: 'ant-design:unordered-list-outlined',
      },
      component: () => import('/@/views/blog/posts/categories/index.vue'),
    },

    {
      path: 'tags',
      name: 'TagList',
      meta: {
        title: t('routes.blog.posts.tagList'),
        ignoreKeepAlive: false,
        icon: 'ant-design:tags-outlined',
      },
      component: () => import('/@/views/blog/posts/tags/index.vue'),
    },
  ],
};

export default posts;
