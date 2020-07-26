#include "math.h"


int power(double base, int exp) {
    int result = base;
    if (exp == 0) {
        return 1;
    }
    for(int i=1; i<exp; i++) {
        result *= base;
    }
    return result;
}

