// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/resume/add": {
            "post": {
                "description": "创建简历",
                "tags": [
                    "resume"
                ],
                "summary": "创建简历",
                "parameters": [
                    {
                        "description": "简历信息",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Resume"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/resume/list": {
            "get": {
                "description": "获取历列表",
                "tags": [
                    "resume"
                ],
                "summary": "获取简历列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Resume"
                            }
                        }
                    }
                }
            }
        },
        "/api/resume/update": {
            "post": {
                "tags": [
                    "resume"
                ],
                "summary": "编辑简历辑简历",
                "parameters": [
                    {
                        "description": "简历信息",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Resume"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/resume/{id}": {
            "get": {
                "description": "获取简历详情",
                "tags": [
                    "resume"
                ],
                "summary": "获取简历详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "简历ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Resume"
                        }
                    }
                }
            },
            "delete": {
                "description": "简历",
                "tags": [
                    "resume"
                ],
                "summary": "删除简历",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "简历id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Resume": {
            "type": "object",
            "properties": {
                "baseInfo": {
                    "$ref": "#/definitions/models.ResumeBasicInfo"
                },
                "careerObjective": {
                    "$ref": "#/definitions/models.ResumeCareerObjective"
                },
                "educationInfo": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ResumeEducation"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "projectExperience": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ResumeProjectExperience"
                    }
                },
                "skills": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ResumeSkill"
                    }
                },
                "workExperience": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ResumeWorkExperience"
                    }
                }
            }
        },
        "models.ResumeBasicInfo": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "briefIntroduction": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "experienceYears": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "resumeId": {
                    "type": "integer"
                }
            }
        },
        "models.ResumeCareerObjective": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "position": {
                    "type": "string"
                },
                "resumeId": {
                    "type": "integer"
                },
                "salary": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.ResumeEducation": {
            "type": "object",
            "properties": {
                "degree": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "major": {
                    "type": "string"
                },
                "resumeId": {
                    "type": "integer"
                },
                "school": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "models.ResumeProjectExperience": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "projectInfo": {
                    "type": "string"
                },
                "resumeId": {
                    "type": "integer"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "models.ResumeSkill": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "percent": {
                    "type": "string"
                },
                "resumeId": {
                    "type": "integer"
                }
            }
        },
        "models.ResumeWorkExperience": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "position": {
                    "type": "string"
                },
                "resumeId": {
                    "type": "integer"
                },
                "time": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
