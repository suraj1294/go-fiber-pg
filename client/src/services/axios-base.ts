import { getAuthState, setAuthState } from "@/auth/auth-store";
import axios, { AxiosResponse, InternalAxiosRequestConfig } from "axios";
import { refreshToken } from "./auth";
import { ApiErrorResponse } from "@/models/api-response";

export const axiosClient = axios.create({
  baseURL: "/api",
  withCredentials: true,
});

axiosClient.interceptors.request.use(
  (config) => {
    const token = getAuthState().token;
    if (!config?.headers?.Authorization && token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (err) => Promise.reject(err)
);

axiosClient.interceptors.response.use(
  (res) => res,
  async (err) => {
    const prevReq = err?.config as InternalAxiosRequestConfig<unknown> & {
      sent: boolean;
    };

    const response = err?.response as AxiosResponse<ApiErrorResponse>;

    const errMessage = response?.data?.message;

    if (
      (errMessage === "unauthorized" || response?.status === 401) &&
      !prevReq?.sent &&
      !prevReq?.url?.includes("refresh")
    ) {
      prevReq.sent = true;
      try {
        const data = await refreshToken();
        setAuthState((prev) => ({ ...prev, token: data?.accessToken }));
        prevReq.headers.Authorization = `Bearer ${data?.accessToken}`;
      } catch {
        setAuthState((prev) => ({ ...prev, token: "" }));
        return Promise.reject(err);
      }

      return axiosClient(prevReq);
    }

    return Promise.reject(err);
  }
);
