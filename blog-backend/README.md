# 克拉托斯博客 - Golang后端 - Kratos Blog Backend

一个Golang的博客系统/CMS。

- 后端基于 golang微服务框架 [go-kratos](https://go-kratos.dev/)
- 前端基于 [VUE3](https://vuejs.org/) + [TypeScript](https://www.typescriptlang.org/)

## 技术栈

* [Kratos](https://go-kratos.dev/) -- B站微服务框架
* [Consul](https://www.consul.io/) -- 服务发现和配置管理
* [OpenTelemetry](https://opentelemetry.io/) -- 分布式可观察系统
* [Wire](https://github.com/google/wire) -- 依赖注入框架
* [OpenAPI](https://www.openapis.org/) -- RESTful API 文档
* [Ent](https://entgo.io/) -- Facebook ORM 数据库实体框架
* [Redis](https://redis.io/) -- 非关系型数据库
* [PostgreSQL](https://www.postgresql.org/) -- 关系型数据库
* [MinIO](https://min.io/) -- 对象存储服务器

## 生成Protobuf API

使用[buf.build](https://buf.build/)进行Protobuf API的构建。

### 安装工具

具体安装方法请参见：[Kratos微服务框架API工程化指南](https://juejin.cn/post/7191095845096259641)

### 执行生成命令

在`blog-backend`根目录下执行命令：

- 更新buf.lock

    ```bash
    buf mod update
    ```

- 生成GO代码

    ```bash
    buf generate
    ```

- 生成OpenAPI v3文档

    ```bash
    buf generate --path api/admin/service/v1 --template api/admin/service/v1/buf.openapi.gen.yaml
    ```

## Make构建

请在`app/{服务名}/service`下执行：

- 初始化开发环境

   ```bash
   make init
   ```

- 生成API

   ```bash
   make api
   ```

- 生成服务器配置结构代码

   ```bash
   make conf
   ```

- 生成ent代码

   ```bash
   make ent
   ```

- 生成wire代码

   ```bash
   make wire
   ```

- 构建程序

   ```bash
   make build
   ```

- 构建Docker镜像

   ```bash
   make docker
   ```

## Bazel构建

使用[bazel.build](https://bazel.build/)进行服务器程序的构建。

### 安装Bazel

如何安装bazel.build的文档，请参考官方文档：<https://bazel.build/install>。

- 安装Gazelle

   ```shell
   go install github.com/bazelbuild/bazel-gazelle/cmd/gazelle@latest
   ```

## Bazel命令使用

在`blog-backend`根目录下执行命令：

- 更新GO依赖库引入的构建配置文件repos.bzl

   ```bash
   gazelle update-repos \
           --from_file=go.mod --to_macro=repos.bzl%go_repositories \
           --build_file_generation=on --build_file_proto_mode=disable \
           --prune
   ```
   
   或者
   
   ```bash
   bazel run //:gazelle-update-repos
   bazel run //:gazelle-update-buf-repos
   ```

- 生成构建配置文件BUILD.bazel

   ```bash
   bazel run //:gazelle
   ```

- 构建单个服务

  ```bash
  bazel build //app/admin/service/cmd/server:server
  bazel build //app/comment/service/cmd/server:server
  bazel build //app/content/service/cmd/server:server
  bazel build //app/file/service/cmd/server:server
  bazel build //app/user/service/cmd/server:server
  ```

  或者

  ```bash
  bazel build //:admin-server
  bazel build //:comment-server
  bazel build //:content-server
  bazel build //:file-server
  bazel build //:user-server
  ```

- 构建全部服务

  ```bash
  bazel build //:all
  ```

## Docker部署
