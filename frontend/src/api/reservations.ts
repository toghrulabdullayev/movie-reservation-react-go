import api from './client';
import type { Reservation, PaginatedResponse } from '../types';

export interface CreateReservationRequest {
  user_id: string;
  showtime_id: string;
  total_price: number;
}

export interface UpdateReservationRequest {
  user_id: string;
  showtime_id: string;
  status: string;
  total_price: number;
}

export const reservationsApi = {
  getAll: (page = 1, limit = 10) =>
    api.get<PaginatedResponse<Reservation>>('/reservations/', {
      params: { page, limit },
    }),

  getById: (id: string) => api.get<Reservation>(`/reservations/${id}`),

  create: (data: CreateReservationRequest) =>
    api.post<Reservation>('/reservations/', data),

  update: (id: string, data: UpdateReservationRequest) =>
    api.put<Reservation>(`/reservations/${id}`, data),

  delete: (id: string) => api.delete(`/reservations/${id}`),
};
