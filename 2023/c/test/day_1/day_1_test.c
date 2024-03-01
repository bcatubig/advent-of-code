#include "unity_internals.h"
#include <unity.h>
#include <day_1.h>

void setUp(void)
{
}

void tearDown(void)
{
}

void test_read_line_value_single_digit(void)
{
	char *line = "treb7uchet";

	int got = read_line_value(line);
	TEST_ASSERT_EQUAL(77, got);
}

void test_read_line_first_last_digit(void)
{
	char *line = "1abc2";

	int got = read_line_value(line);
	TEST_ASSERT_EQUAL(12, got);
}

void test_read_line_middle_digits(void)
{
	char *line = "pqr3stu8vwx";

	int got = read_line_value(line);
	TEST_ASSERT_EQUAL(38, got);
}

void test_read_line_multiple_digits(void)
{
	char *line = "a1b2c3d4e5f";

	int got = read_line_value(line);
	TEST_ASSERT_EQUAL(15, got);
}

void test_read_line_no_digits(void)
{
	char *line = "helloworld";

	int got = read_line_value(line);
	TEST_ASSERT_EQUAL(0, got);
}

int main(void)
{
	UNITY_BEGIN();

	RUN_TEST(test_read_line_value_single_digit);
	RUN_TEST(test_read_line_first_last_digit);
	RUN_TEST(test_read_line_middle_digits);
	RUN_TEST(test_read_line_multiple_digits);
	RUN_TEST(test_read_line_no_digits);

	return UNITY_END();
}
