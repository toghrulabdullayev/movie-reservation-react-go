import { useEffect, useState } from 'react';
import { showtimesApi, type CreateShowtimeRequest } from '../api/showtimes';
import { moviesApi } from '../api/movies';
import { roomsApi } from '../api/rooms';
import type { Showtime, Movie, Room } from '../types';

const STATUSES = ['scheduled', 'ongoing', 'completed', 'cancelled'];

export default function ShowtimesPage() {
  const [showtimes, setShowtimes] = useState<Showtime[]>([]);
  const [movies, setMovies] = useState<Movie[]>([]);
  const [rooms, setRooms] = useState<Room[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [editingId, setEditingId] = useState<string | null>(null);
  const [form, setForm] = useState<CreateShowtimeRequest>({
    movie_id: '',
    room_id: '',
    start_time: '',
    end_time: '',
    base_price: 10,
    status: 'scheduled',
  });

  const load = () => {
    setLoading(true);
    Promise.all([
      showtimesApi.getAll(),
      moviesApi.getAll(1, 100),
      roomsApi.getAll(1, 100),
    ])
      .then(([s, m, r]) => {
        setShowtimes(s.data);
        setMovies(m.data.data);
        setRooms(r.data.data);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => load(), []);

  const resetForm = () => {
    setForm({ movie_id: '', room_id: '', start_time: '', end_time: '', base_price: 10, status: 'scheduled' });
    setEditingId(null);
    setShowForm(false);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (editingId) {
      await showtimesApi.update(editingId, form);
    } else {
      await showtimesApi.create(form);
    }
    resetForm();
    load();
  };

  const handleEdit = (st: Showtime) => {
    setForm({
      movie_id: st.movie_id,
      room_id: st.room_id,
      start_time: st.start_time.slice(0, 16),
      end_time: st.end_time.slice(0, 16),
      base_price: st.base_price,
      status: st.status,
    });
    setEditingId(st.id);
    setShowForm(true);
  };

  const handleDelete = async (id: string) => {
    if (!confirm('Delete this showtime?')) return;
    await showtimesApi.delete(id);
    load();
  };

  const getMovieTitle = (id: string) => movies.find((m) => m.id === id)?.title ?? id.slice(0, 8);
  const getRoomName = (id: string) => rooms.find((r) => r.id === id)?.name ?? id.slice(0, 8);

  return (
    <div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">Showtimes</h1>
        <button
          onClick={() => { resetForm(); setShowForm(!showForm); }}
          className="px-4 py-2 bg-indigo-600 hover:bg-indigo-500 text-white text-sm rounded cursor-pointer"
        >
          {showForm ? 'Cancel' : '+ New Showtime'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-gray-800 border border-gray-700 rounded-lg p-4 mb-6 space-y-3">
          <div className="grid grid-cols-2 sm:grid-cols-3 gap-3">
            <select
              value={form.movie_id}
              onChange={(e) => setForm({ ...form, movie_id: e.target.value })}
              required
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            >
              <option value="">Select Movie</option>
              {movies.map((m) => (
                <option key={m.id} value={m.id}>{m.title}</option>
              ))}
            </select>
            <select
              value={form.room_id}
              onChange={(e) => setForm({ ...form, room_id: e.target.value })}
              required
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            >
              <option value="">Select Room</option>
              {rooms.map((r) => (
                <option key={r.id} value={r.id}>{r.name}</option>
              ))}
            </select>
            <select
              value={form.status}
              onChange={(e) => setForm({ ...form, status: e.target.value })}
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            >
              {STATUSES.map((s) => (
                <option key={s} value={s}>{s}</option>
              ))}
            </select>
            <input
              type="datetime-local"
              value={form.start_time}
              onChange={(e) => setForm({ ...form, start_time: e.target.value })}
              required
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            />
            <input
              type="datetime-local"
              value={form.end_time}
              onChange={(e) => setForm({ ...form, end_time: e.target.value })}
              required
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            />
            <input
              type="number"
              step="0.01"
              placeholder="Base Price"
              value={form.base_price}
              onChange={(e) => setForm({ ...form, base_price: +e.target.value })}
              required
              min={0}
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            />
          </div>
          <button type="submit" className="px-4 py-2 bg-green-600 hover:bg-green-500 text-white text-sm rounded cursor-pointer">
            {editingId ? 'Update' : 'Create'}
          </button>
        </form>
      )}

      {loading ? (
        <div className="text-gray-400 py-10 text-center">Loading...</div>
      ) : (
        <div className="overflow-x-auto">
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b border-gray-700 text-gray-400">
                <th className="text-left py-2 px-3">Movie</th>
                <th className="text-left py-2 px-3">Room</th>
                <th className="text-left py-2 px-3">Start</th>
                <th className="text-left py-2 px-3">End</th>
                <th className="text-left py-2 px-3">Price</th>
                <th className="text-left py-2 px-3">Status</th>
                <th className="text-right py-2 px-3">Actions</th>
              </tr>
            </thead>
            <tbody>
              {showtimes.map((s) => (
                <tr key={s.id} className="border-b border-gray-800">
                  <td className="py-2 px-3 truncate max-w-[150px]">{getMovieTitle(s.movie_id)}</td>
                  <td className="py-2 px-3">{getRoomName(s.room_id)}</td>
                  <td className="py-2 px-3">{new Date(s.start_time).toLocaleString()}</td>
                  <td className="py-2 px-3">{new Date(s.end_time).toLocaleString()}</td>
                  <td className="py-2 px-3">${s.base_price}</td>
                  <td className="py-2 px-3">
                    <span className={`px-2 py-0.5 rounded text-xs ${
                      s.status === 'scheduled' ? 'bg-blue-900 text-blue-300' :
                      s.status === 'ongoing' ? 'bg-green-900 text-green-300' :
                      s.status === 'completed' ? 'bg-gray-700 text-gray-300' :
                      'bg-red-900 text-red-300'
                    }`}>
                      {s.status}
                    </span>
                  </td>
                  <td className="py-2 px-3 text-right space-x-2">
                    <button onClick={() => handleEdit(s)} className="text-indigo-400 hover:text-indigo-300 text-xs cursor-pointer">Edit</button>
                    <button onClick={() => handleDelete(s.id)} className="text-red-400 hover:text-red-300 text-xs cursor-pointer">Delete</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
