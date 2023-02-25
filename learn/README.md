# learn
learnings

# Requirements

## Bazel

This repo is configured to use bazel as a build tool.

It expects Bazel 6.0.0 at the time of this writing, the recommended
way to install it is to download
[Bazelisk](https://github.com/bazelbuild/bazelisk) and place it in your PATH
as `bazel`.

# c
c/
 - cprogramming.com/ cprogramming.com tutorial

# golang
golang/
 - scratch/ : quick scratch projects
 - tutorial/ : official tutorial
 - with-tests/ : learn go with tests

## Bazel example

To run, for instance, the tests in the "greet" package in the
with-tests collection:

```
bazel test //golang/with-tests/greet:greet_test
```