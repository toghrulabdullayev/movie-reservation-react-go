import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';
import Layout from './components/Layout';
import ProtectedRoute from './components/ProtectedRoute';
import HomePage from './pages/HomePage';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import MoviesPage from './pages/MoviesPage';
import MovieDetailPage from './pages/MovieDetailPage';
import RoomsPage from './pages/RoomsPage';
import SeatTypesPage from './pages/SeatTypesPage';
import SeatsPage from './pages/SeatsPage';
import ShowtimesPage from './pages/ShowtimesPage';
import ShowtimeSeatsPage from './pages/ShowtimeSeatsPage';
import ReservationsPage from './pages/ReservationsPage';

export default function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <Routes>
          <Route element={<Layout />}>
            <Route index element={<HomePage />} />
            <Route path="login" element={<LoginPage />} />
            <Route path="register" element={<RegisterPage />} />
            <Route path="movies" element={<MoviesPage />} />
            <Route path="movies/:id" element={<MovieDetailPage />} />
            <Route element={<ProtectedRoute />}>
              <Route path="rooms" element={<RoomsPage />} />
              <Route path="seat-types" element={<SeatTypesPage />} />
              <Route path="seats" element={<SeatsPage />} />
              <Route path="showtimes" element={<ShowtimesPage />} />
              <Route path="showtime-seats" element={<ShowtimeSeatsPage />} />
              <Route path="reservations" element={<ReservationsPage />} />
            </Route>
          </Route>
        </Routes>
      </BrowserRouter>
    </AuthProvider>
  );
}
