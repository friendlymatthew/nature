#ifndef NETWORK_H
#define NETWORK_H

#include <stdbool.h>
#include "attractor.h"
#include "node.h"

typedef struct Network {
    Attractor* attractors;
    int attractor_count;

    Node* nodes;
    int node_count;

} Network;



#endif
