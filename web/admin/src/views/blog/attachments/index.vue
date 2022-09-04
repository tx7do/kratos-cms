<template>
  <PageWrapper title="附件列表">
    <template #extra>
      <Space size="middle">
        <Button preIcon="ant-design:upload-outlined" type="primary" @click="handleClickUpload"
          >上传</Button
        >
      </Space>
    </template>

    <CardList :params="params" :api="ListAttachment" />
  </PageWrapper>
</template>

<script lang="ts" setup>
  import CardList from './card-list.vue';
  import { Button } from '/@/components/Button';
  import { PageWrapper } from '/@/components/Page';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { Space } from 'ant-design-vue';
  import { ListAttachment } from '/@/api/blog/attachment';

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
