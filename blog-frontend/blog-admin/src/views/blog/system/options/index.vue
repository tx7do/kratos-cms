<template>
  <PageWrapper dense contentFullHeight fixedHeight>
    <Card>
      <Tabs>
        <template v-for="item in tabList" :key="item.key">
          <TabPane :tab="item.name">
            <component :is="item.component" />
          </TabPane>
        </template>
      </Tabs>
    </Card>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { Tabs, Card } from 'ant-design-vue';
  import { PageWrapper } from '/@/components/Page';
  import { tabList } from './data';

  const TabPane = Tabs.TabPane;
</script>

<style lang="less" scoped></style>
