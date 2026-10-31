import api from './client';
import type { Movie, PaginatedResponse } from '../types';

export const moviesApi = {
  getAll: (page = 1, limit = 10) =>
    api.get<PaginatedResponse<Movie>>('/movies/', { params: { page, limit } }),

  getById: (id: string) => api.get<Movie>(`/movies/${id}`),
};
