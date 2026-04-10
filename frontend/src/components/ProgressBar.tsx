interface ProgressBarProps {
  value: number;
  className?: string;
  showLabel?: boolean;
  size?: 'sm' | 'md' | 'lg';
}

export default function ProgressBar({ value, className = '', showLabel = true, size = 'md' }: ProgressBarProps) {
  const heights = { sm: 'h-1.5', md: 'h-2.5', lg: 'h-4' };

  return (
    <div className={className}>
      {showLabel && (
        <div className="flex justify-between mb-1">
          <span className="text-sm font-medium text-gray-700">Progreso</span>
          <span className="text-sm font-medium text-indigo-600">{Math.round(value)}%</span>
        </div>
      )}
      <div className={`w-full bg-gray-200 rounded-full ${heights[size]}`}>
        <div
          className={`bg-indigo-600 ${heights[size]} rounded-full transition-all duration-500 ease-out`}
          style={{ width: `${Math.min(value, 100)}%` }}
        />
      </div>
    </div>
  );
}
