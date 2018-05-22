import flask
import tensorflow as tf
import python_task.task_pb2 as task_pb2

# bazel build --define grpc_no_ares=true python_task:app

app = flask.Flask(__name__)

@app.route("/")
def hello():
    task = task_pb2.Task()
    task.name = "sdfds"
    print task

    hello = tf.constant('Hello, TensorFlow!')
    sess = tf.Session()
    print(sess.run(hello))
    return "Python Home \n"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
