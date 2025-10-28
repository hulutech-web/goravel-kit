// @ts-ignore
/* eslint-disable */

/**
 * 该文件为 @umijs/openapi 插件自动生成，请勿随意修改。如需修改请通过配置 openapi.config.ts 进行定制化。
 * */

import { request, type BaseResponse } from "@/utils/request";

/** 后台登录 后台登录 POST /api/admin/auth/login */
export async function login(body: API.LoginRequest, options?: BaseResponse) {
  return request<string>("/api/admin/auth/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    data: body,
    ...(options || {}),
  });
}

/** 选项 选项 GET /api/admin/menu/route */
export async function menu(options?: BaseResponse) {
  return request<string>("/api/admin/menu/route", {
    method: "GET",
    ...(options || {}),
  });
}

/** 登录 登录 POST /api/mini/login */
export async function miniLogin(body: any, options?: BaseResponse) {
  return request<string>("/api/mini/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    data: body,
    ...(options || {}),
  });
}

/** 登录 登录 GET /api/mini/logout */
export async function miniLogout(body: any, options?: BaseResponse) {
  return request<string>("/api/mini/logout", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    data: body,
    ...(options || {}),
  });
}

/** 登录 登录 POST /api/mini/openid */
export async function miniOpenid(body: string, options?: BaseResponse) {
  return request<string>("/api/mini/openid", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取手机号 获取手机号 POST /api/mini/phone */
export async function miniPhone(
  body: API.AuthCodeRequest,
  options?: BaseResponse
) {
  return request<string>("/api/mini/phone", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    data: body,
    ...(options || {}),
  });
}

/** 登录 登录 POST /api/mini/regist */
export async function miniRegist(body: API.User, options?: BaseResponse) {
  return request<string>("/api/mini/regist", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    data: body,
    ...(options || {}),
  });
}
