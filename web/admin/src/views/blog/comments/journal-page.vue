<template>
  <PageWrapper dense contentFullHeight fixedHeight>
    <BasicTable @register="registerTable" :searchInfo="postSearchFormSchema">
      <template #action="{ record }">
        <TableAction :actions="createActions(record)" />
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
  } from '/@/components/Table';
  import { PageWrapper } from '/@/components/Page';
  import { useGo } from '/@/hooks/web/usePage';

  import { journalColumns, postSearchFormSchema } from './data';

  const go = useGo();

  const [registerTable, { reload }] = useTable({
    //title: '列表',
    //api: ListPost,
    columns: journalColumns,
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
      slots: {
        customRender: 'action',
      },
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
