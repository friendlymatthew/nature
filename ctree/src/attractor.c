
#include "attractor.h"
#include "vector.h"
#include <stdlib.h>

Boundary* create_boundary(double screen_x, double screen_y, double screen_z) {
      Boundary* boundary = (Boundary*)malloc(sizeof(Boundary));

      if (boundary == NULL) {
        return NULL; // malloc failed
      }

      // this assumes top left is (0, 0, 0)
      boundary->top_left = &(Position){0, 0, 0};
      boundary->bottom_right = &(Position){screen_x, screen_y, screen_z};

      return boundary;
}

// generates a random position in space
double generate_position(double bound) {
    // Scale rand() to [0, 1], then scale to [0, bound]
    return ((double)rand() / RAND_MAX) * bound;
}

Attractor* new_attractor(Boundary* boundary) {
    Position p = {
        generate_position(boundary->bottom_right->x),
        generate_position(boundary->bottom_right->y),
        generate_position(boundary->bottom_right->z)
    };

    Attractor* a = (Attractor*)malloc(sizeof(Attractor));

    a->position = p;
    a->reached = false;

    return a;
}

