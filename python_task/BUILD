load(
  "@io_bazel_rules_python//python:python.bzl",
  "py_binary", "py_library", "py_test",
)
load("@python_task_import//:requirements.bzl", python_task_requirements="all_requirements")
load("@org_pubref_rules_protobuf//python:rules.bzl", "py_proto_library")
load("@io_bazel_rules_docker//python:image.bzl", "py_image")

py_proto_library(
  name = "task_proto",
  protos = ["task.proto"],
  with_grpc = True,
)

py_binary(
  name = "app",
	srcs = ["app.py"],
	deps = [
    ":task_proto",
    "@org_tensorflow//tensorflow:tensorflow_py",
	] + python_task_requirements,
)

py_test(
  name = "app_test",
	srcs = ["app_test.py"],
	deps = [],
)


py_image(
  name = "app_image",
	srcs = ["app.py"],
	deps = [
    ":task_proto",
    "@org_tensorflow//tensorflow:tensorflow_py",
	] + python_task_requirements,
  main = "app.py",
)
