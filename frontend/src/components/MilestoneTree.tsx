import { MilestoneResponse } from '@/lib/api';

interface MilestoneTreeProps {
  milestones: MilestoneResponse[];
  completedMilestones?: Set<number>;
  currentMilestoneId?: number;
  onMilestoneClick?: (milestone: MilestoneResponse) => void;
}

export default function MilestoneTree({
  milestones,
  completedMilestones = new Set(),
  currentMilestoneId,
  onMilestoneClick,
}: MilestoneTreeProps) {
  const sorted = [...milestones].sort((a, b) => a.order_index - b.order_index);

  return (
    <div className="relative">
      {sorted.map((milestone, index) => {
        const isCompleted = completedMilestones.has(milestone.id);
        const isCurrent = milestone.id === currentMilestoneId;
        const isLast = index === sorted.length - 1;

        return (
          <div key={milestone.id} className="relative flex gap-4">
            {/* Timeline line */}
            <div className="flex flex-col items-center">
              <div
                className={`w-8 h-8 rounded-full flex items-center justify-center z-10 border-2 ${
                  isCompleted
                    ? 'bg-green-500 border-green-500 text-white'
                    : isCurrent
                    ? 'bg-indigo-500 border-indigo-500 text-white'
                    : 'bg-white border-gray-300 text-gray-400'
                }`}
              >
                {isCompleted ? (
                  <svg className="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                    <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                  </svg>
                ) : (
                  <span className="text-xs font-bold">{index + 1}</span>
                )}
              </div>
              {!isLast && (
                <div className={`w-0.5 flex-1 min-h-[2rem] ${isCompleted ? 'bg-green-300' : 'bg-gray-200'}`} />
              )}
            </div>

            {/* Content */}
            <div
              className={`flex-1 pb-6 ${onMilestoneClick ? 'cursor-pointer' : ''}`}
              onClick={() => onMilestoneClick?.(milestone)}
            >
              <div
                className={`p-4 rounded-lg border transition-all ${
                  isCurrent
                    ? 'border-indigo-300 bg-indigo-50'
                    : isCompleted
                    ? 'border-green-200 bg-green-50'
                    : 'border-gray-200 bg-white hover:border-gray-300'
                }`}
              >
                <div className="flex items-center justify-between mb-1">
                  <h4 className="font-medium text-gray-900">{milestone.title}</h4>
                  <span className="text-xs text-gray-400">{milestone.estimated_hours}h</span>
                </div>
                <p className="text-sm text-gray-500">{milestone.description}</p>
                {milestone.resources && milestone.resources.length > 0 && (
                  <div className="mt-2 flex items-center gap-2 text-xs text-gray-400">
                    <svg className="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
                    </svg>
                    {milestone.resources.length} recursos
                  </div>
                )}
              </div>
            </div>
          </div>
        );
      })}
    </div>
  );
}
