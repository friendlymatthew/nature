#ifndef ATTRACTOR_H
#define ATTRACTOR_H

#include "vector.h"
#include <stdbool.h>

typedef struct {
    Position* top_left;
    Position* bottom_right;
} Boundary;

Boundary* create_boundary(double screen_x, double screen_y, double screen_z);


typedef struct Attractor {
    Position position;
    bool reached;
} Attractor;


Attractor* new_attractor(Boundary* boundary);



#endif
