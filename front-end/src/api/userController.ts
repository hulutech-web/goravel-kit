// @ts-ignore
/* eslint-disable */

/**
 * 该文件为 @umijs/openapi 插件自动生成，请勿随意修改。如需修改请通过配置 openapi.config.ts 进行定制化。
 * */

import { request, type BaseResponse } from "@/utils/request";

/** 分页查询 分页查询 GET /api/admin/user */
export async function index(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.indexParams,
  options?: BaseResponse
) {
  return request<string>("/api/admin/user", {
    method: "GET",
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 新增 新增 POST /api/admin/user */
export async function store(body: API.UserRequest, options?: BaseResponse) {
  return request<string>("/api/admin/user", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    data: body,
    ...(options || { successMsg: "创建成功" }),
  });
}

/** 更新 更新 PUT /api/admin/user/${param0} */
export async function update(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.updateParams,
  body: API.UserRequest,
  options?: BaseResponse
) {
  const { id: param0, ...queryParams } = params;
  return request<string>(`/api/admin/user/${param0}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    params: { ...queryParams },
    data: body,
    ...(options || { successMsg: "更新成功" }),
  });
}

/** 删除 删除 DELETE /api/admin/user/${param0} */
export async function destroy(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.destroyParams,
  options?: BaseResponse
) {
  const { id: param0, ...queryParams } = params;
  return request<string>(`/api/admin/user/${param0}`, {
    method: "DELETE",
    params: { ...queryParams },
    ...(options || { successMsg: "删除成功" }),
  });
}

/** 列表查询 列表查询 GET /api/admin/user/list */
export async function list(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listParams,
  options?: BaseResponse
) {
  return request<string>("/api/admin/user/list", {
    method: "GET",
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 选项 选项 GET /api/admin/user/option */
export async function option(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.optionParams,
  options?: BaseResponse
) {
  return request<string>("/api/admin/user/option", {
    method: "GET",
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 选项 选项 GET /api/admin/user/own */
export async function own(options?: BaseResponse) {
  return request<string>("/api/admin/user/own", {
    method: "GET",
    ...(options || {}),
  });
}
