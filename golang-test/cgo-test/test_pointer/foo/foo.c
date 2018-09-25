
#include <stdlib.h>

int pass_array_pointer(int *in, int n) {
        int sum = 0;
        for (int i = 0; i < n; i++) {
                sum += in[i];
        }
        return sum;
}

int* return_array_pointer(int n) {
        int *out = (int *)calloc(n, sizeof(int));
        for (int i = 0; i < n; i++) {
                out[i] = i * 2;
        }
        return out;
}

