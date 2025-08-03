// 配置全局基础配置
import _ from "lodash";
import axios from "axios";
import router from "@/router";
import webConfig from "./global.config";
// 使用element-ui Message做消息提醒
import { Message } from "element-ui";
import Vue from "vue";
import domMessage from "@/assets/tool/message_once";
import store from "../store";

import { JumpDialogBox } from "@/assets/tool/login";
// import base64 from 'js-base64'
const messageOnce = new domMessage();
let request = axios.create({
    // 基础配置
    // baseURL: process.env.VUE_APP_DASHBOARD_SERVER_PATH || '/ws',
    baseURL:
        process.env.NODE_ENV === "out"
            ? window.WEBCONFIG.VUE_APP_DASHBOARD_SERVER_PATH
            : process.env.VUE_APP_DASHBOARD_SERVER_PATH,
    timeout: 30 * 1000,
    // responseType: 'json',
});
request.interceptors.request.use(
    (config) => {
        config.cancelToken = new axios.CancelToken(function (cancel) {
            store.commit("pushToken", cancel);
        });
        if (!config.headers || !config.headers["Content-Type"]) {
            // config.headers["Content-Type"] =
            //     "application/x-www-form-urlencoded";
            config.headers["Content-Type"] =
                "application/json";
        }

        let token = localStorage.getItem("token");
        // if (webConfig.indexOf(url) === -1 && token) {
        //     config.headers.token = token
        // }
        if (token) {
            config.headers.token = token;
        }
        // token加密
        // let _secret = base64.encode(webConfig.secretId + new Date().toString())
        // config.headers.secret = _secret
        return config;
    },
    (error) => {
        Promise.reject(error);
    }
);

request.interceptors.response.use(
    (response) => {
        let data = response.data;
        if (data && data instanceof Blob) {
            let blob = new Blob([data], { type: "application/octet-stream" });
            let name = response.headers["filename"];
            let fileName = decodeURIComponent(name);
            let url = URL.createObjectURL(blob);
            let a = document.createElement("a");
            a.download = fileName; // 设置下载的文件名
            a.href = url;
            a.click();
            window.URL.revokeObjectURL(url);
        } else if (data) {
            if (data.resultCode !== 200) {
                if (data.resultCode !== 401 && data.resultCode !== 403) {
                    messageOnce.error({
                        message: response.data.message,
                        duration: 2000,
                        showClose: true,
                    });
                } else {
                    setTimeout(() => {
                        JumpDialogBox.install({
                            code: data.resultCode,
                            msg: response.data.message,
                        });
                    }, 0);
                    store.commit("clearToken"); // 取消请求
                }
            }
        }
        return response.data;
    },
    (error) => {
        /***** 接收到异常响应的处理开始 *****/
        if (error && error.response) {
            // 1.公共错误处理
            // 2.根据响应码具体处理
            switch (error.response.status) {
                case 400:
                    error.message = "错误请求";
                    break;
                // case 401:
                //     error.message = '未授权，请重新登录'
                //     break;
                // case 403:
                //     error.message = '拒绝访问'
                //     break;
                case 404:
                    error.message = "请求错误,未找到该资源";
                    // window.location.href = "/404"
                    break;
                case 405:
                    error.message = "请求方法未允许";
                    break;
                case 408:
                    error.message = "请求超时";
                    break;
                case 500:
                    error.message = "服务器端出错";
                    break;
                case 501:
                    error.message = "网络未实现";
                    break;
                case 502:
                    error.message = "网络错误";
                    break;
                case 503:
                    error.message = "服务不可用";
                    break;
                case 504:
                    error.message = "网络超时";
                    break;
                case 505:
                    error.message = "http版本不支持该请求";
                    break;
                default:
                    error.message = `连接错误${error.response.status}`;
            }
        } else {
            // 超时处理
            if (JSON.stringify(error).includes("timeout")) {
                error.message = "服务器响应超时，请刷新当前页";
            } else if (error && error.message !== "canceled") {
                error.message = "连接服务器失败";
            }
        }
        // Message.error(error.message)
        if (error && error.message !== "canceled") {
            messageOnce.error({
                message: error.message,
                type: "error",
                duration: 1000,
                showClose: true,
            });
        }
        return Promise.resolve(error.response);
    }
);
let previewRequest = axios.create({
    // 基础配置
    // baseURL: process.env.VUE_APP_MARKTABLE_SERVER_PATH || '/ws',
    baseURL:
        process.env.NODE_ENV === "out"
            ? window.WEBCONFIG.VUE_APP_MARKTABLE_SERVER_PATH
            : process.env.VUE_APP_MARKTABLE_SERVER_PATH,
    timeout: 30 * 1000,
    // responseType: 'json',
});
previewRequest.interceptors.request.use(
    (config) => {
        config.cancelToken = new axios.CancelToken(function (cancel) {
            store.commit("pushToken", cancel);
        });
        if (!config.headers || !config.headers["Content-Type"]) {
            config.headers["Content-Type"] =
                "application/x-www-form-urlencoded";
        }

        let token = localStorage.getItem("token");
        // if (webConfig.indexOf(url) === -1 && token) {
        //     config.headers.token = token
        // }
        if (token) {
            config.headers.token = token;
        }
        // token加密
        // let _secret = base64.encode(webConfig.secretId + new Date().toString())
        // config.headers.secret = _secret
        return config;
    },
    (error) => {
        Promise.reject(error);
    }
);

