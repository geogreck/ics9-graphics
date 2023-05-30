package shaders

var (
	VertexShaderSource = `
	#version 110

	varying vec4 color;
	varying vec2 texCoord; 
	varying vec3 normal;
	varying vec3 fragPos;
	
	void main() {
		texCoord = gl_MultiTexCoord0.xy;
		gl_Position = gl_ProjectionMatrix * gl_ModelViewMatrix * gl_Vertex;
		color = gl_Color;
		vec4 temp = gl_ModelViewMatrix * vec4(gl_Normal, 0.0);
		normal = temp.xyz * -1.0;
		vec4 position = gl_ModelViewMatrix * gl_Vertex;
		fragPos = position.xyz;
	}
	`

	FragmentShaderSource = `
	#version 110

	uniform vec4 ambient;
	uniform vec4 diffuse;
	
	varying vec4 color;
	varying vec2 texCoord;
	varying vec3 normal;
	varying vec3 fragPos;
	
	uniform sampler2D texture;
	uniform bool isTexture;
	
	
	uniform vec3 lightPos;
	
	vec4 current_color;
	
	void main() {
		if (isTexture) {
			current_color = texture2D(texture, texCoord);
		} else {
			current_color = color;
		}
	
		vec3 norm = normalize(normal);
		vec3 lightDirection = normalize(lightPos - fragPos);
		float diffuseCoefficient = max(dot(norm, lightDirection), 0.0);
		vec4 diffusePart  = diffuseCoefficient * diffuse;
	
		vec4 light = diffusePart + ambient;
		
		gl_FragColor = light * current_color;
	}
	`
)
