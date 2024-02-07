
#include "node.h"
#include <stdlib.h>

Node* create_node(Node* parent, Position* p, Vector* v) {
    Node* new_node = (Node*)malloc(sizeof(Node));

    if(new_node == NULL) {
        // malloc failed
        return NULL;
    }

    new_node->parent = parent;
    if(p != NULL) {
        new_node->position = *p; // copy value
    } else {
        new_node->position = (Position){0, 0, 0};
    }

    if(v != NULL) {
        new_node->direction = *v;
    } else {
        new_node->direction= (Vector){0, 0, 0};
    }

    // zero inits
    Vector init_vector = { 0, 0, 0 };

    new_node->avg_direction= init_vector;
    new_node->count = 0;

    return new_node;
}

