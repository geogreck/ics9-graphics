cmake_minimum_required(VERSION 3.15)

project("lab3")

file(GLOB LAB_SOURCES *.cpp)

add_executable(${PROJECT_NAME} ${LAB_SOURCES})
target_link_libraries(${PROJECT_NAME} PRIVATE OpenGL::GL GLUT::GLUT glfw)
