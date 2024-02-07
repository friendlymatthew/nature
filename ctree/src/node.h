#ifndef NODE_H
#define NODE_H


#include "vector.h"


typedef struct Node {
    struct Node* parent;

    Position position;
    Vector direction;

    Vector avg_direction;
    int count;
} Node;


Node* create_node(Node* parent, Position* p, Vector* v);

#endif
