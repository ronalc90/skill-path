'use client';

import Link from 'next/link';
import { useAuth } from '@/lib/auth-context';
import { useState } from 'react';
import SearchBar from './SearchBar';

export default function Navbar() {
  const { user, isAuthenticated, logout } = useAuth();
  const [menuOpen, setMenuOpen] = useState(false);

  return (
    <nav className="bg-white border-b border-gray-200 sticky top-0 z-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center gap-8">
            <Link href="/" className="flex items-center gap-2">
              <div className="w-8 h-8 bg-indigo-600 rounded-lg flex items-center justify-center">
                <span className="text-white font-bold text-sm">SP</span>
              </div>
              <span className="font-bold text-xl text-gray-900">SkillPath</span>
            </Link>
            <div className="hidden md:flex items-center gap-6">
              <Link href="/paths" className="text-gray-600 hover:text-gray-900 text-sm font-medium">
                Rutas
              </Link>
              {isAuthenticated && (
                <>
                  <Link href="/dashboard" className="text-gray-600 hover:text-gray-900 text-sm font-medium">
                    Dashboard
                  </Link>
                  <Link href="/profile" className="text-gray-600 hover:text-gray-900 text-sm font-medium">
                    Perfil
                  </Link>
                </>
              )}
            </div>
          </div>

          <div className="hidden md:flex items-center gap-4">
            <SearchBar />
            {isAuthenticated ? (
              <div className="flex items-center gap-3">
                <span className="text-sm text-gray-600">{user?.display_name}</span>
                <button
                  onClick={logout}
                  className="text-sm text-gray-500 hover:text-gray-700"
                >
                  Cerrar sesion
                </button>
              </div>
            ) : (
              <div className="flex items-center gap-3">
                <Link
                  href="/login"
                  className="text-sm font-medium text-gray-600 hover:text-gray-900"
                >
                  Iniciar sesion
                </Link>
                <Link
                  href="/register"
                  className="text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 px-4 py-2 rounded-lg"
                >
                  Registrarse
                </Link>
              </div>
            )}
          </div>

          {/* Mobile menu button */}
          <div className="md:hidden flex items-center">
            <button
              onClick={() => setMenuOpen(!menuOpen)}
              className="text-gray-600 hover:text-gray-900"
            >
              <svg className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                {menuOpen ? (
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                ) : (
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
                )}
              </svg>
            </button>
          </div>
        </div>
      </div>

      {/* Mobile menu */}
      {menuOpen && (
        <div className="md:hidden border-t border-gray-200 py-3 px-4 space-y-3">
          <SearchBar />
          <Link href="/paths" className="block text-gray-600 hover:text-gray-900 text-sm font-medium">
            Rutas
          </Link>
          {isAuthenticated ? (
            <>
              <Link href="/dashboard" className="block text-gray-600 hover:text-gray-900 text-sm font-medium">
                Dashboard
              </Link>
              <Link href="/profile" className="block text-gray-600 hover:text-gray-900 text-sm font-medium">
                Perfil
              </Link>
              <button onClick={logout} className="block text-sm text-gray-500 hover:text-gray-700">
                Cerrar sesion
              </button>
            </>
          ) : (
            <>
              <Link href="/login" className="block text-sm font-medium text-gray-600">
                Iniciar sesion
              </Link>
              <Link href="/register" className="block text-sm font-medium text-indigo-600">
                Registrarse
              </Link>
            </>
          )}
        </div>
      )}
    </nav>
  );
}
