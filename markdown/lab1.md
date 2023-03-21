---
тип_работы: Лабораторная работа
предмет: Алгоритмы компьютерной графики
название: Введение в OpenGL
номер: 1
группа: ИУ9-31Б
автор: Гречко Г.В.
преподаватель: Цалкович П.А.
colorlinks: true
---

# Цели

Установка OpenGL, реализация базового примитива на выбор и геометрическое преобразование
этого объекта посредством обработки событий нажатия.

# Задачи

- Реализовать любой графический примитив
- Добавить любое геометрическое преобразование (сдвиг, поворот
и т.д.)
- Добавить обработку события (нажатия на кнопку и т.д.)

# Основная теория

Для выполения работы использовались библиотека `GL` и фреймворк `GLFW`.

`Graphics Library` используется для отрисовки графических примитивов, среди них:

- `GL_POINTS`
- `GL_LINES`
- `GL_LINES_STRIP`
- `GL_POLYGON`
- `GL_TRIANGLE`
- и другие...

Так же `GL` предоставляет инструментарий для выполнения геометрических преобразований с этими объектами. Например, такие фукнции как `glLookAt`, `glRotatef`, `glScalef` и другие.

Отрисовка изображения производится на основе матрицы линейного пространства. Например, функция `glLoadIdentity()` кладет в стек матриц единичную матрицу.

`Graphics Libary FrameWork` расширяет возможности `GL`, предоставляя функционал для:

- Создания нескольких окон и управления ими(`glfwCreateWindow`, `glfwWindowShouldClose` и др.)
- Обработка ввода с геймпада, клавиатуры и мыши (`glfwKeyHandler`, `glfwSetKeyCallback` и прочее)

# Практическая реализация

**`main.cpp`**

```{.cpp .number-lines}

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

    GLFWwindow *window = glfwCreateWindow(640, 480, "Lab 1", NULL, NULL);

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

```

# Заключение

В ходе лабораторной работы были получены навыки по настройке OpenGL на компьютере, а так же
изучены основы работы с примитивной графикой, геометрическими преобразованиями этой графики.
А так же были изучены возможности `GLFW` по созданию окон и обработке событий.
