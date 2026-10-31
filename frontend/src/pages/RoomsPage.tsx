import { useEffect, useState } from 'react';
import { roomsApi, type CreateRoomRequest } from '../api/rooms';
import type { Room, PaginationMetadata } from '../types';

export default function RoomsPage() {
  const [rooms, setRooms] = useState<Room[]>([]);
  const [meta, setMeta] = useState<PaginationMetadata | null>(null);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [editingId, setEditingId] = useState<string | null>(null);
  const [form, setForm] = useState<CreateRoomRequest>({
    name: '',
    capacity: 100,
    rows: 10,
    columns: 10,
  });

  const load = (p: number) => {
    setLoading(true);
    roomsApi.getAll(p, 10).then((res) => {
      setRooms(res.data.data);
      setMeta(res.data.metadata);
      setLoading(false);
    });
  };

  useEffect(() => load(page), [page]);

  const resetForm = () => {
    setForm({ name: '', capacity: 100, rows: 10, columns: 10 });
    setEditingId(null);
    setShowForm(false);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (editingId) {
      await roomsApi.update(editingId, form);
    } else {
      await roomsApi.create(form);
    }
    resetForm();
    load(page);
  };

  const handleEdit = (room: Room) => {
    setForm({ name: room.name, capacity: room.capacity, rows: room.rows, columns: room.columns });
    setEditingId(room.id);
    setShowForm(true);
  };

  const handleDelete = async (id: string) => {
    if (!confirm('Delete this room?')) return;
    await roomsApi.delete(id);
    load(page);
  };

  return (
    <div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">Rooms</h1>
        <button
          onClick={() => { resetForm(); setShowForm(!showForm); }}
          className="px-4 py-2 bg-indigo-600 hover:bg-indigo-500 text-white text-sm rounded cursor-pointer"
        >
          {showForm ? 'Cancel' : '+ New Room'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-gray-800 border border-gray-700 rounded-lg p-4 mb-6 space-y-3">
          <div className="grid grid-cols-2 sm:grid-cols-4 gap-3">
            <input
              placeholder="Name"
              value={form.name}
              onChange={(e) => setForm({ ...form, name: e.target.value })}
              required
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            />
            <input
              type="number"
              placeholder="Capacity"
              value={form.capacity}
              onChange={(e) => setForm({ ...form, capacity: +e.target.value })}
              required
              min={1}
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            />
            <input
              type="number"
              placeholder="Rows"
              value={form.rows}
              onChange={(e) => setForm({ ...form, rows: +e.target.value })}
              required
              min={1}
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            />
            <input
              type="number"
              placeholder="Columns"
              value={form.columns}
              onChange={(e) => setForm({ ...form, columns: +e.target.value })}
              required
              min={1}
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            />
          </div>
          <button
            type="submit"
            className="px-4 py-2 bg-green-600 hover:bg-green-500 text-white text-sm rounded cursor-pointer"
          >
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
                <th className="text-left py-2 px-3">Name</th>
                <th className="text-left py-2 px-3">Capacity</th>
                <th className="text-left py-2 px-3">Rows</th>
                <th className="text-left py-2 px-3">Columns</th>
                <th className="text-right py-2 px-3">Actions</th>
              </tr>
            </thead>
            <tbody>
              {rooms.map((r) => (
                <tr key={r.id} className="border-b border-gray-800">
                  <td className="py-2 px-3">{r.name}</td>
                  <td className="py-2 px-3">{r.capacity}</td>
                  <td className="py-2 px-3">{r.rows}</td>
                  <td className="py-2 px-3">{r.columns}</td>
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
