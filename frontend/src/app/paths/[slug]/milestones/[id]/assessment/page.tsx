'use client';

import { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import Link from 'next/link';
import { api, AssessmentQuestion, AssessmentResult, AnswerDetail } from '@/lib/api';
import { useAuth } from '@/lib/auth-context';
import QuizQuestion from '@/components/QuizQuestion';

export default function AssessmentPage() {
  const params = useParams();
  const router = useRouter();
  const { isAuthenticated } = useAuth();
  const [questions, setQuestions] = useState<AssessmentQuestion[]>([]);
  const [answers, setAnswers] = useState<Record<number, number>>({});
  const [result, setResult] = useState<AssessmentResult | null>(null);
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState('');

  const milestoneId = parseInt(params.id as string);

  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/login');
      return;
    }

    api
      .getAssessment(milestoneId)
      .then(setQuestions)
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false));
  }, [milestoneId, isAuthenticated, router]);

  const handleAnswer = (questionId: number, selectedIndex: number) => {
    setAnswers((prev) => ({ ...prev, [questionId]: selectedIndex }));
  };

  const handleSubmit = async () => {
    if (Object.keys(answers).length < questions.length) {
      setError('Responde todas las preguntas antes de enviar.');
      return;
    }

    setSubmitting(true);
    setError('');

    try {
      const submission = Object.entries(answers).map(([qid, idx]) => ({
        question_id: parseInt(qid),
        selected_index: idx,
      }));
      const res = await api.submitAssessment(milestoneId, submission);
      setResult(res);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Error al enviar evaluacion');
    } finally {
      setSubmitting(false);
    }
  };

  const getResultForQuestion = (questionId: number): AnswerDetail | undefined => {
    return result?.details.find((d) => d.question_id === questionId);
  };

  if (loading) {
    return (
      <div className="max-w-3xl mx-auto px-4 py-12">
        <div className="animate-pulse space-y-6">
          {[1, 2, 3].map((i) => (
            <div key={i} className="bg-white rounded-xl border border-gray-200 p-6">
              <div className="h-5 bg-gray-200 rounded w-3/4 mb-4" />
              <div className="space-y-3">
                {[1, 2, 3, 4].map((j) => (
                  <div key={j} className="h-12 bg-gray-200 rounded" />
                ))}
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  }

  return (
    <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <Link
        href={`/paths/${params.slug}/learn`}
        className="text-sm text-indigo-600 hover:text-indigo-700 mb-4 inline-flex items-center gap-1"
      >
        <svg className="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
        </svg>
        Volver al aprendizaje
      </Link>

      <h1 className="text-2xl font-bold text-gray-900 mb-2">Evaluacion</h1>
      <p className="text-gray-500 mb-8">
        Responde las siguientes preguntas para validar tu conocimiento. Necesitas un 70% para aprobar.
      </p>

      {error && (
        <div className="p-4 bg-red-50 border border-red-200 rounded-lg text-sm text-red-600 mb-6">
          {error}
        </div>
      )}

      {/* Result banner */}
      {result && (
        <div
          className={`p-6 rounded-xl border mb-8 ${
            result.passed
              ? 'bg-green-50 border-green-200'
              : 'bg-orange-50 border-orange-200'
          }`}
        >
          <div className="flex items-center justify-between">
            <div>
              <h2 className={`text-xl font-bold ${result.passed ? 'text-green-800' : 'text-orange-800'}`}>
                {result.passed ? 'Aprobado!' : 'No aprobado'}
              </h2>
              <p className={`text-sm mt-1 ${result.passed ? 'text-green-600' : 'text-orange-600'}`}>
                Obtuviste {result.score} de {result.total_questions} respuestas correctas
                ({Math.round((result.score / result.total_questions) * 100)}%)
              </p>
            </div>
            <div className="text-4xl font-bold">
              <span className={result.passed ? 'text-green-600' : 'text-orange-600'}>
                {result.score}/{result.total_questions}
              </span>
            </div>
          </div>
        </div>
      )}

      {/* Questions */}
      <div className="space-y-6 mb-8">
        {questions.map((question, index) => (
          <QuizQuestion
            key={question.id}
            question={question}
            questionIndex={index}
            selectedAnswer={answers[question.id]}
            onAnswer={handleAnswer}
            result={getResultForQuestion(question.id)}
            showResult={!!result}
          />
        ))}
      </div>

      {/* Submit */}
      {!result && questions.length > 0 && (
        <div className="flex justify-end">
          <button
            onClick={handleSubmit}
            disabled={submitting}
            className="px-8 py-3 bg-indigo-600 text-white font-medium rounded-lg hover:bg-indigo-700 disabled:opacity-50 transition-colors"
          >
            {submitting ? 'Enviando...' : 'Enviar respuestas'}
          </button>
        </div>
      )}

      {result && (
        <div className="flex justify-center gap-4">
          <Link
            href={`/paths/${params.slug}/learn`}
            className="px-6 py-3 bg-indigo-600 text-white font-medium rounded-lg hover:bg-indigo-700 transition-colors"
          >
            Continuar aprendiendo
          </Link>
          {!result.passed && (
            <button
              onClick={() => {
                setResult(null);
                setAnswers({});
                setError('');
              }}
              className="px-6 py-3 border border-gray-300 text-gray-700 font-medium rounded-lg hover:bg-gray-50 transition-colors"
            >
              Intentar de nuevo
            </button>
          )}
        </div>
      )}
    </div>
  );
}
