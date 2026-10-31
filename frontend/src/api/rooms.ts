import api from './client';
import type { Room, PaginatedResponse } from '../types';

export interface CreateRoomRequest {
  name: string;
  capacity: number;
  rows: number;
  columns: number;
}

export const roomsApi = {
  getAll: (page = 1, limit = 10) =>
    api.get<PaginatedResponse<Room>>('/rooms/', { params: { page, limit } }),

  getById: (id: string) => api.get<Room>(`/rooms/${id}`),

  create: (data: CreateRoomRequest) => api.post<Room>('/rooms/', data),

  update: (id: string, data: CreateRoomRequest) =>
    api.put<Room>(`/rooms/${id}`, data),

  delete: (id: string) => api.delete(`/rooms/${id}`),
};
