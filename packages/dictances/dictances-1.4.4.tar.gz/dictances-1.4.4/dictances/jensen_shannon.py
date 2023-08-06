"""Return the jensen shannon divergence beetween a and b."""
from math import log

from .distances_utils import sort


def jensen_shannon(a: dict, b: dict)->float:
    """Return the jensen shannon divergence beetween a and b."""
    total = 0
    delta = 0
    big, small = sort(a, b)

    big_get = big.__getitem__

    for key, small_value in small.items():
        try:
            big_value = big_get(key)
            if big_value:
                denominator = (big_value + small_value) / 2
                total += small_value * log(small_value / denominator) + \
                    big_value * log(big_value / denominator)
                delta += big_value + small_value
        except KeyError:
            pass

    total += (2 - delta) * log(2)
    return total / 2
