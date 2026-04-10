'use client';

import { AssessmentQuestion, AnswerDetail } from '@/lib/api';

interface QuizQuestionProps {
  question: AssessmentQuestion;
  questionIndex: number;
  selectedAnswer: number | undefined;
  onAnswer: (questionId: number, selectedIndex: number) => void;
  result?: AnswerDetail;
  showResult: boolean;
}

export default function QuizQuestion({
  question,
  questionIndex,
  selectedAnswer,
  onAnswer,
  result,
  showResult,
}: QuizQuestionProps) {
  return (
    <div className="bg-white rounded-xl border border-gray-200 p-6">
      <h3 className="font-medium text-gray-900 mb-4">
        <span className="text-indigo-600 font-bold mr-2">#{questionIndex + 1}</span>
        {question.question}
      </h3>

      <div className="space-y-3">
        {question.options.map((option, index) => {
          const isSelected = selectedAnswer === index;
          let optionClass = 'border-gray-200 hover:border-indigo-300';

          if (showResult && result) {
            if (index === result.correct_answer) {
              optionClass = 'border-green-500 bg-green-50';
            } else if (isSelected && !result.is_correct) {
              optionClass = 'border-red-500 bg-red-50';
            }
          } else if (isSelected) {
            optionClass = 'border-indigo-500 bg-indigo-50';
          }

          return (
            <button
              key={index}
              onClick={() => !showResult && onAnswer(question.id, index)}
              disabled={showResult}
              className={`w-full text-left p-4 rounded-lg border-2 transition-all ${optionClass} ${
                showResult ? 'cursor-default' : 'cursor-pointer'
              }`}
            >
              <div className="flex items-center gap-3">
                <div
                  className={`w-6 h-6 rounded-full flex items-center justify-center text-xs font-bold ${
                    isSelected && !showResult
                      ? 'bg-indigo-600 text-white'
                      : showResult && result && index === result.correct_answer
                      ? 'bg-green-500 text-white'
                      : showResult && isSelected && !result?.is_correct
                      ? 'bg-red-500 text-white'
                      : 'bg-gray-100 text-gray-600'
                  }`}
                >
                  {String.fromCharCode(65 + index)}
                </div>
                <span className="text-sm text-gray-700">{option}</span>
              </div>
            </button>
          );
        })}
      </div>

      {showResult && result && (
        <div
          className={`mt-4 p-4 rounded-lg text-sm ${
            result.is_correct ? 'bg-green-50 text-green-800' : 'bg-orange-50 text-orange-800'
          }`}
        >
          <p className="font-medium mb-1">
            {result.is_correct ? 'Correcto!' : 'Incorrecto'}
          </p>
          <p>{result.explanation}</p>
        </div>
      )}
    </div>
  );
}
