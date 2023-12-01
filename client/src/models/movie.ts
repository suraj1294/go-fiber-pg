import { ApiResponse } from "./api-response";

export interface Movie {
  id: number;
  title: string;
  release_date: string;
  runtime: number;
  MPAARating: string;
  description: string;
  image: string;
}

export type MoviesApiResponse = ApiResponse<Movie[]>;
