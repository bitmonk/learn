load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "countdown",
    srcs = ["countdown.go"],
    importpath = "learn/golang/with-tests/countdown",
    visibility = ["//visibility:public"],
)

go_test(
    name = "countdown_test",
    srcs = ["countdown_test.go"],
    embed = [":countdown"],
)
