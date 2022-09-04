import { BasicColumn, FormSchema } from '/@/components/Table';
import { h } from 'vue';
import { Badge } from 'ant-design-vue';
import { Post } from '/&/post';

export const columns: BasicColumn[] = [
  {
    title: '标题',
    dataIndex: 'title',
    sorter: false,
  },
  {
    title: '状态',
    align: 'center',
    dataIndex: 'status',
    sorter: false,
    width: 120,
    customRender: ({ record }) => {
      const rd = record as Post;
      return h(Badge, {
        status: 'success',
        text: '已发布',
      });
    },
  },
  {
    title: '分类',
    dataIndex: 'categoryId',
    sorter: false,
  },
  {
    title: '标签',
    dataIndex: 'tags',
    sorter: false,
  },
  {
    title: '评论',
    dataIndex: 'commentCount',
    sorter: false,
    width: 60,
    customRender: ({ record }) => {
      const rd = record as Post;
      return h(Badge, {
        count: rd.commentCount,
        numberStyle: { backgroundColor: 'volcano' },
        overflowCount: 10000,
      });
    },
  },
  {
    title: '访问',
    dataIndex: 'visits',
    sorter: false,
    width: 60,
    customRender: ({ record }) => {
      const rd = record as Post;
      return h(Badge, {
        count: rd.visits,
        numberStyle: { backgroundColor: '#2db7f5' },
        overflowCount: 100000,
      });
    },
  },
  {
    title: '发布时间',
    dataIndex: 'createTime',
    sorter: false,
  },
];

export const searchFormSchema: FormSchema[] = [
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
  },
];
