load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "shapes",
    srcs = ["shapes.go"],
    importpath = "learn/golang/with-tests/shapes",
    visibility = ["//visibility:public"],
)

go_test(
    name = "shapes_test",
    srcs = ["shapes_test.go"],
    embed = [":shapes"],
)
