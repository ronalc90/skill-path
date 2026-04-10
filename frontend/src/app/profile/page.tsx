'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { api, UserPathProgress } from '@/lib/api';
import { useAuth } from '@/lib/auth-context';
import PathCard from '@/components/PathCard';

export default function ProfilePage() {
  const { user, isAuthenticated, isLoading: authLoading, refreshUser } = useAuth();
  const router = useRouter();
  const [myPaths, setMyPaths] = useState<UserPathProgress[]>([]);
  const [editing, setEditing] = useState(false);
  const [formData, setFormData] = useState({
    display_name: '',
    bio: '',
    github_url: '',
    linkedin_url: '',
  });
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    if (authLoading) return;
    if (!isAuthenticated) {
      router.push('/login');
      return;
    }

    api.getMyPaths().then(setMyPaths).catch(() => {});
  }, [isAuthenticated, authLoading, router]);

  useEffect(() => {
    if (user) {
      setFormData({
        display_name: user.display_name || '',
        bio: user.bio || '',
        github_url: user.github_url || '',
        linkedin_url: user.linkedin_url || '',
      });
    }
  }, [user]);

  const handleSave = async () => {
    setSaving(true);
    try {
      await api.updateProfile(formData);
      await refreshUser();
      setEditing(false);
    } catch {
      // Error
    } finally {
      setSaving(false);
    }
  };

  if (authLoading || !user) {
    return (
      <div className="max-w-4xl mx-auto px-4 py-12">
        <div className="animate-pulse space-y-4">
          <div className="h-24 bg-gray-200 rounded-xl" />
          <div className="h-8 bg-gray-200 rounded w-1/3" />
        </div>
      </div>
    );
  }

  const completedPaths = myPaths.filter((p) => p.completed_at);

  return (
    <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      {/* Profile header */}
      <div className="bg-white rounded-xl border border-gray-200 p-8 mb-8">
        <div className="flex items-start justify-between">
          <div className="flex items-center gap-6">
            <div className="w-20 h-20 bg-indigo-100 rounded-full flex items-center justify-center">
              <span className="text-2xl font-bold text-indigo-600">
                {user.display_name.charAt(0).toUpperCase()}
              </span>
            </div>
            <div>
              {editing ? (
                <div className="space-y-3">
                  <input
                    type="text"
                    value={formData.display_name}
                    onChange={(e) => setFormData({ ...formData, display_name: e.target.value })}
                    className="block w-full px-3 py-2 border border-gray-300 rounded-lg text-lg font-bold"
                  />
                  <textarea
                    value={formData.bio}
                    onChange={(e) => setFormData({ ...formData, bio: e.target.value })}
                    placeholder="Escribe algo sobre ti..."
                    rows={2}
                    className="block w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
                  />
                  <input
                    type="url"
                    value={formData.github_url}
                    onChange={(e) => setFormData({ ...formData, github_url: e.target.value })}
                    placeholder="https://github.com/username"
                    className="block w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
                  />
                  <input
                    type="url"
                    value={formData.linkedin_url}
                    onChange={(e) => setFormData({ ...formData, linkedin_url: e.target.value })}
                    placeholder="https://linkedin.com/in/username"
                    className="block w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
                  />
                </div>
              ) : (
                <>
                  <h1 className="text-2xl font-bold text-gray-900">{user.display_name}</h1>
                  <p className="text-gray-500">{user.email}</p>
                  {user.bio && <p className="text-sm text-gray-600 mt-2">{user.bio}</p>}
                  <div className="flex gap-3 mt-2">
                    {user.github_url && (
                      <a href={user.github_url} target="_blank" rel="noopener noreferrer" className="text-sm text-indigo-600 hover:text-indigo-700">
                        GitHub
                      </a>
                    )}
                    {user.linkedin_url && (
                      <a href={user.linkedin_url} target="_blank" rel="noopener noreferrer" className="text-sm text-indigo-600 hover:text-indigo-700">
                        LinkedIn
                      </a>
                    )}
                  </div>
                </>
              )}
            </div>
          </div>
          <div>
            {editing ? (
              <div className="flex gap-2">
                <button
                  onClick={handleSave}
                  disabled={saving}
                  className="px-4 py-2 bg-indigo-600 text-white text-sm rounded-lg hover:bg-indigo-700 disabled:opacity-50"
                >
                  {saving ? 'Guardando...' : 'Guardar'}
                </button>
                <button
                  onClick={() => setEditing(false)}
                  className="px-4 py-2 border border-gray-300 text-gray-700 text-sm rounded-lg hover:bg-gray-50"
                >
                  Cancelar
                </button>
              </div>
            ) : (
              <button
                onClick={() => setEditing(true)}
                className="px-4 py-2 border border-gray-300 text-gray-700 text-sm rounded-lg hover:bg-gray-50"
              >
                Editar perfil
              </button>
            )}
          </div>
        </div>

        {/* Stats */}
        <div className="grid grid-cols-3 gap-4 mt-6 pt-6 border-t border-gray-200">
          <div className="text-center">
            <p className="text-2xl font-bold text-gray-900">{myPaths.length}</p>
            <p className="text-sm text-gray-500">Rutas inscritas</p>
          </div>
          <div className="text-center">
            <p className="text-2xl font-bold text-gray-900">{completedPaths.length}</p>
            <p className="text-sm text-gray-500">Rutas completadas</p>
          </div>
          <div className="text-center">
            <p className="text-2xl font-bold text-gray-900">
              {myPaths.reduce((sum, p) => sum + Math.round(p.progress_pct), 0)}
            </p>
            <p className="text-sm text-gray-500">Puntos totales</p>
          </div>
        </div>
      </div>

      {/* Completed paths */}
      {completedPaths.length > 0 && (
        <div className="mb-8">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">Rutas completadas</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {completedPaths.map((p) => (
              <PathCard key={p.path_id} path={p.path} progress={100} />
            ))}
          </div>
        </div>
      )}

      {/* In-progress paths */}
      {myPaths.filter((p) => !p.completed_at).length > 0 && (
        <div>
          <h2 className="text-xl font-semibold text-gray-900 mb-4">En progreso</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {myPaths
              .filter((p) => !p.completed_at)
              .map((p) => (
                <PathCard key={p.path_id} path={p.path} progress={p.progress_pct} />
              ))}
          </div>
        </div>
      )}
    </div>
  );
}
