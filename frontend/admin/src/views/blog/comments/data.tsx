import { BasicColumn, FormSchema } from '/@/components/Table';

import PostPage from './post-page.vue';
import SheetPage from './sheet-page.vue';
import JournalPage from './journal-page.vue';

export interface TabItem {
  key: string;
  name: string;
  component: any;
}

export const tabList: TabItem[] = [
  {
    key: '1',
    name: '文章',
    component: PostPage,
  },
  {
    key: '2',
    name: '页面',
    component: SheetPage,
  },
  {
    key: '3',
    name: '日志',
    component: JournalPage,
  },
];

export const postColumns: BasicColumn[] = [
  {
    title: '昵称',
    dataIndex: 'nickName',
    sorter: false,
  },
  {
    title: '内容',
    dataIndex: 'content',
    sorter: false,
  },
  {
    title: '状态',
    dataIndex: 'status',
    sorter: false,
  },
  {
    title: '评论文章',
    dataIndex: 'postId',
    sorter: false,
  },
  {
    title: '时间',
    dataIndex: 'createTime',
    sorter: false,
  },
];

export const sheetColumns: BasicColumn[] = [
  {
    title: '昵称',
    dataIndex: 'nickName',
    sorter: false,
  },
  {
    title: '内容',
    dataIndex: 'content',
    sorter: false,
  },
  {
    title: '状态',
    dataIndex: 'status',
    sorter: false,
  },
  {
    title: '评论页面',
    dataIndex: 'sheetId',
    sorter: false,
  },
  {
    title: '时间',
    dataIndex: 'createTime',
    sorter: false,
  },
];

export const journalColumns: BasicColumn[] = [
  {
    title: '昵称',
    dataIndex: 'nickName',
    sorter: false,
  },
  {
    title: '内容',
    dataIndex: 'content',
    sorter: false,
  },
  {
    title: '状态',
    dataIndex: 'status',
    sorter: false,
  },
  {
    title: '评论日志',
    dataIndex: 'journalId',
    sorter: false,
  },
  {
    title: '时间',
    dataIndex: 'createTime',
    sorter: false,
  },
];

export const postSearchFormSchema: FormSchema[] = [
  {
    field: 'keyword',
    label: '关键词',
    component: 'Input',
    colProps: { span: 8 },
  },
  {
    field: 'status',
    label: '评论状态',
    colProps: { span: 8 },
    component: 'Select',
    defaultValue: '0',
    componentProps: {
      options: [
        { label: '已发布', value: '0' },
        { label: '待审核', value: '1' },
        { label: '回收站', value: '2' },
      ],
    },
    required: false,
  },
];
