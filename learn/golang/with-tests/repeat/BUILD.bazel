load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "repeat",
    srcs = ["repeat.go"],
    importpath = "learn/golang/with-tests/repeat",
    visibility = ["//visibility:public"],
)

go_test(
    name = "repeat_test",
    srcs = ["repeat_test.go"],
    embed = [":repeat"],
)
