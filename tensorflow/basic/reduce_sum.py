# coding=utf-8
import tensorflow as tf

# reduce_sum
x = [[2, 2], [1, 2], [1, 2]]
with tf.Session() as sess:
    print(sess.run(tf.reduce_sum(x)))  # 10
    print(sess.run(tf.reduce_sum(x, 0)))  # [4 6]
    print(sess.run(tf.reduce_sum(x, 1)))  # [4 3 3]
    print(sess.run(tf.reduce_sum(x, 1, True)))  # [[4] [3 [3]]
    print(sess.run(tf.reduce_sum(x, [0, 1])))  # 10
    print(sess.run(tf.reduce_sum(x, [1, 0])))  # 10
    print(sess.run(tf.reduce_sum(x, [1, 0], True)))  # [[10]]
# Args:
#   input_tensor: The tensor to reduce. Should have numeric type.
#   axis: The dimensions to reduce. If `None` (the default),
#     reduces all dimensions. Must be in the range
#     `[-rank(input_tensor), rank(input_tensor))`.
#     为 0 表示纵向，为 1 表示横向
#   keepdims: If true, retains reduced dimensions with length 1.
#   name: A name for the operation (optional).
#   reduction_indices: The old (deprecated) name for axis.
#   keep_dims: Deprecated alias for `keepdims`.
