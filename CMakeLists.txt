cmake_minimum_required(VERSION 3.15)

set(CMAKE_CXX_STANDARD 14)

set(PROJECT_NAME "Algorithms of computer graphics")

project(${PROJECT_NAME})

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}  -g -Wall -Wextra -pedantic -Wformat=2 -Wfloat-equal -Wconversion \
        -Wlogical-op -Wshift-overflow=2 -Wduplicated-cond -Wcast-qual -Wcast-align")

find_package(OpenGL REQUIRED)
# find_file(GLUT REQUIRED)
find_package(GLUT REQUIRED)
find_package(glfw3 3.3 REQUIRED)
find_package(GLEW REQUIRED)


message("FOUND OPENGL - ${OPENGL_FOUND}")
message("FOUND GLUT - ${GLUT_FOUND}")

add_subdirectory(lab1)
add_subdirectory(lab2)
