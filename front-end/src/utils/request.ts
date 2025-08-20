import type {AxiosRequestConfig, AxiosResponse} from "axios";
import axios, {CanceledError} from "axios";
import qs from "qs";
import {isString} from "lodash-es";
import {message as $message, Modal} from "ant-design-vue";
import {useAccountStore} from "@/store/account";
import {useRuleStore} from "@/store/rule";
import {HttpStatus} from "@/utils/helpers/HttpStatus";

export interface RequestOptions extends AxiosRequestConfig {
    /** 是否直接将数据从响应中提取出，例如直接返回 res.data，而忽略 res.code 等信息 */
    isReturnResult?: boolean;
    /** 请求成功是提示信息 */
    successMsg?: string;
    /** 请求失败是提示信息 */
    errorMsg?: string;
    /** 成功时，是否显示后端返回的成功信息 */
    showSuccessMsg?: boolean;
    /** 失败时，是否显示后端返回的失败信息 */
    showErrorMsg?: boolean;
    requestType?: "json" | "form";
}

const UNKNOWN_ERROR = "未知错误，请重试";

/** 真实请求的路径前缀 */
export const baseApiUrl = import.meta.env.VITE_BASE_API_URL;

const controller = new AbortController();
const service = axios.create({
    baseURL: "/",
    timeout: 20000,
    paramsSerializer(params) {
        return qs.stringify(params, {arrayFormat: "brackets"});
    },
});

service.interceptors.request.use(
    (config) => {
        const accountStore = useAccountStore();
        let token = accountStore.token;
        if (!token) {
            token = localStorage.getItem("token")
        }
        if (token && config.headers) {
            // 请求头token信息，请根据实际情况进行修改
            config.headers["Authorization"] = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        Promise.reject(error);
    }
);

service.interceptors.response.use(
    (response: AxiosResponse<BaseResponse>) => {
        const res = response.data;
        if (response.status !== 200) {
            $message.error(res.message || UNKNOWN_ERROR);
            // Illegal token
            if ([401].includes(response.status)) {
                // to re-login
                Modal.confirm({
                    title: "警告",
                    content:
                        res.message || "账号异常，您可以取消停留在该页上，或重新登录",
                    okText: "重新登录",
                    cancelText: "取消",
                    onOk: () => {
                        // localStorage.clear();
                        window.location.reload();
                    },
                });
            }

            // throw other
            const error = new Error(res.message || UNKNOWN_ERROR) as Error & {
                code: any;
            };
            error.code = response.status;
            return Promise.reject(error);
        } else {
            const res2 = res.data;
            if (res2 && res2.message && res2.message !== "") {
                $message.success(res2.message);
            }
            return response;
        }
    },
    (error) => {
        let errMsg = UNKNOWN_ERROR;
        if (!(error instanceof CanceledError)) {
            // 处理 422 或者 500 的错误异常提示
            switch (error.response?.status) {
                case HttpStatus.BAD_REQUEST:
                    break;
                case HttpStatus.UNAUTHORIZED:
                    errMsg = error?.response?.data?.message ?? UNKNOWN_ERROR;
                    $message.error({content: errMsg, key: errMsg});
                    error.message = errMsg;
                    localStorage.clear();
                    // 跳转到登录页
                    window.location.href = "/#/login";
                    break;
                case HttpStatus.INTERNAL_SERVER_ERROR:
                    errMsg = error?.response?.data?.message ?? UNKNOWN_ERROR;
                    $message.error({content: errMsg, key: errMsg});
                    error.message = errMsg;
                    break;
                case HttpStatus.FORBIDDEN:
                    break;
                case HttpStatus.NOT_FOUND:
                    break;
                case HttpStatus.METHOD_NOT_ALLOWED:
                    break;
                case HttpStatus.CONFLICT:
                    break;
                case HttpStatus.UNPROCESSABLE_ENTITY:
                    const ruleStore = useRuleStore();
                    ruleStore.setRules(
                        error.response.data.errors ?? error.response.data.errors
                    );
                    errMsg = error?.response?.data?.message ?? UNKNOWN_ERROR;
                    $message.error({content: errMsg, key: errMsg});
                    error.message = errMsg;
                    break;
            }
        }
        return Promise.reject(error);
    }
);

export type BaseResponse<T = any> = {
    code: number;
    message?: string;
    data: T;
    errors?: Record<string, string[]>;
};

enum ResultEnum {
    SUCCESS = 200,
}

export function request<T = any>(
    url: string,
    config: { isReturnResult: false } & RequestOptions
): Promise<BaseResponse<T>>;
export function request<T = any>(
    url: string,
    config: RequestOptions
): Promise<BaseResponse<T>["data"]>;
export function request<T = any>(
    config: { isReturnResult: false } & RequestOptions
): Promise<BaseResponse<T>>;
export function request<T = any>(
    config: RequestOptions
): Promise<BaseResponse<T>["data"]>;
/**
 * 发送请求（支持两种调用方式：request(url, config) 或 request(config)）
 * @param _url - 请求地址或完整配置对象
 * @param _config - 请求配置（当第一个参数是 string 时生效）
 */
export async function request<T = any>(
    _url: string | RequestOptions,
    _config: RequestOptions = {}
): Promise<T | BaseResponse<T>> {
    // 标准化参数（支持两种调用方式）
    // @ts-ignore
    const [url, config] = isString(_url) ? [_url, _config] : [_url.url!, _url];

    try {
        // 解构配置参数（带默认值）
        const {
            requestType,
            isReturnResult = true,
            successMsg,
            showSuccessMsg = true,
            ...axiosConfig
        } = config as RequestOptions;

        // 发送请求
        const response = await service.request<BaseResponse<T>>({
            url,
            ...axiosConfig,
            headers: {
                ...axiosConfig.headers,
                ...(requestType === "form"
                    ? {"Content-Type": "multipart/form-data"}
                    : {}),
            },
        });

        const {data} = response;
        // 检查请求是否成功（根据业务状态码）
        const isSuccess = response.status === ResultEnum.SUCCESS;

        // 成功提示处理
        if (isSuccess) {
            if (successMsg) {
                $message.success(data.message);
            } else if (showSuccessMsg && data.message) {
                $message.success(data.message);
            }
        }

        // 根据配置返回完整响应或仅返回数据
        return isReturnResult ? data.data : data;

    } catch (error: unknown) {
        // 增强错误处理
        if (error instanceof Error) {
            // 如果是取消请求的错误，不显示提示
            if (!(error instanceof CanceledError)) {
                const errMsg = error.message || UNKNOWN_ERROR;
                $message.error(errMsg);
            }
            return Promise.reject(error);
        }
        // 处理非 Error 对象的异常情况
        $message.error(UNKNOWN_ERROR);
        return Promise.reject(new Error(UNKNOWN_ERROR));
    }
}