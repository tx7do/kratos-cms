<template>
  <PageWrapper dense contentFullHeight>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button preIcon="ant-design:plus" type="primary" @click="handleCreate">
          创建标签
        </a-button>
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

  import { columns } from './data';
  import { ListTag } from '/@/api/blog/tag';

  const go = useGo();

  const [registerTable] = useTable({
    title: '全部标签',
    api: ListTag,
    columns,
    rowKey: 'id',
    showIndexColumn: false,
    canResize: true,
    useSearchForm: false,
    showTableSetting: true,
    bordered: true,
    actionColumn: {
      width: 120,
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
    ];
  }

  function handleCreate(record: EditRecordRow) {
    go('/tags/write');
  }
  function handleEdit(record: EditRecordRow) {
    go('/tags/edit/' + record.id);
  }
  function handleDelete(record: EditRecordRow) {}
</script>
