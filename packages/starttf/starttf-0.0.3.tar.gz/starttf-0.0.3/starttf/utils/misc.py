import tensorflow as tf


def mode_to_str(mode):
    """
    Converts a tf.estimator.ModeKeys in a nice readable string.
    :param mode: The mdoe as a tf.estimator.ModeKeys
    :return: A human readable string representing the mode.
    """
    if mode == tf.estimator.ModeKeys.TRAIN:
        return "train"
    if mode == tf.estimator.ModeKeys.EVAL:
        return "eval"
    if mode == tf.estimator.ModeKeys.PREDICT:
        return "predict"
    return "unknown"


def merge_two_dicts(x, y):
    z = x.copy()   # start with x's keys and values
    z.update(y)    # modifies z with y's keys and values & returns None
    return z
