import { BasicColumn, FormSchema } from '/@/components/Table';

export const columns: BasicColumn[] = [
  {
    title: '标题',
    dataIndex: 'title',
    sorter: false,
  },
  {
    title: '状态',
    dataIndex: 'status',
    sorter: false,
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
  },
  {
    title: '访问',
    dataIndex: 'visits',
    sorter: false,
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
    required: true,
  },
];
