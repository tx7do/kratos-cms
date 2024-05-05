load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_layer", "container_push")

# 发布服务
def publish_service(service_name, repository_name = "", repository_version = "", publish = False):
    service_new_name = "{}-service".format(service_name)
    image_name = "{}-service-image".format(service_name)
    conf_file_group_name = "{}-service-configs".format(service_name)
    conf_layer_name = "{}-service-configs-layer".format(service_name)

    app_path = "/app/{}/service/bin".format(service_name)
    conf_path = "/app/{}/service/configs".format(service_name)

    if repository_version == "":
        repository_version = "{BUILD_TIMESTAMP}"

    # 为服务的编译目标定义一个别名
    native.alias(
        name = service_new_name,
        actual = "//app/{}/service/cmd/server:server".format(service_name),
        visibility = ["//visibility:private"],
    )

    # 将配置文件打包
    native.filegroup(
        name = conf_file_group_name,
        srcs = native.glob(["app/{}/service/configs/**".format(service_name)]),
        visibility = ["//visibility:public"],
    )

    container_layer(
        name = conf_layer_name,
        directory = "/{}".format(conf_path),
        files = [
            "//:{}".format(conf_file_group_name),
        ],
        mode = "0o755",
        visibility = ["//visibility:public"],
    )

    # 生成Docker镜像
    container_image(
        # 镜像名，可用于：编译目标名，镜像标签。
        name = image_name,

        # OS
        base = "@slim_linux_amd64//image",

        # 容器启动时运行的命令
        # https://docs.docker.com/engine/reference/builder/#entrypoint
        entrypoint = [
            "./server",
            "-conf",
            "../configs",
            "-chost",
            "host.docker.internal:8500",
            "-ctype",
            "consul",
        ],

        # 存放files/tars/debs文件的路径
        directory = app_path,

        # https://docs.docker.com/engine/reference/builder/#workdir
        workdir = app_path,

        # https://docs.docker.com/engine/reference/builder/#user
        # user = "appuser",

        # 需要打包进镜像去的文件
        files = [
            "//:{}".format(service_new_name),
        ],
        layers = ["//:{}".format(conf_layer_name)],

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
            image = "//:{}".format(image_name),
            # 镜像库的注册链接
            registry = "index.docker.io",
            ## 目标镜像库中的镜像名
            repository = "{}/kratoscms-{}-service".format(repository_name, service_name),
            # 镜像标签
            tag = repository_version,
        )
