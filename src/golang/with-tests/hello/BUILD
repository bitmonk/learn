load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "hello",
    srcs = ["hello.go"],
    importpath = "github.com/bitmonk/learn/golang/with-tests/hello",
    visibility = ["//visibility:public"],
)

go_test(
    name = "hello-test",
    srcs = ["hello_test.go"],
    visibility = ["//visibility:public"],
    embed = [":hello"],
)
