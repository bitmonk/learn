load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "wallet",
    srcs = ["wallet.go"],
    importpath = "learn/golang/with-tests/wallet",
    visibility = ["//visibility:public"],
)

go_test(
    name = "wallet_test",
    srcs = ["wallet_test.go"],
    embed = [":wallet"],
)
