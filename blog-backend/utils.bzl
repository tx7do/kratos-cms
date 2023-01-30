load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")

# 发布服务
def publish_service(service_name, repository_name = "", repository_version = "latest", publish = False):
    image_name = "{}-service-image".format(service_name)

    # 为服务的编译目标定义一个别名
    native.alias(
        name = "{}-service".format(service_name),
        actual = "//app/{}/service/cmd/server:server".format(service_name),
        visibility = ["//visibility:private"],
    )

    # 将配置文件打包
    native.filegroup(
        name = "{}-service-configs".format(service_name),
        srcs = native.glob(["app/{}/service/configs/**".format(service_name)]),
        visibility = ["//visibility:private"],
    )

    # 生成Docker镜像
    container_image(
        name = image_name,
        base = "@alpine_linux_amd64//image",
        entrypoint = ["/server"],
        directory = "/app/{}/service/bin".format(service_name),
        # 需要打包进镜像去的文件
        files = [
            ":{}-service".format(service_name),
            ":{}-service-configs".format(service_name),
        ],
        # 资源库的用户名
        repository = repository_name,
    )

    # 推送到DockerHub
    if publish:
        container_push(
            name = "{}-push".format(image_name),
            # 镜像的格式，可选项：Docker、OCI；默认为：Docker。
            format = "Docker",
            # 要被推送的镜像
            image = ":{}".format(image_name),
            # 镜像库的注册链接
            registry = "index.docker.io",
            ## 目标镜像库中的镜像名
            repository = "{}/kratos-blog_{}-service".format(repository_name, service_name),
            # 镜像标签
            tag = repository_version,
        )
