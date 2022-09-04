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
  const name = Random.first();
  return {
    id: Random.id(),
    name: name,
    slug: name.toLowerCase(),
    fullPath: '/tags/' + name.toLowerCase(),
    color: Random.color(),
    thumbnail: Random.image(),
    description: Random.cparagraph(),
    postCount: Random.integer(0, 1000),
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
    thumbnail: Random.image(),
    postCount: Random.integer(0, 1000),
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

export function createAttachment() {
  return {
    id: Random.id(),
    name: Random.first(),
    path: Random.image(),
    thumbPath: Random.image(),
    mediaType: 'image/jpeg',
    suffix: 'jpeg',
    width: 6016,
    height: 4016,
    size: 476059,
    type: 'LOCAL',
    createTime: Random.datetime(),
  };
}

export function createAttachments(count) {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createAttachment());
  }
  return items;
}
