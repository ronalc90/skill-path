'use client';

import { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { api, PathDetailResponse } from '@/lib/api';
import { useAuth } from '@/lib/auth-context';
import MilestoneTree from '@/components/MilestoneTree';
import ResourceItem from '@/components/ResourceItem';

const difficultyLabels: Record<string, string> = {
  beginner: 'Principiante',
  intermediate: 'Intermedio',
  advanced: 'Avanzado',
};

export default function PathDetailPage() {
  const params = useParams();
  const router = useRouter();
  const { isAuthenticated } = useAuth();
  const [path, setPath] = useState<PathDetailResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [enrolling, setEnrolling] = useState(false);
  const [selectedMilestone, setSelectedMilestone] = useState<number | null>(null);

  useEffect(() => {
    if (params.slug) {
      api
        .getPathBySlug(params.slug as string)
        .then(setPath)
        .catch(() => router.push('/paths'))
        .finally(() => setLoading(false));
    }
  }, [params.slug, router]);

  const handleStart = async () => {
    if (!isAuthenticated) {
      router.push('/login');
      return;
    }
    if (!path) return;

    setEnrolling(true);
    try {
      await api.startPath(path.id);
      router.push(`/paths/${path.slug}/learn`);
    } catch (err) {
      if (err instanceof Error && err.message.includes('already')) {
        router.push(`/paths/${path.slug}/learn`);
      }
    } finally {
      setEnrolling(false);
    }
  };

  if (loading) {
    return (
      <div className="max-w-4xl mx-auto px-4 py-12">
        <div className="animate-pulse space-y-4">
          <div className="h-8 bg-gray-200 rounded w-1/2" />
          <div className="h-4 bg-gray-200 rounded w-3/4" />
          <div className="h-4 bg-gray-200 rounded w-2/3" />
        </div>
      </div>
    );
  }

  if (!path) return null;

  const activeMilestone = selectedMilestone !== null
    ? path.milestones.find((m) => m.id === selectedMilestone)
    : null;

  return (
    <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      {/* Header */}
      <div className="bg-white rounded-xl border border-gray-200 p-8 mb-8">
        <div className="flex items-start gap-6">
          {path.icon_url && <img src={path.icon_url} alt="" className="w-16 h-16" />}
          <div className="flex-1">
            <div className="flex items-center gap-3 mb-2">
              <span className="text-xs font-medium px-2.5 py-1 rounded-full bg-indigo-100 text-indigo-700">
                {path.category}
              </span>
              <span className="text-xs font-medium px-2.5 py-1 rounded-full bg-gray-100 text-gray-600">
                {difficultyLabels[path.difficulty] || path.difficulty}
              </span>
            </div>
            <h1 className="text-3xl font-bold text-gray-900 mb-3">{path.title}</h1>
            <p className="text-gray-500 mb-4">{path.description}</p>
            <div className="flex items-center gap-6 text-sm text-gray-400">
              <span>{path.estimated_hours} horas estimadas</span>
              <span>{path.milestones.length} hitos</span>
              <span>
                {path.milestones.reduce((acc, m) => acc + m.resources.length, 0)} recursos
              </span>
            </div>
          </div>
          <button
            onClick={handleStart}
            disabled={enrolling}
            className="px-6 py-3 bg-indigo-600 text-white font-medium rounded-lg hover:bg-indigo-700 disabled:opacity-50 transition-colors whitespace-nowrap"
          >
            {enrolling ? 'Iniciando...' : 'Comenzar ruta'}
          </button>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* Milestone Tree */}
        <div className="lg:col-span-1">
          <h2 className="text-lg font-semibold text-gray-900 mb-4">Hitos</h2>
          <MilestoneTree
            milestones={path.milestones}
            currentMilestoneId={selectedMilestone || undefined}
            onMilestoneClick={(m) => setSelectedMilestone(m.id === selectedMilestone ? null : m.id)}
          />
        </div>

        {/* Resources panel */}
        <div className="lg:col-span-2">
          {activeMilestone ? (
            <div>
              <h2 className="text-lg font-semibold text-gray-900 mb-2">{activeMilestone.title}</h2>
              <p className="text-sm text-gray-500 mb-4">{activeMilestone.description}</p>
              <div className="space-y-3">
                {activeMilestone.resources.map((resource) => (
                  <ResourceItem key={resource.id} resource={resource} />
                ))}
              </div>
            </div>
          ) : (
            <div className="text-center py-16 text-gray-400">
              <svg className="w-12 h-12 mx-auto mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M15 15l-2 5L9 9l11 4-5 2zm0 0l5 5M7.188 2.239l.777 2.897M5.136 7.965l-2.898-.777M13.95 4.05l-2.122 2.122m-5.657 5.656l-2.12 2.122" />
              </svg>
              <p>Selecciona un hito para ver sus recursos</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
