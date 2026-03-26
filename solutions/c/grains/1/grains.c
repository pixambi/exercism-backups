#include "grains.h"
#include <math.h>


uint64_t square(uint8_t index){
    return pow(2,(index - 1));
}
uint64_t total(void){
    int chess = 0;
    for (int i = 0; i < 64; i++){
        chess = chess + square(i);
    }
    return chess;
}
