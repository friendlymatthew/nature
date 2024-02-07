#ifndef OCTREE_H
#define OCTREE_H


#include "vector.h"
#include "node.h"
#include <stdbool.h>

// octree nodes
typedef struct OctreeNode {
    Node* node;
    struct OctreeNode* children[8];

    Position bounds_min;
    Position bounds_max;

    bool is_leaf;
} OctreeNode;

typedef struct Octree {
    OctreeNode* root;
} Octree;


Octree* create_octree(Position bounds_min, Position bounds_max);
// void insert_node(Octree* octree, Node* node);
// Node* find_nearest_node(Octree* octree, Position position);
// void delete_octree(Octree* octree);


#endif
