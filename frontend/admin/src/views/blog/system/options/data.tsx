import BasePage from './base-page.vue';
import SEOPage from './seo-page.vue';
import PostPage from './post-page.vue';
import CommentPage from './comment-page.vue';
import AttachmentPage from './attachment-page.vue';
import SMTPPage from './smtp-page.vue';
import OtherPage from './other-page.vue';

export interface TabItem {
  key: string;
  name: string;
  component: any;
}

export const tabList: TabItem[] = [
  {
    key: '1',
    name: '常规设置',
    component: BasePage,
  },
  {
    key: '2',
    name: 'SEO 设置',
    component: SEOPage,
  },
  {
    key: '3',
    name: '文章设置',
    component: PostPage,
  },
  {
    key: '4',
    name: '评论设置',
    component: CommentPage,
  },
  {
    key: '5',
    name: '附件设置',
    component: AttachmentPage,
  },
  {
    key: '6',
    name: '邮件设置',
    component: SMTPPage,
  },
  {
    key: '7',
    name: '其他设置',
    component: OtherPage,
  },
];
