import { AuthLoginApiResponse, AuthUserApiResponse } from "@/models/auth";
import { axiosClient } from "./axios-base";

import { handlerApiError } from "@/utils/handler-error";
import { ApiResponse } from "@/models/api-response";

export const login = async ({
  email,
  password,
}: {
  email: string;
  password: string;
}) => {
  try {
    const { data } = await axiosClient.post<AuthLoginApiResponse>(
      "/auth/login",
      {
        email,
        password,
      }
    );

    return Promise.resolve(data?.response);
  } catch (error) {
    return Promise.reject(handlerApiError(error));
  }
};

export const me = async () => {
  try {
    const { data } = await axiosClient.get<AuthUserApiResponse>("/auth/me");

    return Promise.resolve(data?.response);
  } catch (error) {
    return Promise.reject(handlerApiError(error));
  }
};

export const logout = async () => {
  try {
    const { data } = await axiosClient.get<ApiResponse<string>>("/auth/logout");

    return Promise.resolve(data?.response);
  } catch (error) {
    return Promise.reject(handlerApiError(error));
  }
};
