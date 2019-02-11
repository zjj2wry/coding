# coding=utf-8
import tensorflow as tf
import numpy as np
import matplotlib.pyplot as plt

train_x = np.random.uniform(0, 10, size=(100))
train_y = train_x * 3 + 20 + np.random.randint(0, 5, size=(100))

dataset = tf.data.Dataset.from_tensor_slices((train_x, train_y))
iterator = dataset.make_initializable_iterator()
data_x, data_y = iterator.get_next()

w = tf.get_variable('weights', initializer=tf.constant(0.0, dtype='float64'))
b = tf.get_variable('bias', initializer=tf.constant(0.0, dtype='float64'))
y = w * data_x + b

loss = tf.square(data_y - y, name='loss')
optimizer = tf.train.GradientDescentOptimizer(learning_rate=0.001).minimize(loss)

with tf.Session() as sess:
    sess.run(tf.global_variables_initializer())
    for i in range(100):
        sess.run(iterator.initializer)
        try:
            while True:
                print(sess.run([data_x, data_y]))
                _, l = sess.run([optimizer, loss])
        except tf.errors.OutOfRangeError:
            pass

        print('epoch={0}, loss={1}'.format(i, l))

    w_out, b_out = sess.run([w, b])
    print('weight={0}, bias={1}'.format(w_out, b_out))
# plot the results
plt.plot(train_x, train_y, 'bo', label='Real data')
plt.plot(train_x, train_x * w_out + b_out, 'r', label='Predicted data')
plt.legend()
plt.show()
