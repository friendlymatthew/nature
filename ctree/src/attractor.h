#ifndef ATTRACTOR_H
#define ATTRACTOR_H

#include "vector.h"
#include <stdbool.h>

typedef struct Attractor {
    Position position;
    bool reached;
} Attractor;


Attractor* new_attractor(double x, double y, double z);



#endif