previewRequest.interceptors.response.use(
    (response) => {
        let data = response.data;
        if (data && data instanceof Blob) {
            return response.data;
        }
    },
    (error) => {
        /***** 接收到异常响应的处理开始 *****/
        if (error && error.response) {
            // 1.公共错误处理
            // 2.根据响应码具体处理
            switch (error.response.status) {
                case 400:
                    error.message = "错误请求";
                    break;
                // case 401:
                //     error.message = '未授权，请重新登录'
                //     break;
                // case 403:
                //     error.message = '拒绝访问'
                //     break;
                case 404:
                    error.message = "请求错误,未找到该资源";
                    // window.location.href = "/404"
                    break;
                case 405:
                    error.message = "请求方法未允许";
                    break;
                case 408:
                    error.message = "请求超时";
                    break;
                case 500:
                    error.message = "服务器端出错";
                    break;
                case 501:
                    error.message = "网络未实现";
                    break;
                case 502:
                    error.message = "网络错误";
                    break;
                case 503:
                    error.message = "服务不可用";
                    break;
                case 504:
                    error.message = "网络超时";
                    break;
                case 505:
                    error.message = "http版本不支持该请求";
                    break;
                default:
                    error.message = `连接错误${error.response.status}`;
            }
        } else {
            // 超时处理
            if (JSON.stringify(error).includes("timeout")) {
                error.message = "服务器响应超时，请刷新当前页";
            } else if (error && error.message !== "canceled") {
                error.message = "连接服务器失败";
            }
        }
        // Message.error(error.message)
        if (error && error.message !== "canceled") {
            messageOnce.error({
                message: error.message,
                type: "error",
                duration: 1000,
                showClose: true,
            });
        }
        return Promise.resolve(error.response);
    }
);
const http = {
    /**
     * methods: 请求
     * @param url 请求地址
     * @param params 请求参数
     */
    get(url, params) {
        const config = {
            method: "get",
            url: url,
        };
        if (params) config.params = params;
        return request(config);
    },
    post(url, params) {
        const config = {
            method: "post",
            url: url,
        };
        if (params) config.data = JSON.stringify(params);
        return request(config);
    },
    put(url, params) {
        const config = {
            method: "put",
            url: url,
        };
        if (params) config.params = params;
        return request(config);
    },
    delete(url, params) {
        const config = {
            method: "delete",
            url: url,
        };
        if (params) config.params = params;
        return request(config);
    },
    uplodFile(url, params, uploadChange) {
        const headers = {
            "Content-Type": "multipart/form-data",
        };
        const config = {
            method: "post",
            url: url,
            headers,
        };
        if (params) config.data = params;
        if (uploadChange) config.onUploadProgress = uploadChange
        return request(config);
    },
    downloadFile(url, params) {
        const config = {
            method: "get",
            url: url,
            responseType: "blob",
        };
        if (params) config.params = params;
        return request(config);
    },
    // 文件预览
    previewFile(url, params) {
        const config = {
            method: "get",
            url: url,
            responseType: "blob",
        };
        if (params) config.params = params;
        return previewRequest(config);
    },
};

export default http;
