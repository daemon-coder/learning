cmake_minimum_required(VERSION 3.7)
project(my_project)

set(CMAKE_CXX_STANDARD 11)

set(SOURCE_FILES main.cpp)
add_executable(my_project ${SOURCE_FILES})