import api from './client';
import type { User } from '../types';

export interface LoginRequest {
  identity: string;
  password: string;
}

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

export interface AuthResponse {
  message: string;
  user: User;
}

export interface MeResponse {
  user: Pick<User, 'id' | 'role'>;
}

export const authApi = {
  register: (data: RegisterRequest) =>
    api.post<AuthResponse>('/auth/register', data),

  login: (data: LoginRequest) =>
    api.post<AuthResponse>('/auth/login', data),

  refresh: () => api.post<AuthResponse>('/auth/refresh'),

  logout: () => api.post<{ message: string }>('/auth/logout'),

  me: () => api.get<MeResponse>('/auth/me'),
};
