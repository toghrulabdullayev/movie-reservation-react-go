import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { moviesApi } from '../api/movies';
import type { Movie } from '../types';

const TMDB_IMG = 'https://image.tmdb.org/t/p/w500';

export default function MovieDetailPage() {
  const { id } = useParams<{ id: string }>();
  const [movie, setMovie] = useState<Movie | null>(null);
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!id) return;
    moviesApi
      .getById(id)
      .then((res) => setMovie(res.data))
      .catch(() => setError('Movie not found'))
      .finally(() => setLoading(false));
  }, [id]);

  if (loading) return <div className="text-gray-400 py-10 text-center">Loading...</div>;
  if (error) return <div className="text-red-400 py-10 text-center">{error}</div>;
  if (!movie) return null;

  return (
    <div>
      {movie.backdrop_path && (
        <img
          src={`${TMDB_IMG}${movie.backdrop_path}`}
          alt={movie.title}
          className="w-full h-64 sm:h-80 object-cover rounded-lg mb-6"
        />
      )}
      <div className="flex gap-6 flex-col sm:flex-row">
        {movie.poster_path && (
          <img
            src={`${TMDB_IMG}${movie.poster_path}`}
            alt={movie.title}
            className="w-48 rounded-lg self-start"
          />
        )}
        <div>
          <h1 className="text-3xl font-bold mb-2">{movie.title}</h1>
          <div className="flex items-center gap-3 text-sm text-gray-400 mb-4">
            <span>{movie.release_date?.split('T')[0] ?? 'N/A'}</span>
            <span>&middot;</span>
            <span>{movie.runtime} min</span>
            <span>&middot;</span>
            <span className="text-yellow-400">★ {movie.vote_average}</span>
          </div>
          <div className="flex flex-wrap gap-2 mb-4">
            {movie.genres.map((g) => (
              <span
                key={g}
                className="px-2 py-0.5 bg-gray-800 border border-gray-700 rounded text-xs text-gray-300"
              >
                {g}
              </span>
            ))}
          </div>
          <p className="text-gray-300 leading-relaxed">{movie.overview}</p>
        </div>
      </div>
    </div>
  );
}
