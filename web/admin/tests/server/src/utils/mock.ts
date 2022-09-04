import { Random } from 'mockjs';

export const postStatus = ['PUBLISHED', 'INTIMATE', 'DRAFT', 'RECYCLE'];

export function createPost() {
  return {
    id: Random.id(),
    title: Random.ctitle(),
    visits: Random.integer(0, 101000),
    commentCount: Random.integer(0, 10100),
    createTime: Random.datetime(),
    editorType: 'MARKDOWN',
    status: postStatus[Random.integer(0, 3)],
    tags: createTags(Random.integer(1, 6)),
    categories: createCategories(Random.integer(1, 3)),
  };
}

export function createPosts(count) {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createPost());
  }
  return items;
}

export function createTag() {
  return {
    id: Random.id(),
    name: Random.first(),
    slug: Random.first(),
    color: Random.color(),
    createTime: Random.datetime(),
  };
}

export function createTags(count) {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createTag());
  }
  return items;
}

export function createCategory() {
  return {
    id: Random.id(),
    name: Random.first(),
    slug: Random.first(),
    createTime: Random.datetime(),
  };
}

export function createCategories(count) {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createCategory());
  }
  return items;
}
