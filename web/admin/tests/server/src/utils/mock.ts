import { Random } from 'mockjs';

import { Post } from '/&/post';
import { Tag } from '/&/tag';
import { Category } from '/&/category';
import { Attachment } from '/&/attachment';
import { Comment } from '/&/comment';
import { Link } from '/&/link';
import { Menu } from '/&/menu';
import { Photo } from '/&/photo';

Random.extend({
  postStatus: function () {
    const postStatusList = ['PUBLISHED', 'INTIMATE', 'DRAFT', 'RECYCLE'];
    return this.pick(postStatusList);
  },
});

export function createPost(): Post {
  return {
    id: Random.increment(),
    title: Random.ctitle(),
    visits: Random.integer(0, 101000),
    commentCount: Random.integer(0, 10100),
    createTime: Random.datetime(),
    editorType: 'MARKDOWN',
    status: Random.postStatus(),
    tags: createTags(Random.integer(1, 6)),
    categories: createCategories(Random.integer(1, 3)),
  };
}

export function createPosts(count): Post[] {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createPost());
  }
  return items;
}

export function createTag(): Tag {
  const name = Random.first();
  return {
    id: Random.increment(),
    name: name,
    slug: name.toLowerCase(),
    fullPath: '/tags/' + name.toLowerCase(),
    color: Random.color(),
    thumbnail: Random.image(),
    postCount: Random.integer(0, 1000),
    createTime: Random.datetime(),
  };
}

export function createTags(count): Tag[] {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createTag());
  }
  return items;
}

export function createCategory(): Category {
  return {
    id: Random.increment(),
    name: Random.first(),
    slug: Random.first(),
    thumbnail: Random.image(),
    postCount: Random.integer(0, 1000),
    createTime: Random.datetime(),
  };
}

export function createCategories(count): Category[] {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createCategory());
  }
  return items;
}

export function createAttachment(): Attachment {
  return {
    id: Random.increment(),
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

export function createAttachments(count): Attachment[] {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createAttachment());
  }
  return items;
}

export function createComment(): Comment {
  return {};
}

export function createComments(count): Comment[] {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createComment());
  }
  return items;
}

export function createLink(): Link {
  return {};
}

export function createLinks(count): Link[] {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createLink());
  }
  return items;
}

export function createMenu(): Menu {
  return {};
}

export function createMenus(count): Menu[] {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createMenu());
  }
  return items;
}

export function createPhoto(): Photo {
  return {};
}

export function createPhotos(count): Photo[] {
  const items: any[] = [];
  for (let index = 0; index < count; index++) {
    items.push(createPhoto());
  }
  return items;
}
