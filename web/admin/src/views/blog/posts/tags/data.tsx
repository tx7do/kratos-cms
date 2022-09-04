import { BasicColumn } from '/@/components/Table';
import { h } from 'vue';
import { Badge, Tag, Avatar } from 'ant-design-vue';
import { Tag as TAG } from '/&/tag';

export const columns: BasicColumn[] = [
  {
    title: '封面图',
    dataIndex: 'thumbnail',
    sorter: false,
    width: 80,
    customRender: ({ record }) => {
      const tag = record as TAG;
      return h(Avatar, { src: tag.thumbnail });
    },
  },
  {
    title: '名称',
    dataIndex: 'name',
    align: 'center',
    sorter: false,
    width: 150,
    customRender: ({ record }) => {
      const tag = record as TAG;
      return h(Tag, { color: tag.color }, () => tag.name);
    },
  },
  {
    title: '别名',
    dataIndex: 'slug',
    sorter: false,
    width: 150,
  },
  {
    title: '全路径名',
    dataIndex: 'fullPath',
    sorter: false,
  },
  {
    title: '关联文章数',
    dataIndex: 'postCount',
    sorter: false,
    width: 120,
    customRender: ({ record }) => {
      const tag = record as TAG;
      return h(Badge, {
        count: tag.postCount,
        numberStyle: { backgroundColor: 'volcano' },
        overflowCount: 10000,
      });
    },
  },
  {
    title: '创建时间',
    dataIndex: 'createTime',
    sorter: false,
  },
];
