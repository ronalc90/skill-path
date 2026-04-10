import { ResourceResponse } from '@/lib/api';

const typeIcons: Record<string, string> = {
  video: 'M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z',
  article: 'M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9M7 16h6M7 8h6v4H7V8z',
  course: 'M12 14l9-5-9-5-9 5 9 5z M12 14l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14z',
  book: 'M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253',
  tutorial: 'M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z',
  exercise: 'M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4',
};

const typeLabels: Record<string, string> = {
  video: 'Video',
  article: 'Articulo',
  course: 'Curso',
  book: 'Libro',
  tutorial: 'Tutorial',
  exercise: 'Ejercicio',
};

interface ResourceItemProps {
  resource: ResourceResponse;
  isCompleted?: boolean;
  onComplete?: () => void;
}

export default function ResourceItem({ resource, isCompleted, onComplete }: ResourceItemProps) {
  return (
    <div className={`flex items-center gap-3 p-3 rounded-lg border transition-all ${
      isCompleted ? 'bg-green-50 border-green-200' : 'bg-white border-gray-200 hover:border-gray-300'
    }`}>
      {onComplete && (
        <button
          onClick={(e) => {
            e.preventDefault();
            onComplete();
          }}
          className={`flex-shrink-0 w-5 h-5 rounded border-2 flex items-center justify-center transition-all ${
            isCompleted
              ? 'bg-green-500 border-green-500 text-white'
              : 'border-gray-300 hover:border-indigo-400'
          }`}
        >
          {isCompleted && (
            <svg className="w-3 h-3" fill="currentColor" viewBox="0 0 20 20">
              <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
            </svg>
          )}
        </button>
      )}

      <div className="flex-shrink-0">
        <svg className="w-5 h-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d={typeIcons[resource.type] || typeIcons.article} />
        </svg>
      </div>

      <div className="flex-1 min-w-0">
        <a
          href={resource.url}
          target="_blank"
          rel="noopener noreferrer"
          className="text-sm font-medium text-gray-900 hover:text-indigo-600 truncate block"
        >
          {resource.title}
        </a>
        <div className="flex items-center gap-2 mt-0.5">
          <span className="text-xs px-1.5 py-0.5 bg-gray-100 rounded text-gray-500">
            {typeLabels[resource.type] || resource.type}
          </span>
          {resource.provider && (
            <span className="text-xs text-gray-400">{resource.provider}</span>
          )}
          <span className="text-xs text-gray-400">{resource.estimated_minutes} min</span>
          {resource.is_free ? (
            <span className="text-xs text-green-600 font-medium">Gratis</span>
          ) : (
            <span className="text-xs text-orange-500 font-medium">Pago</span>
          )}
        </div>
      </div>

      <a
        href={resource.url}
        target="_blank"
        rel="noopener noreferrer"
        className="flex-shrink-0 text-gray-400 hover:text-indigo-600"
      >
        <svg className="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
        </svg>
      </a>
    </div>
  );
}
