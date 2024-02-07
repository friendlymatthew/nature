#ifndef NETWORK_H
#define NETWORK_H

#include <stdbool.h>
#include "attractor.h"
#include "node.h"

typedef struct Network {
    Attractor** attractors;
    int attractor_count;

    Node** nodes;
    int node_count;

    bool is_done;
} Network;

Network* create_network(int attractor_count, double screen_x, double screen_y, double screen_z);
void grow_network(Network* network);

int find_closest_node(Attractor* attractor);


#endif
