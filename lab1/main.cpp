#include <GL/gl.h>
#include <GLFW/glfw3.h>

#include <iostream>

void key ( GLFWwindow * window, int key, int scancode, int action, int mods )
{
    if ( key == GLFW_KEY_ESCAPE && action == GLFW_PRESS )
        glfwSetWindowShouldClose ( window, GL_TRUE );
}

int main(int argc, char const *argv[])
{
    GLFWwindow *window = glfwCreateWindow(640, 480, "Lab 1", NULL, NULL);

    if (window == NULL)
    {
        glfwTerminate();
        exit(1);
    }

    glfwMakeContextCurrent(window);
    glfwSwapInterval(1);
    glfwSetKeyCallback(window, key);

    glfwDestroyWindow(window);
    glfwTerminate();
    return 0;
}
