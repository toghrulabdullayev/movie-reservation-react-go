import api from './client';
import type { SeatType } from '../types';

export interface CreateSeatTypeRequest {
  name: string;
  multiplier: number;
}

export const seatTypesApi = {
  getAll: () => api.get<SeatType[]>('/seats/types/'),

  getById: (id: string) => api.get<SeatType>(`/seats/types/${id}`),

  create: (data: CreateSeatTypeRequest) =>
    api.post<SeatType>('/seats/types/', data),

  update: (id: string, data: CreateSeatTypeRequest) =>
    api.put<SeatType>(`/seats/types/${id}`, data),

  delete: (id: string) => api.delete(`/seats/types/${id}`),
};
