<template>
  <PageWrapper dense contentFullHeight>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button preIcon="ant-design:plus" type="primary" @click="handleCreate"> 写文章 </a-button>
        <a-button preIcon="ant-design:delete" @click="handleTrash"> 回收站 </a-button>
      </template>
      <template #bodyCell="{ column, record }">
        <TableAction
          v-if="(column as BasicColumn).dataIndex === 'action'"
          :actions="createActions(record)"
        />
      </template>
    </BasicTable>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import {
    BasicTable,
    useTable,
    TableAction,
    ActionItem,
    EditRecordRow,
    BasicColumn,
  } from '/@/components/Table';
  import { PageWrapper } from '/@/components/Page';
  import { useGo } from '/@/hooks/web/usePage';

  import { columns, searchFormSchema } from './data';
  import { ListPost } from '/@/api/blog/post';

  const go = useGo();

  const [registerTable] = useTable({
    title: '列表',
    api: ListPost,
    columns,
    rowKey: 'id',
    formConfig: {
      labelWidth: 120,
      schemas: searchFormSchema,
      autoSubmitOnEnter: true,
    },
    showIndexColumn: false,
    canResize: false,
    useSearchForm: true,
    showTableSetting: true,
    bordered: true,
    actionColumn: {
      width: 180,
      title: '操作',
      dataIndex: 'action',
    },
  });

  function createActions(record: EditRecordRow): ActionItem[] {
    return [
      {
        label: '编辑',
        onClick: handleEdit.bind(null, record),
      },
      {
        label: '删除',
        onClick: handleDelete.bind(null, record),
      },
      {
        label: '设置',
        onClick: handleSetting.bind(null, record),
      },
    ];
  }

  function handleCreate(record: EditRecordRow) {
    go('/posts/write');
  }
  function handleTrash(record: EditRecordRow) {}

  function handleEdit(record: EditRecordRow) {
    go('/posts/edit/' + record.id);
  }
  function handleDelete(record: EditRecordRow) {}
  function handleSetting(record: EditRecordRow) {}
</script>
