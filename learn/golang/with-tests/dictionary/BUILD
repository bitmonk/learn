load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "dictionary",
    srcs = ["dictionary.go"],
    importpath = "learn/golang/with-tests/dictionary",
    visibility = ["//visibility:public"],
)

go_test(
    name = "dictionary_test",
    srcs = ["dictionary_test.go"],
    embed = [":dictionary"],
)
