# TensorFlow Serving external dependencies that can be loaded in WORKSPACE
# files.

load('@org_tensorflow//tensorflow:workspace.bzl', 'tf_workspace')

def tf_serving_workspace():
  tf_workspace(path_prefix = "", tf_repo_name = "org_tensorflow")
