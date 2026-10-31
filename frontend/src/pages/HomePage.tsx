import { Link } from 'react-router-dom';

export default function HomePage() {
  return (
    <div>
      <div className="text-center py-20">
        <h1 className="text-4xl font-bold mb-4">CineBook</h1>
        <p className="text-gray-400 mb-8 max-w-md mx-auto">
          Movie reservation system. Browse movies, manage rooms and showtimes,
          and book your seats.
        </p>
      </div>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <Link
          to="/movies"
          className="bg-gray-800 hover:bg-gray-750 border border-gray-700 rounded-lg p-6 transition-colors"
        >
          <h2 className="text-lg font-semibold mb-2">Movies</h2>
          <p className="text-sm text-gray-400">Browse available movies</p>
        </Link>
        <Link
          to="/rooms"
          className="bg-gray-800 hover:bg-gray-750 border border-gray-700 rounded-lg p-6 transition-colors"
        >
          <h2 className="text-lg font-semibold mb-2">Rooms</h2>
          <p className="text-sm text-gray-400">Manage cinema rooms</p>
        </Link>
        <Link
          to="/seat-types"
          className="bg-gray-800 hover:bg-gray-750 border border-gray-700 rounded-lg p-6 transition-colors"
        >
          <h2 className="text-lg font-semibold mb-2">Seat Types</h2>
          <p className="text-sm text-gray-400">Configure seat types & pricing</p>
        </Link>
        <Link
          to="/showtimes"
          className="bg-gray-800 hover:bg-gray-750 border border-gray-700 rounded-lg p-6 transition-colors"
        >
          <h2 className="text-lg font-semibold mb-2">Showtimes</h2>
          <p className="text-sm text-gray-400">Schedule and manage showtimes</p>
        </Link>
        <Link
          to="/reservations"
          className="bg-gray-800 hover:bg-gray-750 border border-gray-700 rounded-lg p-6 transition-colors"
        >
          <h2 className="text-lg font-semibold mb-2">Reservations</h2>
          <p className="text-sm text-gray-400">View and manage reservations</p>
        </Link>
      </div>
    </div>
  );
}
