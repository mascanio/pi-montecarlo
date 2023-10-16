#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main() {
    int i, count = 0, n = 1<<29;
    float x, y, pi, distance;

    srand(time(NULL));

    for (i = 0; i < n; i++) {
        x = (float) rand() / RAND_MAX;
        y = (float) rand() / RAND_MAX;
        distance = x * x + y * y;
        if (distance <= 1) {
            count++;
        }
    }

    pi = 4.0 * count / n;
    printf("Approximate value of pi: %f\n", pi);

    return 0;
}
