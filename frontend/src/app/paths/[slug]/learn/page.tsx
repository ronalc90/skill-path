'use client';

import { useEffect, useState, useCallback } from 'react';
import { useParams, useRouter } from 'next/navigation';
import Link from 'next/link';
import { api, PathDetailResponse, UserPathProgress, MilestoneResponse } from '@/lib/api';
import { useAuth } from '@/lib/auth-context';
import ProgressBar from '@/components/ProgressBar';
import MilestoneTree from '@/components/MilestoneTree';
import ResourceItem from '@/components/ResourceItem';

export default function LearnPage() {
  const params = useParams();
  const router = useRouter();
  const { isAuthenticated } = useAuth();
  const [path, setPath] = useState<PathDetailResponse | null>(null);
  const [progress, setProgress] = useState<UserPathProgress | null>(null);
  const [currentMilestone, setCurrentMilestone] = useState<MilestoneResponse | null>(null);
  const [completedResources, setCompletedResources] = useState<Set<number>>(new Set());
  const [loading, setLoading] = useState(true);

  const completedMilestoneIds = new Set(
    progress?.milestones?.filter((m) => m.status === 'completed').map((m) => m.milestone_id) || []
  );

  const fetchProgress = useCallback(async (pathId: number) => {
    try {
      const prog = await api.getPathProgress(pathId);
      setProgress(prog);
    } catch {
      // Not enrolled yet
    }
  }, []);

  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/login');
      return;
    }

    const load = async () => {
      try {
        const pathData = await api.getPathBySlug(params.slug as string);
        setPath(pathData);
        await fetchProgress(pathData.id);

        if (pathData.milestones.length > 0) {
          setCurrentMilestone(pathData.milestones[0]);
        }
      } catch {
        router.push('/paths');
      } finally {
        setLoading(false);
      }
    };

    load();
  }, [params.slug, isAuthenticated, router, fetchProgress]);

  const handleCompleteMilestone = async (milestoneId: number) => {
    if (!path) return;
    try {
      await api.completeMilestone(path.id, milestoneId);
      await fetchProgress(path.id);
    } catch {
      // Error handling
    }
  };

  const handleCompleteResource = async (resourceId: number) => {
    if (!path) return;
    try {
      await api.completeResource(path.id, resourceId);
      setCompletedResources((prev) => new Set([...Array.from(prev), resourceId]));
    } catch {
      // Error handling
    }
  };

  if (loading || !path) {
    return (
      <div className="max-w-6xl mx-auto px-4 py-12">
        <div className="animate-pulse space-y-4">
          <div className="h-8 bg-gray-200 rounded w-1/3" />
          <div className="h-4 bg-gray-200 rounded w-full" />
        </div>
      </div>
    );
  }

  return (
    <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      {/* Header */}
      <div className="bg-white rounded-xl border border-gray-200 p-6 mb-6">
        <div className="flex items-center justify-between mb-4">
          <div>
            <Link href={`/paths/${path.slug}`} className="text-sm text-indigo-600 hover:text-indigo-700 mb-1 inline-block">
              Volver a la ruta
            </Link>
            <h1 className="text-2xl font-bold text-gray-900">{path.title}</h1>
          </div>
          {progress && progress.progress_pct === 100 && (
            <span className="px-4 py-2 bg-green-100 text-green-700 font-medium rounded-lg text-sm">
              Ruta completada!
            </span>
          )}
        </div>
        <ProgressBar value={progress?.progress_pct || 0} />
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Milestones sidebar */}
        <div className="lg:col-span-1">
          <div className="bg-white rounded-xl border border-gray-200 p-4">
            <h2 className="font-semibold text-gray-900 mb-4">Progreso</h2>
            <MilestoneTree
              milestones={path.milestones}
              completedMilestones={completedMilestoneIds}
              currentMilestoneId={currentMilestone?.id}
              onMilestoneClick={(m) => setCurrentMilestone(m)}
            />
          </div>
        </div>

        {/* Active milestone content */}
        <div className="lg:col-span-2 space-y-6">
          {currentMilestone && (
            <>
              <div className="bg-white rounded-xl border border-gray-200 p-6">
                <div className="flex items-center justify-between mb-4">
                  <div>
                    <h2 className="text-xl font-bold text-gray-900">{currentMilestone.title}</h2>
                    <p className="text-sm text-gray-500 mt-1">{currentMilestone.description}</p>
                  </div>
                  {!completedMilestoneIds.has(currentMilestone.id) && (
                    <button
                      onClick={() => handleCompleteMilestone(currentMilestone.id)}
                      className="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 transition-colors whitespace-nowrap"
                    >
                      Completar hito
                    </button>
                  )}
                </div>

                {/* Resources */}
                <h3 className="font-medium text-gray-700 mb-3">Recursos ({currentMilestone.resources.length})</h3>
                <div className="space-y-2">
                  {currentMilestone.resources.map((resource) => (
                    <ResourceItem
                      key={resource.id}
                      resource={resource}
                      isCompleted={completedResources.has(resource.id)}
                      onComplete={() => handleCompleteResource(resource.id)}
                    />
                  ))}
                </div>
              </div>

              {/* Assessment link */}
              <div className="bg-indigo-50 rounded-xl border border-indigo-200 p-6">
                <div className="flex items-center justify-between">
                  <div>
                    <h3 className="font-semibold text-indigo-900">Evaluacion del hito</h3>
                    <p className="text-sm text-indigo-700 mt-1">
                      Pon a prueba tus conocimientos con una evaluacion rapida.
                    </p>
                  </div>
                  <Link
                    href={`/paths/${path.slug}/milestones/${currentMilestone.id}/assessment`}
                    className="px-4 py-2 bg-indigo-600 text-white text-sm font-medium rounded-lg hover:bg-indigo-700 transition-colors"
                  >
                    Tomar evaluacion
                  </Link>
                </div>
              </div>
            </>
          )}
        </div>
      </div>
    </div>
  );
}
