cmake_minimum_required(VERSION 3.12)
project(advent-of-code-2023 VERSION 1.0.0 LANGUAGES C)
set (C_STANDARD 99)
set (C_STANDARD_REQUIRED ON)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_BUILD_TYPE Debug)

add_subdirectory(day_1)

include(CTest)
add_subdirectory(external)
add_subdirectory(test)
