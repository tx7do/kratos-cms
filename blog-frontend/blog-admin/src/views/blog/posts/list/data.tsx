import { BasicColumn, FormSchema } from '/@/components/Table';
import { h } from 'vue';
import { Badge, Tag, Row, Col } from 'ant-design-vue';

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
      let statusText = '已发布';
      let _status: 'default' | 'error' | 'success' | 'warning' | 'processing' | undefined =
        'success';
      if (rd.status == 'INTIMATE') {
        statusText = '私密';
        _status = 'success';
      } else if (rd.status == 'DRAFT') {
        statusText = '私密';
        _status = 'warning';
      } else if (rd.status == 'RECYCLE') {
        statusText = '回收站';
        _status = 'error';
      }
      return <Badge status={_status} text={statusText} />;
    },
  },
  {
    title: '分类',
    dataIndex: 'categoryId',
    align: 'left',
    sorter: false,
    customRender: ({ record }) => {
      const rd = record as Post;
      const categories = rd?.categories;
      return h(Row, { gutter: [6, 6] }, () =>
        categories?.map((cate) => {
          return h(Col, {}, () => h(Tag, { color: 'processing' }, () => cate.name));
        }),
      );
    },
  },
  {
    title: '标签',
    dataIndex: 'tags',
    align: 'left',
    sorter: false,
    customRender: ({ record }) => {
      const rd = record as Post;
      const tags = rd?.tags;
      return h(Row, { gutter: [6, 6] }, () =>
        tags?.map((tag) => {
          return h(Col, {}, () => h(Tag, { color: tag.color }, () => tag.name));
        }),
      );
    },
  },
  {
    title: '评论',
    dataIndex: 'commentCount',
    sorter: false,
    width: 80,
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
    width: 80,
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
