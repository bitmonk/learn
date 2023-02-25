load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "sum",
    srcs = ["sum.go"],
    importpath = "learn/golang/with-tests/sum",
    visibility = ["//visibility:public"],
)

go_test(
    name = "sum_test",
    srcs = ["sum_test.go"],
    embed = [":sum"],
)
