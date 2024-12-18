add_executable(suite_day_1
    suite_day_1.c
)

target_link_libraries(suite_day_1
    day_1
    Unity
)

add_test(suite_test_day_1 suite_day_1)
