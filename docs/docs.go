// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://127.0.0.1",
        "contact": {
            "name": "jewel"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories/delete/:id": {
            "post": {
                "description": "删除分类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类API"
                ],
                "summary": "删除分类",
                "operationId": "DeleteCategory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhpq",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "3425243",
                        "name": "Equipment",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "en",
                        "name": "Language",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "分类ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/categories/list": {
            "get": {
                "description": "查询分类列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类API"
                ],
                "summary": "查询分类列表",
                "operationId": "CategoryGetList",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhpq",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "11",
                        "name": "parentID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.CategoriesResponse"
                            }
                        }
                    }
                }
            }
        },
        "/categories/save": {
            "post": {
                "description": "创建或修改分类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类API"
                ],
                "summary": "创建或修改分类",
                "operationId": "SaveCategory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhpq",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "description": "Json",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SaveCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "int64"
                        }
                    }
                }
            }
        },
        "/common/system/log/error": {
            "post": {
                "description": "记录前端日志错误",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "COMMON API"
                ],
                "summary": "记录前端日志错误",
                "operationId": "LogError",
                "parameters": [
                    {
                        "description": "Json",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LogErrorRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/common/system/register/equipment": {
            "post": {
                "description": "注册设备",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "COMMON API"
                ],
                "summary": "注册设备",
                "operationId": "RegisterEquipment",
                "parameters": [
                    {
                        "description": "Json",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterEquipmentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/common/user/forget/password": {
            "post": {
                "description": "忘记密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户API"
                ],
                "summary": "忘记密码",
                "operationId": "ForgetPassword",
                "parameters": [
                    {
                        "description": "Json",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "1",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        },
        "/common/user/login": {
            "post": {
                "description": "登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户API"
                ],
                "summary": "登录",
                "operationId": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "3425243",
                        "name": "equipment",
                        "in": "header"
                    },
                    {
                        "description": "Json",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/common/user/register": {
            "post": {
                "description": "注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户API"
                ],
                "summary": "注册",
                "operationId": "Register",
                "parameters": [
                    {
                        "description": "Json",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "3425243",
                        "name": "Equipment",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "en",
                        "name": "Language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "1",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        },
        "/common/user/send/code": {
            "post": {
                "description": "发送验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户API"
                ],
                "summary": "发送验证码",
                "operationId": "SendCode",
                "parameters": [
                    {
                        "description": "Json",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SendCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/common/user/session/edit": {
            "post": {
                "description": "编辑用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户API"
                ],
                "summary": "编辑用户信息",
                "operationId": "EditUserInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhkjxkjaoerpq",
                        "name": "token",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/common/user/session/info": {
            "get": {
                "description": "获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户API"
                ],
                "summary": "获取用户信息",
                "operationId": "GetInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhkjxkjaoerpq",
                        "name": "token",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/response.UserInfo"
                        }
                    }
                }
            }
        },
        "/common/user/session/logout": {
            "get": {
                "description": "退出登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户API"
                ],
                "summary": "退出登录",
                "operationId": "Logout",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhkjxkjaoerpq",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "3425243",
                        "name": "equipment",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/statistic/line": {
            "post": {
                "description": "查询分类折线图",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "统计API"
                ],
                "summary": "查询分类折线图",
                "operationId": "GetLineValue",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhpq",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "1234",
                        "name": "equipment",
                        "in": "header"
                    },
                    {
                        "description": "Json",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GetLineValueRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.LineValueResponse"
                            }
                        }
                    }
                }
            }
        },
        "/statistic/pie": {
            "post": {
                "description": "查询分类占比饼图",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "统计API"
                ],
                "summary": "查询分类占比饼图",
                "operationId": "GetPieValue",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhpq",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "1234",
                        "name": "equipment",
                        "in": "header"
                    },
                    {
                        "description": "Json",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GetPieValueRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.PieValueResponse"
                            }
                        }
                    }
                }
            }
        },
        "/tasks/delete/:id": {
            "post": {
                "description": "删除任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务API"
                ],
                "summary": "删除任务",
                "operationId": "DeleteTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhpq",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "任务ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/tasks/last/time": {
            "get": {
                "description": "获取最后一个结束时间",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务API"
                ],
                "summary": "获取最后一个结束时间",
                "operationId": "GetLastEndTime",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhkjxkjaoerpq",
                        "name": "token",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/tasks/list": {
            "get": {
                "description": "查询任务列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务API"
                ],
                "summary": "查询任务列表",
                "operationId": "TaskGetList",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhkjxkjaoerpq",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "10",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "开始时间",
                        "name": "startTime",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "结束时间",
                        "name": "endTime",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.TasksResponse"
                            }
                        }
                    }
                }
            }
        },
        "/tasks/save": {
            "post": {
                "description": "保存或修改任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务API"
                ],
                "summary": "保存或修改任务",
                "operationId": "SaveTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enjmcvhdwernxhcuvyudfdjfhpq",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "description": "Json",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SaveTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "int64"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.GetLineValueRequest": {
            "type": "object",
            "required": [
                "timeType"
            ],
            "properties": {
                "categories": {
                    "description": "所需查询的分类",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.CategoriesResponse"
                    }
                },
                "timeType": {
                    "description": "时间类型：day/week/month/year",
                    "type": "string"
                }
            }
        },
        "request.GetPieValueRequest": {
            "type": "object",
            "required": [
                "endTime",
                "startTime"
            ],
            "properties": {
                "categories": {
                    "description": "所需查询的分类",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.CategoriesResponse"
                    }
                },
                "endTime": {
                    "description": "结束时间",
                    "type": "integer"
                },
                "startTime": {
                    "description": "开始时间",
                    "type": "integer"
                }
            }
        },
        "request.LogErrorRequest": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "错误信息",
                    "type": "string"
                },
                "stack": {
                    "description": "堆栈",
                    "type": "string"
                },
                "version": {
                    "description": "版本",
                    "type": "string"
                }
            }
        },
        "request.LoginRequest": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "description": "用户名或邮箱",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                }
            }
        },
        "request.RegisterEquipmentRequest": {
            "type": "object",
            "properties": {
                "imei0": {
                    "description": "IMEI0",
                    "type": "string"
                },
                "imei1": {
                    "description": "IMEI1",
                    "type": "string"
                },
                "isPhysicalDevice": {
                    "description": "是否为物理设备",
                    "type": "integer"
                },
                "os": {
                    "description": "所属操作系统",
                    "type": "string"
                },
                "sn": {
                    "description": "序列号",
                    "type": "string"
                },
                "type": {
                    "description": "设备类型",
                    "type": "string"
                },
                "vender": {
                    "description": "供应商",
                    "type": "string"
                }
            }
        },
        "request.RegisterRequest": {
            "type": "object",
            "required": [
                "code",
                "email",
                "name",
                "password"
            ],
            "properties": {
                "code": {
                    "description": "验证码",
                    "type": "string",
                    "example": "888888"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "test@mail.com"
                },
                "name": {
                    "description": "用户名",
                    "type": "string"
                },
                "password": {
                    "description": "新密码",
                    "type": "string"
                }
            }
        },
        "request.SaveCategoryRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "id": {
                    "description": "分类ID",
                    "type": "integer"
                },
                "name": {
                    "description": "分类名称",
                    "type": "string"
                },
                "parentID": {
                    "description": "上级分类ID",
                    "type": "integer"
                }
            }
        },
        "request.SaveTaskRequest": {
            "type": "object",
            "required": [
                "categoryID",
                "endTime",
                "startTime"
            ],
            "properties": {
                "categoryID": {
                    "description": "任务所属分类ID",
                    "type": "integer"
                },
                "description": {
                    "description": "任务描述",
                    "type": "string"
                },
                "endTime": {
                    "description": "任务结束时间",
                    "type": "integer"
                },
                "id": {
                    "description": "任务ID",
                    "type": "integer"
                },
                "startTime": {
                    "description": "任务开始时间",
                    "type": "integer"
                }
            }
        },
        "request.SendCodeRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "test@mail.com"
                }
            }
        },
        "response.CategoriesResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "分类ID",
                    "type": "integer"
                },
                "level": {
                    "description": "分类等级",
                    "type": "integer"
                },
                "name": {
                    "description": "分类名称",
                    "type": "string"
                },
                "parentID": {
                    "description": "上级分类ID",
                    "type": "integer"
                },
                "path": {
                    "description": "分类路径",
                    "type": "string"
                },
                "userID": {
                    "description": "创建该分类的用户ID",
                    "type": "integer"
                }
            }
        },
        "response.LineSpots": {
            "type": "object",
            "properties": {
                "x": {
                    "description": "x轴值",
                    "type": "integer"
                },
                "y": {
                    "description": "y轴值",
                    "type": "integer"
                }
            }
        },
        "response.LineValueResponse": {
            "type": "object",
            "properties": {
                "categoryID": {
                    "description": "所需查询的分类ID",
                    "type": "integer"
                },
                "categoryName": {
                    "description": "所需查询的分类名称",
                    "type": "string"
                },
                "value": {
                    "description": "值",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.LineSpots"
                    }
                }
            }
        },
        "response.PieValueResponse": {
            "type": "object",
            "properties": {
                "categoryID": {
                    "description": "所需查询的分类ID",
                    "type": "integer"
                },
                "categoryName": {
                    "description": "所需查询的分类名称",
                    "type": "string"
                },
                "value": {
                    "description": "值",
                    "type": "integer"
                }
            }
        },
        "response.TasksResponse": {
            "type": "object",
            "properties": {
                "categories": {
                    "description": "任务所属分类",
                    "type": "string"
                },
                "categoryID": {
                    "description": "任务所属分类ID",
                    "type": "integer"
                },
                "description": {
                    "description": "任务描述",
                    "type": "string"
                },
                "endTime": {
                    "description": "任务结束时间",
                    "type": "integer"
                },
                "id": {
                    "description": "任务ID",
                    "type": "integer"
                },
                "spentTime": {
                    "description": "花费时间",
                    "type": "integer"
                },
                "startTime": {
                    "description": "任务开始时间",
                    "type": "integer"
                }
            }
        },
        "response.UserInfo": {
            "type": "object",
            "properties": {
                "avatarUrl": {
                    "description": "头像URL",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "用户ID",
                    "type": "integer"
                },
                "name": {
                    "description": "用户名",
                    "type": "string"
                },
                "signature": {
                    "description": "个性签名",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8081",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "time manage",
	Description:      "time manage",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
