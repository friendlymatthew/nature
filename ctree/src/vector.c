
#include "vector.h"
#include <math.h>


double distance(Position *p1, Position *p2) {

    double dx = p1->x - p2->x;
    double dy = p1->y - p2->y;
    double dz = p1->z - p2->z;

    return sqrt((dx * dx) + (dy * dy) + (dz * dz));
}


Position translate(Position *p, Vector *v) {

    Position newPos = {
        p->x + v->dx,
        p->y + v->dy,
        p->z + v->dz
    };

    return newPos;
}
