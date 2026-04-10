'use client';

import Link from 'next/link';
import { useEffect, useState } from 'react';
import { api, PathSummary } from '@/lib/api';
import PathCard from '@/components/PathCard';

export default function Home() {
  const [featuredPaths, setFeaturedPaths] = useState<PathSummary[]>([]);

  useEffect(() => {
    api.getPaths({ page: 1 }).then((res) => setFeaturedPaths(res.paths)).catch(() => {});
  }, []);

  return (
    <div>
      {/* Hero */}
      <section className="bg-gradient-to-br from-indigo-600 via-indigo-700 to-purple-800 text-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20 md:py-32">
          <div className="text-center max-w-3xl mx-auto">
            <h1 className="text-4xl md:text-6xl font-bold mb-6 leading-tight">
              Tu camino hacia la <span className="text-indigo-200">excelencia tech</span>
            </h1>
            <p className="text-lg md:text-xl text-indigo-100 mb-8 leading-relaxed">
              Rutas de aprendizaje curadas por expertos. Aprende Backend, Frontend, Data Science
              y Mobile Development con recursos gratuitos y de calidad.
            </p>
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <Link href="/paths" className="px-8 py-3.5 bg-white text-indigo-700 font-semibold rounded-lg hover:bg-indigo-50 transition-colors">
                Explorar rutas
              </Link>
              <Link href="/register" className="px-8 py-3.5 border-2 border-white text-white font-semibold rounded-lg hover:bg-white/10 transition-colors">
                Comenzar gratis
              </Link>
            </div>
          </div>
        </div>
      </section>

      {/* Stats */}
      <section className="bg-white border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8 text-center">
            {[
              { value: '4+', label: 'Rutas completas' },
              { value: '27+', label: 'Hitos de aprendizaje' },
              { value: '80+', label: 'Recursos curados' },
              { value: '395+', label: 'Horas de contenido' },
            ].map((stat) => (
              <div key={stat.label}>
                <p className="text-3xl font-bold text-indigo-600">{stat.value}</p>
                <p className="text-sm text-gray-500 mt-1">{stat.label}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Featured Paths */}
      <section className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
        <div className="text-center mb-12">
          <h2 className="text-3xl font-bold text-gray-900 mb-3">Rutas de aprendizaje</h2>
          <p className="text-gray-500 max-w-2xl mx-auto">
            Elige tu camino y comienza a aprender con recursos curados, hitos claros y evaluaciones para validar tu progreso.
          </p>
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          {featuredPaths.map((path) => (
            <PathCard key={path.id} path={path} />
          ))}
        </div>
        <div className="text-center mt-10">
          <Link href="/paths" className="inline-flex items-center gap-2 text-indigo-600 hover:text-indigo-700 font-medium">
            Ver todas las rutas
            <svg className="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
            </svg>
          </Link>
        </div>
      </section>

      {/* How it works */}
      <section className="bg-white border-t border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
          <h2 className="text-3xl font-bold text-gray-900 text-center mb-12">Como funciona</h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {[
              { icon: 'M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4', title: '1. Elige tu ruta', desc: 'Selecciona entre rutas curadas de Backend, Frontend, Data o Mobile.' },
              { icon: 'M13 10V3L4 14h7v7l9-11h-7z', title: '2. Aprende paso a paso', desc: 'Sigue los hitos en orden con recursos gratuitos y de alta calidad.' },
              { icon: 'M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z', title: '3. Valida tu conocimiento', desc: 'Completa evaluaciones para demostrar tu dominio en cada tema.' },
            ].map((step) => (
              <div key={step.title} className="text-center">
                <div className="w-14 h-14 bg-indigo-100 rounded-xl flex items-center justify-center mx-auto mb-4">
                  <svg className="w-7 h-7 text-indigo-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d={step.icon} />
                  </svg>
                </div>
                <h3 className="font-semibold text-lg mb-2">{step.title}</h3>
                <p className="text-sm text-gray-500">{step.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-900 text-gray-400">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
          <div className="flex flex-col md:flex-row justify-between items-center">
            <div className="flex items-center gap-2 mb-4 md:mb-0">
              <div className="w-8 h-8 bg-indigo-600 rounded-lg flex items-center justify-center">
                <span className="text-white font-bold text-sm">SP</span>
              </div>
              <span className="font-bold text-white">SkillPath</span>
            </div>
            <p className="text-sm">Desarrollado por Ronald. Todos los derechos reservados.</p>
          </div>
        </div>
      </footer>
    </div>
  );
}
