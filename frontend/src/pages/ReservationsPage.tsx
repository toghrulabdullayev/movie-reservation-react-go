import { useEffect, useState } from 'react';
import { reservationsApi, type CreateReservationRequest, type UpdateReservationRequest } from '../api/reservations';
import { showtimesApi } from '../api/showtimes';
import type { Reservation, Showtime, PaginationMetadata } from '../types';

const STATUSES = ['pending', 'confirmed', 'cancelled'];

export default function ReservationsPage() {
  const [reservations, setReservations] = useState<Reservation[]>([]);
  const [showtimes, setShowtimes] = useState<Showtime[]>([]);
  const [meta, setMeta] = useState<PaginationMetadata | null>(null);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [editingId, setEditingId] = useState<string | null>(null);
  const [form, setForm] = useState<CreateReservationRequest>({
    user_id: '',
    showtime_id: '',
    total_price: 0,
  });
  const [formStatus, setFormStatus] = useState('pending');

  const load = (p: number) => {
    setLoading(true);
    Promise.all([reservationsApi.getAll(p, 10), showtimesApi.getAll()])
      .then(([r, s]) => {
        setReservations(r.data.data);
        setMeta(r.data.metadata);
        setShowtimes(s.data);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => load(page), [page]);

  const resetForm = () => {
    setForm({ user_id: '', showtime_id: '', total_price: 0 });
    setFormStatus('pending');
    setEditingId(null);
    setShowForm(false);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (editingId) {
      const updateData: UpdateReservationRequest = {
        ...form,
        status: formStatus,
      };
      await reservationsApi.update(editingId, updateData);
    } else {
      await reservationsApi.create(form);
    }
    resetForm();
    load(page);
  };

  const handleEdit = (r: Reservation) => {
    setForm({ user_id: r.user_id, showtime_id: r.showtime_id, total_price: r.total_price });
    setFormStatus(r.status);
    setEditingId(r.id);
    setShowForm(true);
  };

  const handleDelete = async (id: string) => {
    if (!confirm('Delete this reservation?')) return;
    await reservationsApi.delete(id);
    load(page);
  };

  const getShowtimeLabel = (id: string) => {
    const st = showtimes.find((s) => s.id === id);
    return st ? new Date(st.start_time).toLocaleDateString() : id.slice(0, 8);
  };

  return (
    <div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">Reservations</h1>
        <button
          onClick={() => { resetForm(); setShowForm(!showForm); }}
          className="px-4 py-2 bg-indigo-600 hover:bg-indigo-500 text-white text-sm rounded cursor-pointer"
        >
          {showForm ? 'Cancel' : '+ New Reservation'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-gray-800 border border-gray-700 rounded-lg p-4 mb-6 space-y-3">
          <div className="grid grid-cols-2 sm:grid-cols-4 gap-3">
            <input
              placeholder="User ID"
              value={form.user_id}
              onChange={(e) => setForm({ ...form, user_id: e.target.value })}
              required
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            />
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
            <input
              type="number"
              step="0.01"
              placeholder="Total Price"
              value={form.total_price}
              onChange={(e) => setForm({ ...form, total_price: +e.target.value })}
              required
              min={0}
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            />
            {editingId && (
              <select
                value={formStatus}
                onChange={(e) => setFormStatus(e.target.value)}
                className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
              >
                {STATUSES.map((s) => (
                  <option key={s} value={s}>{s}</option>
                ))}
              </select>
            )}
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
                <th className="text-left py-2 px-3">ID</th>
                <th className="text-left py-2 px-3">User</th>
                <th className="text-left py-2 px-3">Showtime</th>
                <th className="text-left py-2 px-3">Status</th>
                <th className="text-left py-2 px-3">Total</th>
                <th className="text-left py-2 px-3">Reserved At</th>
                <th className="text-right py-2 px-3">Actions</th>
              </tr>
            </thead>
            <tbody>
              {reservations.map((r) => (
                <tr key={r.id} className="border-b border-gray-800">
                  <td className="py-2 px-3 font-mono text-xs">{r.id.slice(0, 8)}...</td>
                  <td className="py-2 px-3 font-mono text-xs">{r.user_id.slice(0, 8)}...</td>
                  <td className="py-2 px-3">{getShowtimeLabel(r.showtime_id)}</td>
                  <td className="py-2 px-3">
                    <span className={`px-2 py-0.5 rounded text-xs ${
                      r.status === 'confirmed' ? 'bg-green-900 text-green-300' :
                      r.status === 'cancelled' ? 'bg-red-900 text-red-300' :
                      'bg-yellow-900 text-yellow-300'
                    }`}>
                      {r.status}
                    </span>
                  </td>
                  <td className="py-2 px-3">${r.total_price}</td>
                  <td className="py-2 px-3">{new Date(r.reserved_at).toLocaleString()}</td>
                  <td className="py-2 px-3 text-right space-x-2">
                    <button onClick={() => handleEdit(r)} className="text-indigo-400 hover:text-indigo-300 text-xs cursor-pointer">Edit</button>
                    <button onClick={() => handleDelete(r.id)} className="text-red-400 hover:text-red-300 text-xs cursor-pointer">Delete</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {meta && meta.total_pages > 1 && (
        <div className="flex items-center justify-center gap-4 mt-6">
          <button onClick={() => setPage((p) => Math.max(1, p - 1))} disabled={page === 1} className="px-3 py-1 bg-gray-800 border border-gray-700 rounded text-sm disabled:opacity-40 cursor-pointer">Prev</button>
          <span className="text-sm text-gray-400">Page {meta.current_page} / {meta.total_pages}</span>
          <button onClick={() => setPage((p) => Math.min(meta.total_pages, p + 1))} disabled={page === meta.total_pages} className="px-3 py-1 bg-gray-800 border border-gray-700 rounded text-sm disabled:opacity-40 cursor-pointer">Next</button>
        </div>
      )}
    </div>
  );
}
