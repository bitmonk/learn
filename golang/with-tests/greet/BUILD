load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "greet",
    srcs = ["greet.go"],
    importpath = "learn/golang/with-tests/greet",
    visibility = ["//visibility:public"],
)

go_test(
    name = "greet_test",
    srcs = ["greet_test.go"],
    embed = [":greet"],
)
