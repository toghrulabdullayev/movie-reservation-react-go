export interface User {
  id: string;
  username: string;
  email: string;
  role: 'user' | 'admin';
}

export interface Movie {
  id: string;
  tmdb_id: number;
  title: string;
  overview: string;
  poster_path: string | null;
  backdrop_path: string | null;
  release_date: string | null;
  runtime: number;
  vote_average: number;
  genres: string[];
}

export interface Room {
  id: string;
  name: string;
  capacity: number;
  rows: number;
  columns: number;
}

export interface SeatType {
  id: string;
  name: string;
  multiplier: number;
}

export interface Seat {
  id: string;
  room_id: string;
  seat_type_id: string;
  row: number;
  number: number;
}

export interface Showtime {
  id: string;
  movie_id: string;
  room_id: string;
  start_time: string;
  end_time: string;
  base_price: number;
  status: 'scheduled' | 'ongoing' | 'completed' | 'cancelled';
}

export interface ShowtimeSeat {
  id: string;
  showtime_id: string;
  seat_id: string;
  status: 'available' | 'held' | 'booked' | 'blocked';
  price: number;
}

export interface Reservation {
  id: string;
  user_id: string;
  showtime_id: string;
  status: 'pending' | 'confirmed' | 'cancelled';
  total_price: number;
  reserved_at: string;
}

export interface PaginationMetadata {
  current_page: number;
  page_size: number;
  total_records: number;
  total_pages: number;
}

export interface PaginatedResponse<T> {
  metadata: PaginationMetadata;
  data: T[];
}
