export interface ApiResponse<T> {
  success: string;
  response?: T;
  message?: string;
}

export interface ApiErrorResponse {
  success: string;
  message?: string;
}
