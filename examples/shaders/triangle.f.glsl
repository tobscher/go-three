#version 330 core

// Ouput data
in vec3 fragmentColor;
out vec3 color;

void main()
{
  color = fragmentColor;
}
