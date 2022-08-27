<template>
  <PageWrapper title="附件列表">
    <template #extra>
      <a-space size="middle">
        <Button preIcon="ant-design:upload-outlined" type="primary" @click="handleClickUpload"
          >上传</Button
        >
      </a-space>
    </template>

    <CardList :params="params" />
  </PageWrapper>
</template>

<script lang="ts" setup>
  import CardList from './card-list.vue';
  import { Button } from '/@/components/Button';
  import { PageWrapper } from '/@/components/Page';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { Space } from 'ant-design-vue';

  const ASpace = Space;

  const { notification } = useMessage();
  // 请求api时附带参数
  const params = {};

  let reload = () => {};

  // 获取内部fetch方法;
  function getMethod(m: any) {
    reload = m;
  }

  function handleClickUpload() {}
</script>
