<template>
  <PageWrapper dense contentFullHeight>
    <BasicTable @register="registerTable">
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

  import { postColumns, postSearchFormSchema } from './data';
  import { ListComment } from '/@/api/blog/comment';

  const [registerTable] = useTable({
    //title: '列表',
    api: ListComment,
    columns: postColumns,
    rowKey: 'id',
    formConfig: {
      labelWidth: 120,
      schemas: postSearchFormSchema,
      autoSubmitOnEnter: true,
    },
    showIndexColumn: false,
    canResize: false,
    useSearchForm: true,
    showTableSetting: false,
    bordered: true,
    actionColumn: {
      width: 100,
      title: '操作',
      dataIndex: 'action',
    },
  });

  function createActions(record: EditRecordRow): ActionItem[] {
    return [
      {
        label: '回复',
        onClick: handleReply.bind(null, record),
      },
      {
        label: '回收站',
        onClick: handleDelete.bind(null, record),
      },
    ];
  }

  function handleReply(record: EditRecordRow) {}
  function handleDelete(record: EditRecordRow) {}
</script>

<style lang="less" scoped></style>
