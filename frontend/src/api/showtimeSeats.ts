import api from './client';
import type { ShowtimeSeat } from '../types';

export interface CreateShowtimeSeatRequest {
  showtime_id: string;
  seat_id: string;
  status: string;
  price: number;
}

export const showtimeSeatsApi = {
  getAll: () => api.get<ShowtimeSeat[]>('/showtimes/seats/'),

  getById: (id: string) => api.get<ShowtimeSeat>(`/showtimes/seats/${id}`),

  create: (data: CreateShowtimeSeatRequest) =>
    api.post('/showtimes/seats/', data),

  update: (id: string, data: CreateShowtimeSeatRequest) =>
    api.put(`/showtimes/seats/${id}`, data),

  delete: (id: string) => api.delete(`/showtimes/seats/${id}`),
};
