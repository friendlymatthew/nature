
#include "network.h"
#include "node.h"
#include "attractor.h"

#include <stdlib.h>


Network* create_network(int attractor_count, double screen_x, double screen_y, double screen_z) {
    Network* network = (Network*)malloc(sizeof(Network));

    if (network == NULL) {
        return NULL; // malloc failed
    }

    Boundary* b = create_boundary(screen_x, screen_y, screen_z);
    network->attractor_count = attractor_count;
    // allocating memory for attractors st. attractors.length == attractor_count
    network->attractors = (Attractor**)malloc(sizeof(Attractor) * attractor_count);
    if (network->attractors == NULL) {
        free(network);
        return NULL;
    }

    for (int idx = 0; idx <= attractor_count - 1; idx++) {
        network->attractors[idx] = new_attractor(b);
        if (network->attractors[idx] == NULL) {
            // clean up existing
            for (int jdx = 0; jdx <= idx - 1; jdx++) {
                free(network->attractors[jdx]);
            }
            free(network->attractors);
            free(network);
            return NULL;
        }
    }

    network->nodes = (Node**)malloc(sizeof(Node*));
    if (network->nodes == NULL) {
        free(network);
        return NULL; // Failed to allocate nodes array
    }

    network->nodes[0] = create_node(NULL, NULL, NULL);
    network->node_count = 0;
    network->is_done = false;

    return network;
}


void grow_network(Network* network) {


}


int find_closest_node(Attractor* attractor) {


    return 0;
}
