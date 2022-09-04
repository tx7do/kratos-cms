<template>
  <div class="p-2">
    <div class="p-2 bg-white">
      <List
        :grid="{ gutter: 14, xs: 1, sm: 2, md: 4, lg: 4, xl: 6, xxl: 6 }"
        :data-source="data"
        :pagination="paginationProp"
      >
        <template #header>
          <div class="flex justify-end space-x-2">
            <slot name="header"></slot>
            <Tooltip @click="fetch">
              <template #title>刷新</template>
              <Button>
                <RedoOutlined />
              </Button>
            </Tooltip>
          </div>
        </template>
        <template #renderItem="{ item }">
          <ListItem>
            <Card hoverable>
              <template #cover>
                <Image :alt="item.name" :src="item.thumbPath" style="width: 222px; height: 130px" />
              </template>
              <CardMeta>
                <template #description>{{ item.name }}</template>
              </CardMeta>
            </Card>
          </ListItem>
        </template>
      </List>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
  import { RedoOutlined } from '@ant-design/icons-vue';
  import { List, Card, Image, Typography, Tooltip, Slider } from 'ant-design-vue';
  import { propTypes } from '/@/utils/propTypes';
  import { Button } from '/@/components/Button';
  import { isFunction } from '/@/utils/is';
  import { useSlider, grid } from './data';

  const ListItem = List.Item;
  const CardMeta = Card.Meta;
  // 获取slider属性
  const sliderProp = computed(() => useSlider(4));
  // 组件接收参数
  const props = defineProps({
    // 请求API的参数
    params: propTypes.object.def({}),
    //api
    api: propTypes.func,
  });
  //暴露内部方法
  const emit = defineEmits(['getMethod', 'delete']);
  //数据
  const data = ref([]);

  function sliderChange(n) {
    pageSize.value = n * 4;
    fetch();
  }

  // 自动请求并暴露内部方法
  onMounted(() => {
    fetch();
    emit('getMethod', fetch);
  });

  async function fetch(p = {}) {
    const { api, params } = props;
    if (api && isFunction(api)) {
      const res = await api({ ...params, page: page.value, pageSize: pageSize.value, ...p });
      data.value = res.items;
      total.value = res.total;
    }
  }

  //分页相关
  const page = ref(1);
  const pageSize = ref(36);
  const total = ref(0);
  const paginationProp = ref({
    showSizeChanger: false,
    showQuickJumper: true,
    pageSize,
    current: page,
    total,
    showTotal: (total) => `总 ${total} 条`,
    onChange: pageChange,
    onShowSizeChange: pageSizeChange,
  });

  function pageChange(p, pz) {
    page.value = p;
    pageSize.value = pz;
    fetch();
  }

  function pageSizeChange(_current, size) {
    pageSize.value = size;
    fetch();
  }

  async function handleDelete(id) {
    emit('delete', id);
  }
</script>
