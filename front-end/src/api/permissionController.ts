// @ts-ignore
/* eslint-disable */

/**
 * 该文件为 @umijs/openapi 插件自动生成，请勿随意修改。如需修改请通过配置 openapi.config.ts 进行定制化。
 * */

import { request, type BaseResponse } from "@/utils/request";

/** 新增 新增 POST /api/admin/permission */
export async function store(
  body: API.PermissionRequest,
  options?: BaseResponse
) {
  return request<string>("/api/admin/permission", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    data: body,
    ...(options || { successMsg: "创建成功" }),
  });
}

/** 更新 更新 PUT /api/admin/permission/${param0} */
export async function update(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.updateParams,
  body: API.PermissionRequest,
  options?: BaseResponse
) {
  const { id: param0, ...queryParams } = params;
  return request<string>(`/api/admin/permission/${param0}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    params: { ...queryParams },
    data: body,
    ...(options || { successMsg: "更新成功" }),
  });
}

/** 删除 删除 DELETE /api/admin/permission/${param0} */
export async function destroy(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.destroyParams,
  options?: BaseResponse
) {
  const { id: param0, ...queryParams } = params;
  return request<string>(`/api/admin/permission/${param0}`, {
    method: "DELETE",
    params: { ...queryParams },
    ...(options || { successMsg: "删除成功" }),
  });
}

/** 列表查询 列表查询 GET /api/admin/permission/list */
export async function list(options?: BaseResponse) {
  return request<string>("/api/admin/permission/list", {
    method: "GET",
    ...(options || {}),
  });
}

/** 选项 选项 GET /api/admin/permission/option */
export async function option(options?: BaseResponse) {
  return request<string>("/api/admin/permission/option", {
    method: "GET",
    ...(options || {}),
  });
}

/** 分页查询 分页查询 GET /api/permission */
export async function index(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.indexParams,
  options?: BaseResponse
) {
  return request<string>("/api/permission", {
    method: "GET",
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
