import { useEffect, useState } from 'react';
import { showtimeSeatsApi, type CreateShowtimeSeatRequest } from '../api/showtimeSeats';
import { showtimesApi } from '../api/showtimes';
import { seatsApi } from '../api/seats';
import type { ShowtimeSeat, Showtime, Seat } from '../types';

const SEAT_STATUSES = ['available', 'held', 'booked', 'blocked'];

export default function ShowtimeSeatsPage() {
  const [items, setItems] = useState<ShowtimeSeat[]>([]);
  const [showtimes, setShowtimes] = useState<Showtime[]>([]);
  const [seats, setSeats] = useState<Seat[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [editingId, setEditingId] = useState<string | null>(null);
  const [form, setForm] = useState<CreateShowtimeSeatRequest>({
    showtime_id: '',
    seat_id: '',
    status: 'available',
    price: 0,
  });

  const load = () => {
    setLoading(true);
    Promise.all([showtimeSeatsApi.getAll(), showtimesApi.getAll(), seatsApi.getAll()])
      .then(([ss, st, s]) => {
        setItems(ss.data);
        setShowtimes(st.data);
        setSeats(s.data);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => load(), []);

  const resetForm = () => {
    setForm({ showtime_id: '', seat_id: '', status: 'available', price: 0 });
    setEditingId(null);
    setShowForm(false);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (editingId) {
      await showtimeSeatsApi.update(editingId, form);
    } else {
      await showtimeSeatsApi.create(form);
    }
    resetForm();
    load();
  };

  const handleEdit = (item: ShowtimeSeat) => {
    setForm({ showtime_id: item.showtime_id, seat_id: item.seat_id, status: item.status, price: item.price });
    setEditingId(item.id);
    setShowForm(true);
  };

  const handleDelete = async (id: string) => {
    if (!confirm('Delete this showtime seat?')) return;
    await showtimeSeatsApi.delete(id);
    load();
  };

  const getShowtimeLabel = (id: string) => {
    const st = showtimes.find((s) => s.id === id);
    return st ? `${new Date(st.start_time).toLocaleDateString()} ${st.base_price}$` : id.slice(0, 8);
  };

  const getSeatLabel = (id: string) => {
    const seat = seats.find((s) => s.id === id);
    return seat ? `R${seat.row}S${seat.number}` : id.slice(0, 8);
  };

  return (
    <div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">Showtime Seats</h1>
        <button
          onClick={() => { resetForm(); setShowForm(!showForm); }}
          className="px-4 py-2 bg-indigo-600 hover:bg-indigo-500 text-white text-sm rounded cursor-pointer"
        >
          {showForm ? 'Cancel' : '+ New Showtime Seat'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-gray-800 border border-gray-700 rounded-lg p-4 mb-6 space-y-3">
          <div className="grid grid-cols-2 sm:grid-cols-4 gap-3">
            <select
              value={form.showtime_id}
              onChange={(e) => setForm({ ...form, showtime_id: e.target.value })}
              required
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            >
              <option value="">Select Showtime</option>
              {showtimes.map((s) => (
                <option key={s.id} value={s.id}>
                  {new Date(s.start_time).toLocaleDateString()} - ${s.base_price}
                </option>
              ))}
            </select>
            <select
              value={form.seat_id}
              onChange={(e) => setForm({ ...form, seat_id: e.target.value })}
              required
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            >
              <option value="">Select Seat</option>
              {seats.map((s) => (
                <option key={s.id} value={s.id}>R{s.row} S{s.number}</option>
              ))}
            </select>
            <select
              value={form.status}
              onChange={(e) => setForm({ ...form, status: e.target.value })}
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            >
              {SEAT_STATUSES.map((s) => (
                <option key={s} value={s}>{s}</option>
              ))}
            </select>
            <input
              type="number"
              step="0.01"
              placeholder="Price"
              value={form.price}
              onChange={(e) => setForm({ ...form, price: +e.target.value })}
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
                <th className="text-left py-2 px-3">Showtime</th>
                <th className="text-left py-2 px-3">Seat</th>
                <th className="text-left py-2 px-3">Status</th>
                <th className="text-left py-2 px-3">Price</th>
                <th className="text-right py-2 px-3">Actions</th>
              </tr>
            </thead>
            <tbody>
              {items.map((ss) => (
                <tr key={ss.id} className="border-b border-gray-800">
                  <td className="py-2 px-3">{getShowtimeLabel(ss.showtime_id)}</td>
                  <td className="py-2 px-3">{getSeatLabel(ss.seat_id)}</td>
                  <td className="py-2 px-3">
                    <span className={`px-2 py-0.5 rounded text-xs ${
                      ss.status === 'available' ? 'bg-green-900 text-green-300' :
                      ss.status === 'held' ? 'bg-yellow-900 text-yellow-300' :
                      ss.status === 'booked' ? 'bg-red-900 text-red-300' :
                      'bg-gray-700 text-gray-300'
                    }`}>
                      {ss.status}
                    </span>
                  </td>
                  <td className="py-2 px-3">${ss.price}</td>
                  <td className="py-2 px-3 text-right space-x-2">
                    <button onClick={() => handleEdit(ss)} className="text-indigo-400 hover:text-indigo-300 text-xs cursor-pointer">Edit</button>
                    <button onClick={() => handleDelete(ss.id)} className="text-red-400 hover:text-red-300 text-xs cursor-pointer">Delete</button>
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
