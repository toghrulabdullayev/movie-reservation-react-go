import api from './client';
import type { Showtime } from '../types';

export interface CreateShowtimeRequest {
  movie_id: string;
  room_id: string;
  start_time: string;
  end_time: string;
  base_price: number;
  status: string;
}

export const showtimesApi = {
  getAll: () => api.get<Showtime[]>('/showtimes/'),

  getById: (id: string) => api.get<Showtime>(`/showtimes/${id}`),

  create: (data: CreateShowtimeRequest) =>
    api.post<Showtime>('/showtimes/', data),

  update: (id: string, data: CreateShowtimeRequest) =>
    api.put(`/showtimes/${id}`, data),

  delete: (id: string) => api.delete(`/showtimes/${id}`),
};
