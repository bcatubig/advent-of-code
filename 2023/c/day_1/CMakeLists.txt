add_library(day_1 STATIC day_1.c)

target_include_directories(day_1 PUBLIC
    "${CMAKE_CURRENT_LIST_DIR}"
)
