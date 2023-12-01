import { MoviesApiResponse } from "@/models/movie";
import { axiosClient } from "./axios-base";

export const getMovies = () => axiosClient.get<MoviesApiResponse>("movies");
