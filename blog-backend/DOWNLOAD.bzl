load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def download_package():
    # Load go bazel rules
    if not native.existing_rule("io_bazel_rules_go"):
        http_archive(
            name = "io_bazel_rules_go",
            sha256 = "56d8c5a5c91e1af73eca71a6fab2ced959b67c86d12ba37feedb0a2dfea441a6",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.37.0/rules_go-v0.37.0.zip",
                "https://github.com/bazelbuild/rules_go/releases/download/v0.37.0/rules_go-v0.37.0.zip",
            ],
        )

    # Load gazelle
    if not native.existing_rule("bazel_gazelle"):
        http_archive(
            name = "bazel_gazelle",
            sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
                "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
            ],
        )

    # Load Bazel/starlark utilities.
    if not native.existing_rule("bazel_skylib"):
        http_archive(
            name = "bazel_skylib",
            sha256 = "74d544d96f4a5bb630d465ca8bbcfe231e3594e5aae57e1edbf17a6eb3ca2506",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
                "https://github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
            ],
        )

    # Load buf.build rule.
    if not native.existing_rule("rules_buf"):
        http_archive(
            name = "rules_buf",
            sha256 = "523a4e06f0746661e092d083757263a249fedca535bd6dd819a8c50de074731a",
            strip_prefix = "rules_buf-0.1.1",
            urls = [
                "https://github.com/bufbuild/rules_buf/archive/refs/tags/v0.1.1.zip",
            ],
        )

    # Load protobuf dependencies.
    if not native.existing_rule("rules_proto"):
        http_archive(
            name = "rules_proto",
            sha256 = "66bfdf8782796239d3875d37e7de19b1d94301e8972b3cbd2446b332429b4df1",
            strip_prefix = "rules_proto-4.0.0",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
                "https://github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
            ],
        )

    # Load grpc dependencies.
    if not native.existing_rule("rules_proto_grpc"):
        http_archive(
            name = "rules_proto_grpc",
            sha256 = "fb7fc7a3c19a92b2f15ed7c4ffb2983e956625c1436f57a3430b897ba9864059",
            strip_prefix = "rules_proto_grpc-4.3.0",
            urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/4.3.0.tar.gz"],
        )

    if not native.existing_rule("build_stack_rules_proto"):
        # Release: v2.0.1
        # TargetCommitish: master
        # Date: 2022-10-20 02:38:27 +0000 UTC
        # URL: https://github.com/stackb/rules_proto/releases/tag/v2.0.1
        # Size: 2071295 (2.1 MB)
        http_archive(
            name = "build_stack_rules_proto",
            sha256 = "ac7e2966a78660e83e1ba84a06db6eda9a7659a841b6a7fd93028cd8757afbfb",
            strip_prefix = "rules_proto-2.0.1",
            urls = ["https://github.com/stackb/rules_proto/archive/v2.0.1.tar.gz"],
        )

    # protoc
    if not native.existing_rule("com_google_protobuf"):
        http_archive(
            name = "com_google_protobuf",
            sha256 = "bc3dbf1f09dba1b2eb3f2f70352ee97b9049066c9040ce0c9b67fb3294e91e4b",
            strip_prefix = "protobuf-3.15.5",
            # latest, as of 2021-03-08
            urls = [
                "https://github.com/protocolbuffers/protobuf/archive/v3.15.5.tar.gz",
                "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.15.5.tar.gz",
            ],
        )
