'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { api, DashboardResponse } from '@/lib/api';
import { useAuth } from '@/lib/auth-context';
import PathCard from '@/components/PathCard';
import StatsCard from '@/components/StatsCard';
import SkillRadar from '@/components/SkillRadar';

export default function DashboardPage() {
  const { isAuthenticated, isLoading: authLoading } = useAuth();
  const router = useRouter();
  const [dashboard, setDashboard] = useState<DashboardResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (authLoading) return;
    if (!isAuthenticated) {
      router.push('/login');
      return;
    }

    api
      .getDashboard()
      .then(setDashboard)
      .catch(() => {})
      .finally(() => setLoading(false));
  }, [isAuthenticated, authLoading, router]);

  if (authLoading || loading) {
    return (
      <div className="max-w-7xl mx-auto px-4 py-12">
        <div className="animate-pulse space-y-6">
          <div className="h-8 bg-gray-200 rounded w-1/4" />
          <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
            {[1, 2, 3, 4].map((i) => (
              <div key={i} className="h-24 bg-gray-200 rounded-xl" />
            ))}
          </div>
        </div>
      </div>
    );
  }

  if (!dashboard) return null;

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <h1 className="text-3xl font-bold text-gray-900 mb-8">Mi Dashboard</h1>

      {/* Stats */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
        <StatsCard
          title="Rutas en progreso"
          value={dashboard.paths_in_progress}
          color="indigo"
          icon={
            <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
            </svg>
          }
        />
        <StatsCard
          title="Rutas completadas"
          value={dashboard.paths_completed}
          color="green"
          icon={
            <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          }
        />
        <StatsCard
          title="Horas aprendidas"
          value={Math.round(dashboard.total_hours_learned)}
          color="blue"
          icon={
            <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          }
        />
        <StatsCard
          title="Tasa de completado"
          value={`${Math.round(dashboard.completion_rate)}%`}
          color="purple"
          icon={
            <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z" />
            </svg>
          }
        />
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* Enrolled paths */}
        <div className="lg:col-span-2">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">Mis rutas</h2>
          {dashboard.enrolled_paths.length === 0 ? (
            <div className="bg-white rounded-xl border border-gray-200 p-12 text-center">
              <p className="text-gray-400 mb-4">No estas inscrito en ninguna ruta aun.</p>
              <a
                href="/paths"
                className="text-indigo-600 hover:text-indigo-700 font-medium"
              >
                Explorar rutas
              </a>
            </div>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {dashboard.enrolled_paths.map((ep) => (
                <PathCard
                  key={ep.path_id}
                  path={ep.path}
                  progress={ep.progress_pct}
                />
              ))}
            </div>
          )}
        </div>

        {/* Skill radar + recent activity */}
        <div className="space-y-6">
          <div className="bg-white rounded-xl border border-gray-200 p-6">
            <h2 className="text-lg font-semibold text-gray-900 mb-4">Radar de habilidades</h2>
            <SkillRadar data={dashboard.skill_radar} />
          </div>

          <div className="bg-white rounded-xl border border-gray-200 p-6">
            <h2 className="text-lg font-semibold text-gray-900 mb-4">Actividad reciente</h2>
            {dashboard.recent_activity.length === 0 ? (
              <p className="text-sm text-gray-400">Sin actividad reciente.</p>
            ) : (
              <div className="space-y-3">
                {dashboard.recent_activity.map((activity, i) => (
                  <div key={i} className="flex items-start gap-3">
                    <div className={`w-2 h-2 rounded-full mt-1.5 ${
                      activity.type === 'assessment' ? 'bg-purple-400' : 'bg-indigo-400'
                    }`} />
                    <div>
                      <p className="text-sm text-gray-700">{activity.title}</p>
                      <p className="text-xs text-gray-400">
                        {new Date(activity.timestamp).toLocaleDateString('es-ES', {
                          day: 'numeric',
                          month: 'short',
                          hour: '2-digit',
                          minute: '2-digit',
                        })}
                      </p>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
