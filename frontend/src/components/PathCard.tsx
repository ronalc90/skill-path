import Link from 'next/link';
import { PathSummary } from '@/lib/api';

const difficultyColors: Record<string, string> = {
  beginner: 'bg-green-100 text-green-800',
  intermediate: 'bg-yellow-100 text-yellow-800',
  advanced: 'bg-red-100 text-red-800',
};

const difficultyLabels: Record<string, string> = {
  beginner: 'Principiante',
  intermediate: 'Intermedio',
  advanced: 'Avanzado',
};

interface PathCardProps {
  path: PathSummary;
  progress?: number;
}

export default function PathCard({ path, progress }: PathCardProps) {
  return (
    <Link href={`/paths/${path.slug}`}>
      <div className="bg-white rounded-xl border border-gray-200 hover:border-indigo-300 hover:shadow-lg transition-all duration-200 overflow-hidden h-full flex flex-col">
        <div className="p-6 flex-1">
          <div className="flex items-start justify-between mb-3">
            {path.icon_url && (
              <img src={path.icon_url} alt="" className="w-10 h-10" />
            )}
            <span className={`text-xs font-medium px-2.5 py-1 rounded-full ${difficultyColors[path.difficulty] || 'bg-gray-100 text-gray-800'}`}>
              {difficultyLabels[path.difficulty] || path.difficulty}
            </span>
          </div>
          <h3 className="font-semibold text-lg text-gray-900 mb-2">{path.title}</h3>
          <p className="text-sm text-gray-500 line-clamp-2 mb-4">{path.description}</p>
          <div className="flex items-center gap-4 text-xs text-gray-400">
            <span className="flex items-center gap-1">
              <svg className="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              {path.estimated_hours}h
            </span>
            <span className="flex items-center gap-1">
              <svg className="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
              {path.milestone_count} hitos
            </span>
            <span className="px-2 py-0.5 bg-gray-100 rounded text-gray-500">{path.category}</span>
          </div>
        </div>
        {typeof progress === 'number' && (
          <div className="px-6 pb-4">
            <div className="flex items-center justify-between mb-1">
              <span className="text-xs font-medium text-indigo-600">{Math.round(progress)}% completado</span>
            </div>
            <div className="w-full bg-gray-200 rounded-full h-2">
              <div
                className="bg-indigo-600 h-2 rounded-full transition-all duration-300"
                style={{ width: `${progress}%` }}
              />
            </div>
          </div>
        )}
      </div>
    </Link>
  );
}
