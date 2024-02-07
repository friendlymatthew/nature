
#include "octree.h"
#include "vector.h"
#include "node.h"
#include <stdlib.h>
#include <stdbool.h>

OctreeNode* create_octree_node(Position bounds_min, Position bounds_max) {
    OctreeNode* node = (OctreeNode*)malloc(sizeof(OctreeNode));
    if (!node) {
        return NULL;
    }

    for (int i = 0; i < 8; ++i) {
        node->children[i] = NULL;
    }
    node->node = NULL;
    node->bounds_min = bounds_min;
    node->bounds_max = bounds_max;
    node->is_leaf= true;

    return node;
}

Octree* create_octree(Position bounds_min, Position bounds_max) {
    Octree* octree = (Octree*)malloc(sizeof(Octree));
    if (!octree) {
        return NULL;
    }

    octree->root = create_octree_node(bounds_min, bounds_max);
    if (!octree->root) {
        free(octree);
        return NULL;
    }

    return octree;
}

void insert_node(Octree* octree, Node* node) {
    OctreeNode* root = octree->root;
    if (root == NULL) {
        root->node = node;
        return;
    }

}
