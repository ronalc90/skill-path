'use client';

import { SkillRadarPoint } from '@/lib/api';

interface SkillRadarProps {
  data: SkillRadarPoint[];
}

export default function SkillRadar({ data }: SkillRadarProps) {
  if (!data || data.length === 0) {
    return (
      <div className="flex items-center justify-center h-64 text-gray-400 text-sm">
        No hay datos de habilidades aun. Completa rutas para ver tu radar de skills.
      </div>
    );
  }

  const size = 250;
  const center = size / 2;
  const levels = 5;
  const radius = (size - 60) / 2;

  const angleStep = (2 * Math.PI) / data.length;

  const getPoint = (index: number, value: number, maxValue: number) => {
    const angle = angleStep * index - Math.PI / 2;
    const normalizedValue = maxValue > 0 ? value / maxValue : 0;
    const r = radius * normalizedValue;
    return {
      x: center + r * Math.cos(angle),
      y: center + r * Math.sin(angle),
    };
  };

  const gridLines = Array.from({ length: levels }, (_, i) => {
    const r = (radius * (i + 1)) / levels;
    const points = data.map((_, j) => {
      const angle = angleStep * j - Math.PI / 2;
      return `${center + r * Math.cos(angle)},${center + r * Math.sin(angle)}`;
    });
    return points.join(' ');
  });

  const dataPoints = data.map((point, i) => getPoint(i, point.score, point.max_score));
  const dataPolygon = dataPoints.map((p) => `${p.x},${p.y}`).join(' ');

  return (
    <div className="flex justify-center">
      <svg viewBox={`0 0 ${size} ${size}`} className="w-full max-w-[300px]">
        {/* Grid */}
        {gridLines.map((points, i) => (
          <polygon
            key={i}
            points={points}
            fill="none"
            stroke="#e5e7eb"
            strokeWidth="1"
          />
        ))}

        {/* Axis lines */}
        {data.map((_, i) => {
          const angle = angleStep * i - Math.PI / 2;
          return (
            <line
              key={i}
              x1={center}
              y1={center}
              x2={center + radius * Math.cos(angle)}
              y2={center + radius * Math.sin(angle)}
              stroke="#e5e7eb"
              strokeWidth="1"
            />
          );
        })}

        {/* Data polygon */}
        <polygon
          points={dataPolygon}
          fill="rgba(99, 102, 241, 0.2)"
          stroke="#6366f1"
          strokeWidth="2"
        />

        {/* Data points */}
        {dataPoints.map((point, i) => (
          <circle
            key={i}
            cx={point.x}
            cy={point.y}
            r="4"
            fill="#6366f1"
            stroke="white"
            strokeWidth="2"
          />
        ))}

        {/* Labels */}
        {data.map((point, i) => {
          const angle = angleStep * i - Math.PI / 2;
          const labelRadius = radius + 25;
          const x = center + labelRadius * Math.cos(angle);
          const y = center + labelRadius * Math.sin(angle);
          return (
            <text
              key={i}
              x={x}
              y={y}
              textAnchor="middle"
              dominantBaseline="middle"
              className="text-[10px] fill-gray-600 font-medium"
            >
              {point.category}
            </text>
          );
        })}
      </svg>
    </div>
  );
}
