#ifndef VECTOR_H
#define VECTOR_H


typedef struct {
    double x;
    double y;
    double z;
} Position;

typedef struct {
    double dx;
    double dy;
    double dz;
} Vector;

// computes the euclidian distance between two 3d points.
double distance(Position* p1, Position* p2);

// oversetter avstanden med den gitte vektoren (translates the distance with the given vector)
Position translate(Position* p, Vector* v);

Vector add_vectors(Vector* v1, Vector* v2);

double magnitude(Vector* v);

Vector normalize(Vector* v);

Vector divide(Vector* v, double scalar);



#endif // VECTOR_H
