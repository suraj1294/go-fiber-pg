import { ApiResponse } from "./api-response";

export interface AuthLogin {
  accessToken: string;
}

export interface AuthUser {
  id?: number;
  first_name?: string;
  last_name?: string;
  accessToken?: string;
}

export type AuthLoginApiResponse = ApiResponse<AuthLogin>;
export type AuthUserApiResponse = ApiResponse<AuthUser>;
