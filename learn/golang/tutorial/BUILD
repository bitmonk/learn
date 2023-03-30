load("@io_bazel_rules_go//go:def.bzl", "go_binary")

# gazelle:ignore

go_binary(
    name = "hello",
    srcs = ["hello.go"],
    deps = ["//golang/tutorial/greetings:greetings"],
    visibility = ["//visibility:public"],
)
