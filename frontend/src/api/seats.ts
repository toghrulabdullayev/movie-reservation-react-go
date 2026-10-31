import api from './client';
import type { Seat } from '../types';

export interface CreateSeatRequest {
  room_id: string;
  seat_type_id: string;
  row: number;
  number: number;
}

export const seatsApi = {
  getAll: () => api.get<Seat[]>('/seats/'),

  getById: (id: string) => api.get<Seat>(`/seats/${id}`),

  create: (data: CreateSeatRequest) => api.post<Seat>('/seats/', data),

  update: (id: string, data: CreateSeatRequest) =>
    api.put<Seat>(`/seats/${id}`, data),

  delete: (id: string) => api.delete(`/seats/${id}`),
};
