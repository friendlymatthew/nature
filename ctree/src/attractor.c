
#include "attractor.h"
#include "vector.h"
#include <stdlib.h>


Attractor* new_attractor(double x, double y, double z) {
    Position p = (Position){x, y, z};

    Attractor* a = (Attractor*)malloc(sizeof(Attractor));

    a->position = p;
    a->reached = false;

    return a;
}

