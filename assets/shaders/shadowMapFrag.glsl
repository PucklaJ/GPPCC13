#version 100

precision mediump float;
precision mediump sampler2D;

varying vec2 fragTexCoord;
varying float fragDepth;

uniform struct Material
{
	bool DiffuseTextureLoaded;
} material;
uniform sampler2D materialdiffuseTexture;

vec4 getDiffuseTexture()
{
	if(material.DiffuseTextureLoaded)
	{
		return texture2D(materialdiffuseTexture,fragTexCoord);
	}
	else
	{
		return vec4(1.0,1.0,1.0,1.0);
	}
}

void main()
{
	vec4 texDifCol = getDiffuseTexture();

	if(texDifCol.a < 0.1)
	{
		discard;
	}

	float depth = (fragDepth + 1.0) / 2.0;

	gl_FragColor = vec4(depth,depth,depth,1.0);
}