import { ApiErrorResponse } from "@/models/api-response";
import axios from "axios";

const defaultError = {
  success: false,
  message: "something went wrong",
  status: 500,
};

export const handlerApiError = (error: unknown): ApiErrorResponse => {
  return {
    ...defaultError,
    ...(axios.isAxiosError(error) && {
      ...error?.response?.data,
      status: error.status,
    }),
  };
};
