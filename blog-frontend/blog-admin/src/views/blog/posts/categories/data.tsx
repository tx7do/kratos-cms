import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { h } from 'vue';
import { Avatar, Badge } from 'ant-design-vue';
import { Category } from '/&/category';

export const columns: BasicColumn[] = [
  {
    title: '封面图',
    dataIndex: 'thumbnail',
    sorter: false,
    width: 40,
    customRender: ({ record }) => {
      const cate = record as Category;
      return h(Avatar, { src: cate.thumbnail });
    },
  },
  {
    title: '名称',
    dataIndex: 'name',
    width: 200,
    align: 'left',
  },
  {
    title: '关联文章数',
    dataIndex: 'postCount',
    sorter: false,
    width: 120,
    customRender: ({ record }) => {
      const cate = record as Category;
      return h(Badge, {
        count: cate.postCount,
        numberStyle: { backgroundColor: 'volcano' },
        overflowCount: 10000,
      });
    },
  },
  {
    title: '创建时间',
    dataIndex: 'createTime',
    width: 180,
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'name',
    label: '名称',
    component: 'Input',
    colProps: { span: 8 },
  },
];

export const formSchema: FormSchema[] = [
  {
    field: 'name',
    label: '名称',
    component: 'Input',
    required: true,
  },

  {
    field: 'parentId',
    label: '上级分类',
    component: 'TreeSelect',
    componentProps: {
      fieldNames: {
        label: 'name',
        key: 'id',
        value: 'id',
      },
      getPopupContainer: () => document.body,
    },
  },

  {
    field: 'priority',
    label: '排序',
    component: 'InputNumber',
    required: true,
  },
];
