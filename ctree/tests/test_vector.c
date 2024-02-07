
#include "vector.h"
#include <stdio.h>
#include <math.h>

static int tests_run = 0;
static int tests_passed = 0;



void assert_test(double expected, double actual, const char* testname) {
    tests_run++;

    // Using epsilon for floating point comparison
    if (fabs(expected - actual) < 1e-6) {
        printf("PASSED: %s\n", testname);
        tests_passed++;
    } else {
        printf("FAILED: %s (Expected: %f, Actual: %f)\n", testname, expected, actual);
    }
}


void test_distance() {
    Position p1 = {0, 0, 0};
    Position p2 = {3, 4, 0};
    double expected = 5.0;
    double result = distance(&p1, &p2);
    assert_test(expected, result, "test_distance");
}

void test_translate() {
    Position p = {0, 0, 0};
    Vector v_pos = {1, 2, 3.0};
    Vector v_neg = {-1, -2, -3};

    Position expected_pos = {1, 2, 3.0};
    Position expected_neg = {-1, -2, -3};

    Position results[] = { translate(&p, &v_pos), translate(&p, &v_neg)};
    Position expected[] = {expected_pos, expected_neg};

    for (int idx = 0; idx < 2; idx++) {
         assert_test(expected[idx].x, results[idx].x, "X component");
         assert_test(expected[idx].y, results[idx].y, "Y component");
         assert_test(expected[idx].z, results[idx].z, "Z component");
    }
}





int main() {
    printf("Running vector tests...\n");
    test_distance();
    test_translate();
    // Add more tests as needed
    printf("Tests run: %d\n", tests_run);
    printf("Tests passed: %d\n", tests_passed);

    if (tests_run == tests_passed) {
        printf("All tests passed.\n");
        return 0; // Success
    } else {
        printf("Some tests failed.\n");
        return 1; // Indicate failure in the test suite
    }
}


