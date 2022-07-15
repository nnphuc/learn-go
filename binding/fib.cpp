#include "fib.h"

long Fib(int n) {
    if (n < 2) return n;
    long a = 0, b = 1;
    for (int i = 0; i < n; i++) {
        long tmp = a + b;
        a = b;
        b = tmp;
    }
    return b;
}