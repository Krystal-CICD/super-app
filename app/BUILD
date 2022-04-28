package(default_visibility = ["//visibility:public"])

# Go boilerplate
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_image(
    name = "super-app",
    srcs = ["hello.go"],
    base = "//base:base_root_amd64_debian11",
)