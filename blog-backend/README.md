# 克拉托斯博客 - Golang后端 - Kratos Blog Backend

一个Golang的博客系统/CMS。

- 后端基于 golang微服务框架 [go-kratos](https://go-kratos.dev/)
- 前端基于 [VUE3](https://vuejs.org/) + [TypeScript](https://www.typescriptlang.org/)

## 技术栈

* [Kratos](https://go-kratos.dev/) -- B站微服务框架
* [Consul](https://www.consul.io/) -- 服务发现和配置管理
* [OpenTelemetry](https://opentelemetry.io/) -- 分布式可观察系统
* [Jaeger](https://www.jaegertracing.io/) -- 分布式跟踪的存储和展示
* [Ent](https://entgo.io/) -- Facebook ORM 数据库实体框架
* [Redis](https://redis.io/) -- 非关系型数据库
* [PostgreSQL](https://www.postgresql.org/) -- 关系型数据库
* [Casbin](https://casbin.org/) -- 访问控制框架
* [Wire](https://github.com/google/wire) -- 依赖注入框架
* [Swagger](https://github.com/swaggo/swag) -- RESTful API 文档
* [MinIO](https://min.io/) -- 对象存储服务器

## 生成API

使用[buf.build](https://buf.build/)进行构建，buf的安装文档请参考官方文档：<https://docs.buf.build/installation>

在`blog-backend`根目录下执行`buf generate`命令：

1. 生成全部

    ```bash
    buf generate
    ```

2. 生成OpenAPI v3文档

    ```bash
    buf generate --path api/admin/service/v1 --template api/admin/service/buf.openapi.gen.yaml
    ```

## Docker部署开发服务器

### Consul

```shell
docker pull bitnami/consul:latest

docker run -itd \
    --name consul-server-standalone \
    -p 8300:8300 \
    -p 8500:8500 \
    -p 8600:8600/udp \
    -e CONSUL_BIND_INTERFACE='eth0' \
    -e CONSUL_AGENT_MODE=server \
    -e CONSUL_ENABLE_UI=true \
    -e CONSUL_BOOTSTRAP_EXPECT=1 \
    -e CONSUL_CLIENT_LAN_ADDRESS=0.0.0.0 \
    bitnami/consul:latest
```

- 管理后台: <http://localhost:8500>

### Jaeger

```shell
docker pull jaegertracing/all-in-one:latest

docker run -d \
    --name jaeger \
    -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
    -e COLLECTOR_OTLP_ENABLED=true \
    -p 6831:6831/udp \
    -p 6832:6832/udp \
    -p 5778:5778 \
    -p 16686:16686 \
    -p 4317:4317 \
    -p 4318:4318 \
    -p 14250:14250 \
    -p 14268:14268 \
    -p 14269:14269 \
    -p 9411:9411 \
    jaegertracing/all-in-one:latest
```

- API：<http://localhost:14268/api/traces>
- Zipkin API：<http://localhost:9411/api/v2/spans>
- 后台: <http://localhost:16686>

### PostgreSQL

```shell
docker pull bitnami/postgresql:latest
docker pull bitnami/postgresql-repmgr:latest
docker pull bitnami/pgbouncer:latest
docker pull bitnami/pgpool:latest
docker pull bitnami/postgres-exporter:latest

docker run -itd \
    --name postgres-test \
    -p 5432:5432 \
    -e POSTGRES_PASSWORD=123456 \
    bitnami/postgresql:latest

docker exec -it postgres-test "apt update"
```

```postgresql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "postgis";

SELECT version();
SELECT postgis_full_version();
```

- 默认账号：postgres
- 默认密码：123456

### Redis

```shell
docker pull bitnami/redis:latest
docker pull bitnami/redis-exporter:latest

docker run -itd \
    --name redis-test \
    -p 6379:6379 \
    -e ALLOW_EMPTY_PASSWORD=yes \
    bitnami/redis:latest
```

### MinIO

```shell
docker pull bitnami/minio:latest
docker pull bitnami/minio-client:latest

docker network create app-tier --driver bridge

# MINIO_ROOT_USER最少3个字符
# MINIO_ROOT_PASSWORD最少8个字符
# 第一次运行的时候,服务会自动关闭,手动再启动就可以正常运行了.
docker run -itd \
    --name minio-server \
    -p 9000:9000 \
    -p 9001:9001 \
    --env MINIO_ROOT_USER="root" \
    --env MINIO_ROOT_PASSWORD="123456789" \
    --env MINIO_DEFAULT_BUCKETS='my-bucket' \
    --env MINIO_FORCE_NEW_KEYS="yes" \
    --env BITNAMI_DEBUG=true \
    --network app-tier \
    bitnami/minio:latest

docker run -itd \
    --name minio-client \
    --env MINIO_SERVER_HOST="minio-server" \
    --env MINIO_SERVER_ACCESS_KEY="root" \
    --env MINIO_SERVER_SECRET_KEY="123456789" \
    --network app-tier \
    bitnami/minio-client:latest
```

- 管理后台: <http://localhost:9001/login>
