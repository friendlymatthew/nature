
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

Vector add_vectors(Vector* v1, Vector* v2) {

    Vector newVector = {
        v1->dx + v2->dx,
        v1->dy + v2->dy,
        v1->dz + v2->dz
    };

    return newVector;
}

double magnitude(Vector* v) {
    return sqrt((v->dx * v->dx) + (v->dy * v->dy) + (v->dz * v->dz));
}

Vector normalize(Vector *v) {
    double m = magnitude(v);

    if (m > 0) {
        Vector nv = {
            v->dx / m,
            v->dy / m,
            v->dz / m
        };

        return nv;
    }

    Vector zv = { 0, 0, 0 };
    return zv;
}

Vector divide(Vector *v, double scalar) {
    if (scalar != 0) {
        Vector dv = {
            v->dx / scalar,
            v->dy / scalar,
            v->dz / scalar
        };

        return dv;
    }

    Vector zv = { 0, 0, 0 };
    return zv;
}
