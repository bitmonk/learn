load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "adder",
    srcs = ["adder.go"],
    importpath = "learn/golang/with-tests/adder",
    visibility = ["//visibility:public"],
)

go_test(
    name = "adder_test",
    srcs = ["adder_test.go"],
    embed = [":adder"],
)
