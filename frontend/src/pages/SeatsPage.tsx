import { useEffect, useState } from 'react';
import { seatsApi, type CreateSeatRequest } from '../api/seats';
import { seatTypesApi } from '../api/seatTypes';
import { roomsApi } from '../api/rooms';
import type { Seat, SeatType, Room } from '../types';

export default function SeatsPage() {
  const [seats, setSeats] = useState<Seat[]>([]);
  const [seatTypes, setSeatTypes] = useState<SeatType[]>([]);
  const [rooms, setRooms] = useState<Room[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [editingId, setEditingId] = useState<string | null>(null);
  const [form, setForm] = useState<CreateSeatRequest>({
    room_id: '',
    seat_type_id: '',
    row: 1,
    number: 1,
  });

  const load = () => {
    setLoading(true);
    Promise.all([seatsApi.getAll(), seatTypesApi.getAll(), roomsApi.getAll(1, 100)])
      .then(([s, st, r]) => {
        setSeats(s.data);
        setSeatTypes(st.data);
        setRooms(r.data.data);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => load(), []);

  const resetForm = () => {
    setForm({ room_id: '', seat_type_id: '', row: 1, number: 1 });
    setEditingId(null);
    setShowForm(false);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (editingId) {
      await seatsApi.update(editingId, form);
    } else {
      await seatsApi.create(form);
    }
    resetForm();
    load();
  };

  const handleEdit = (seat: Seat) => {
    setForm({ room_id: seat.room_id, seat_type_id: seat.seat_type_id, row: seat.row, number: seat.number });
    setEditingId(seat.id);
    setShowForm(true);
  };

  const handleDelete = async (id: string) => {
    if (!confirm('Delete this seat?')) return;
    await seatsApi.delete(id);
    load();
  };

  const getRoomName = (id: string) => rooms.find((r) => r.id === id)?.name ?? id.slice(0, 8);
  const getTypeName = (id: string) => seatTypes.find((st) => st.id === id)?.name ?? id.slice(0, 8);

  return (
    <div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">Seats</h1>
        <button
          onClick={() => { resetForm(); setShowForm(!showForm); }}
          className="px-4 py-2 bg-indigo-600 hover:bg-indigo-500 text-white text-sm rounded cursor-pointer"
        >
          {showForm ? 'Cancel' : '+ New Seat'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-gray-800 border border-gray-700 rounded-lg p-4 mb-6 space-y-3">
          <div className="grid grid-cols-2 sm:grid-cols-4 gap-3">
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
              value={form.seat_type_id}
              onChange={(e) => setForm({ ...form, seat_type_id: e.target.value })}
              required
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            >
              <option value="">Select Type</option>
              {seatTypes.map((st) => (
                <option key={st.id} value={st.id}>{st.name}</option>
              ))}
            </select>
            <input
              type="number"
              placeholder="Row"
              value={form.row}
              onChange={(e) => setForm({ ...form, row: +e.target.value })}
              required
              min={1}
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            />
            <input
              type="number"
              placeholder="Number"
              value={form.number}
              onChange={(e) => setForm({ ...form, number: +e.target.value })}
              required
              min={1}
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
                <th className="text-left py-2 px-3">Room</th>
                <th className="text-left py-2 px-3">Type</th>
                <th className="text-left py-2 px-3">Row</th>
                <th className="text-left py-2 px-3">Number</th>
                <th className="text-right py-2 px-3">Actions</th>
              </tr>
            </thead>
            <tbody>
              {seats.map((s) => (
                <tr key={s.id} className="border-b border-gray-800">
                  <td className="py-2 px-3">{getRoomName(s.room_id)}</td>
                  <td className="py-2 px-3">{getTypeName(s.seat_type_id)}</td>
                  <td className="py-2 px-3">{s.row}</td>
                  <td className="py-2 px-3">{s.number}</td>
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
