import { useEffect, useState } from 'react';
import { seatTypesApi, type CreateSeatTypeRequest } from '../api/seatTypes';
import type { SeatType } from '../types';

export default function SeatTypesPage() {
  const [types, setTypes] = useState<SeatType[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [editingId, setEditingId] = useState<string | null>(null);
  const [form, setForm] = useState<CreateSeatTypeRequest>({ name: '', multiplier: 1.0 });

  const load = () => {
    setLoading(true);
    seatTypesApi.getAll().then((res) => {
      setTypes(res.data);
      setLoading(false);
    });
  };

  useEffect(() => load(), []);

  const resetForm = () => {
    setForm({ name: '', multiplier: 1.0 });
    setEditingId(null);
    setShowForm(false);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (editingId) {
      await seatTypesApi.update(editingId, form);
    } else {
      await seatTypesApi.create(form);
    }
    resetForm();
    load();
  };

  const handleEdit = (st: SeatType) => {
    setForm({ name: st.name, multiplier: st.multiplier });
    setEditingId(st.id);
    setShowForm(true);
  };

  const handleDelete = async (id: string) => {
    if (!confirm('Delete this seat type?')) return;
    await seatTypesApi.delete(id);
    load();
  };

  return (
    <div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">Seat Types</h1>
        <button
          onClick={() => { resetForm(); setShowForm(!showForm); }}
          className="px-4 py-2 bg-indigo-600 hover:bg-indigo-500 text-white text-sm rounded cursor-pointer"
        >
          {showForm ? 'Cancel' : '+ New Seat Type'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-gray-800 border border-gray-700 rounded-lg p-4 mb-6 space-y-3">
          <div className="grid grid-cols-2 gap-3">
            <input
              placeholder="Name (e.g. VIP)"
              value={form.name}
              onChange={(e) => setForm({ ...form, name: e.target.value })}
              required
              className="bg-gray-900 border border-gray-600 rounded px-3 py-2 text-sm text-white"
            />
            <input
              type="number"
              step="0.01"
              placeholder="Multiplier"
              value={form.multiplier}
              onChange={(e) => setForm({ ...form, multiplier: +e.target.value })}
              required
              min={0.01}
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
                <th className="text-left py-2 px-3">Name</th>
                <th className="text-left py-2 px-3">Multiplier</th>
                <th className="text-right py-2 px-3">Actions</th>
              </tr>
            </thead>
            <tbody>
              {types.map((st) => (
                <tr key={st.id} className="border-b border-gray-800">
                  <td className="py-2 px-3">{st.name}</td>
                  <td className="py-2 px-3">{st.multiplier}x</td>
                  <td className="py-2 px-3 text-right space-x-2">
                    <button onClick={() => handleEdit(st)} className="text-indigo-400 hover:text-indigo-300 text-xs cursor-pointer">Edit</button>
                    <button onClick={() => handleDelete(st.id)} className="text-red-400 hover:text-red-300 text-xs cursor-pointer">Delete</button>
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
