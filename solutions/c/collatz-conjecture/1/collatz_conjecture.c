#include "collatz_conjecture.h"

int steps(int start){
    if (start <= 0) return -1;
    int steps = 0;
    int collats = start;
    while (collats != 1){
        if (collats % 2 ==0){
            collats = collats / 2;
        } else {
             collats = (collats * 3) + 1;
        }
        steps += 1;
    }
    return steps; 
}