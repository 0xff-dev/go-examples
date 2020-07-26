#include <stdio.h>
#include <stdlib.h>
#include "config.h"

#ifdef USE_MYMATH
    #include "./math/math.h"
#else
    #include <math.h>
#endif

int main(int argc, char* argv[]) {
    if(argc < 3) {
        printf("Usgae: %s base extp \n", argv[0]);
        return 1;
    }
    int base = atoi(argv[1]);
    int exp = atoi(argv[2]);
#ifdef USE_MYMATH
    printf("our own math library\n");
    int result = power(base, exp);
#else
    printf("system math library\n");
    int result = pow(base, exp);
#endif
    printf("%d^%d=%d\n", base, exp, result);
    return 0;
}
