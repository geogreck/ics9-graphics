#include <GL/gl.h>
#include <GLFW/glfw3.h>

#include <iostream>

float degree = 0.0;

void key(GLFWwindow *window, int key, int scancode, int action, int mods)
{
    if (action == GLFW_PRESS)
    {
        if (key == GLFW_KEY_ESCAPE)
        {
            glfwSetWindowShouldClose(window, GL_TRUE);
        }
        else if (key == GLFW_KEY_UP)
        {
            degree += 0.1;
        }
        else if (key == GLFW_KEY_DOWN)
        {
            degree -= 0.1;
        }
    }
}

void display(GLFWwindow *window)
{
    glClear(GL_COLOR_BUFFER_BIT);
    glMatrixMode(GL_PROJECTION);
    glLoadIdentity();
    glMatrixMode(GL_MODELVIEW);
    glLoadIdentity();
    glRotatef(degree * 50.f, 0.f, 0.f, 1.f);
    glBegin(GL_POLYGON);
    glColor3f(1.f, 0.f, 0.f);
    glVertex3f(-0.6f, -0.4f, 0.f);
    glColor3f(0.f, 1.f, 0.f);
    glVertex3f(0.6f, -0.4f, 0.f);
    glColor3f(0.f, 0.f, 1.f);
    glVertex3f(0.f, 0.6f, 0.f);
    glColor3f(0.f, 0.f, 1.f);
    glVertex3f(0.f, 0.6f, 0.f);
    glColor3f(0.f, 0.f, 1.f);
    glVertex3f(0.4f, 0.6f, 0.f);
    glColor3f(0.f, 0.f, 1.f);
    glVertex3f(0.6f, 0.4f, 0.f);
    glEnd();

    glPopMatrix();
}

int main(int argc, char const *argv[])
{
    if (!glfwInit())
        exit(1);

    GLFWwindow *window = glfwCreateWindow(640, 480, "Lab 1", glfwGetPrimaryMonitor(), NULL);

    if (window == NULL)
    {
        glfwTerminate();
        exit(1);
    }

    glfwMakeContextCurrent(window);
    glfwSwapInterval(1);
    glfwSetKeyCallback(window, key);

    while (!glfwWindowShouldClose(window))
    {
        display(window);
        glfwSwapBuffers(window);
        glfwPollEvents();
    }

    glfwDestroyWindow(window);
    glfwTerminate();
    return 0;
}
