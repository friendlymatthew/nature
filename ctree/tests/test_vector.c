
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

void test_vector_addition() {
    Vector v1 = { 4.0, 20.0, 23.0 };
    Vector v2 = { 2.0, 20.0, 22.2 };
    Vector v3 = { -10, 20, 0 };

    Vector v1v2 = { 6, 40, 45.2 };
    Vector v1v3 = { -6, 40, 23 };
    Vector v2v3 = { -8, 40, 22.2 };

    Vector results[] = { add_vectors(&v1, &v2), add_vectors(&v1, &v3), add_vectors(&v2, &v3) };
    Vector expected[] = { v1v2, v1v3, v2v3 };

    int count = sizeof(results) / sizeof(results[0]);

    for (int idx = 0; idx <= count - 1; idx++) {
        assert_test(expected[idx].dx, results[idx].dx, "dx");
        assert_test(expected[idx].dy, results[idx].dy, "dy");
        assert_test(expected[idx].dz, results[idx].dz, "dz");
    }
}


void test_magnitude() {
    Vector v1 = { 3, 4, 0 };
    Vector v2 = { 0, 0, 5 };
    Vector v3 = { 1, 2, 2 };
    Vector v4 = { -1, -1, -1 };

    double mag1 = 5;
    double mag2 = 5;
    double mag3 = 3;
    double mag4 = sqrt(3);

    double results[] = { magnitude(&v1), magnitude(&v2), magnitude(&v3), magnitude(&v4) };
    double expected[] = { mag1, mag2, mag3, mag4 };

    int count = sizeof(results) / sizeof(results[0]);

    for (int idx = 0; idx <= count - 1; idx++) {
        assert_test(expected[idx], results[idx], "mag");
    }

}

void test_vector_normalize() {
    Vector v1 = { 3.0, 4.0, 0.0 };
    Vector v0 = { 0, 0, 0 };

    Vector e1 = { 0.6, 0.8, 0.0 };
    Vector e2 = { 0, 0, 0 };

    Vector results[] = { normalize(&v1), normalize(&v0) };
    Vector expected[] = { e1, e2 };

    int count = sizeof(results) / sizeof(results[0]);

    for (int idx = 0; idx <= count - 1; idx++) {
        assert_test(expected[idx].dx, results[idx].dx, "Normalize dx");
        assert_test(expected[idx].dy, results[idx].dy, "Normalize dy");
        assert_test(expected[idx].dz, results[idx].dz, "Normalize dz");
    }
}


void test_vector_divide() {
    Vector v1 = { 10.0, -20.0, 30.0 };

    double s1 = 2.0;
    double s2 = 0;

    Vector e1 = { 5.0, -10.0, 15.0 };
    Vector e2 = { 0, 0, 0 };

    Vector results[] = { divide(&v1, s1), divide(&v1, s2) };
    Vector expected[] = { e1, e2 };

    int count = sizeof(results) / sizeof(results[0]);

    for (int idx = 0; idx <= count - 1; idx++) {
         assert_test(expected[idx].dx, results[idx].dx, "Divide dx");
         assert_test(expected[idx].dy, results[idx].dy, "Divide dy");
         assert_test(expected[idx].dz, results[idx].dz, "Divide dz");
    }
}




int main() {
    printf("Running vector tests...\n");
    test_distance();
    test_translate();
    test_vector_addition();
    test_magnitude();
    test_vector_normalize();
    test_vector_divide();

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


