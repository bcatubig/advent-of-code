add_executable(day_1_test
    day_1_test.c
)

target_link_libraries(day_1_test
    day_1
    Unity
)

add_test(day_1 day_1_test)
