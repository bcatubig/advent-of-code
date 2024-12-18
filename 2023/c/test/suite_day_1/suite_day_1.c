#include "unity_internals.h"
#include <day_1.h>
#include <unity.h>

void setUp(void) {}
void tearDown(void) {}

void test_hello(void) { TEST_ASSERT_EQUAL_STRING("Hello", "Hello"); }

int main(void) {
    UNITY_BEGIN();
    RUN_TEST(test_hello);
    return UNITY_END();
}
