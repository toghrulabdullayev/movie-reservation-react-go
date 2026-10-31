import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { moviesApi } from '../api/movies';
import type { Movie, PaginationMetadata } from '../types';

const TMDB_IMG = 'https://image.tmdb.org/t/p/w300';

export default function MoviesPage() {
  const [movies, setMovies] = useState<Movie[]>([]);
  const [meta, setMeta] = useState<PaginationMetadata | null>(null);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setLoading(true);
    moviesApi
      .getAll(page, 12)
      .then((res) => {
        setMovies(res.data.data);
        setMeta(res.data.metadata);
      })
      .finally(() => setLoading(false));
  }, [page]);

  if (loading) return <div className="text-gray-400 py-10 text-center">Loading movies...</div>;

  return (
    <div>
      <h1 className="text-2xl font-bold mb-6">Movies</h1>
      <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-4">
        {movies.map((m) => (
          <Link
            key={m.id}
            to={`/movies/${m.id}`}
            className="bg-gray-800 border border-gray-700 rounded-lg overflow-hidden hover:border-indigo-500 transition-colors"
          >
            {m.poster_path ? (
              <img
                src={`${TMDB_IMG}${m.poster_path}`}
                alt={m.title}
                className="w-full h-64 object-cover"
              />
            ) : (
              <div className="w-full h-64 bg-gray-700 flex items-center justify-center text-gray-500 text-sm">
                No Poster
              </div>
            )}
            <div className="p-3">
              <h3 className="font-medium text-sm truncate">{m.title}</h3>
              <p className="text-xs text-gray-400 mt-1">
                {m.release_date?.split('T')[0] ?? 'N/A'} &middot; {m.runtime}m
              </p>
              <p className="text-xs text-yellow-400 mt-1">
                {'★'} {m.vote_average}
              </p>
            </div>
          </Link>
        ))}
      </div>
      {meta && meta.total_pages > 1 && (
        <div className="flex items-center justify-center gap-4 mt-8">
          <button
            onClick={() => setPage((p) => Math.max(1, p - 1))}
            disabled={page === 1}
            className="px-3 py-1 bg-gray-800 border border-gray-700 rounded text-sm disabled:opacity-40 cursor-pointer"
          >
            Prev
          </button>
          <span className="text-sm text-gray-400">
            Page {meta.current_page} / {meta.total_pages}
          </span>
          <button
            onClick={() => setPage((p) => Math.min(meta.total_pages, p + 1))}
            disabled={page === meta.total_pages}
            className="px-3 py-1 bg-gray-800 border border-gray-700 rounded text-sm disabled:opacity-40 cursor-pointer"
          >
            Next
          </button>
        </div>
      )}
    </div>
  );
}
