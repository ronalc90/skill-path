const API_BASE = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1';

class ApiClient {
  private token: string | null = null;

  constructor() {
    if (typeof window !== 'undefined') {
      this.token = localStorage.getItem('token');
    }
  }

  setToken(token: string) {
    this.token = token;
    if (typeof window !== 'undefined') {
      localStorage.setItem('token', token);
    }
  }

  clearToken() {
    this.token = null;
    if (typeof window !== 'undefined') {
      localStorage.removeItem('token');
    }
  }

  getToken(): string | null {
    if (typeof window !== 'undefined' && !this.token) {
      this.token = localStorage.getItem('token');
    }
    return this.token;
  }

  private async request<T>(path: string, options: RequestInit = {}): Promise<T> {
    const headers: HeadersInit = {
      'Content-Type': 'application/json',
      ...options.headers,
    };

    const currentToken = this.getToken();
    if (currentToken) {
      (headers as Record<string, string>)['Authorization'] = `Bearer ${currentToken}`;
    }

    const response = await fetch(`${API_BASE}${path}`, {
      ...options,
      headers,
    });

    if (response.status === 401) {
      this.clearToken();
      if (typeof window !== 'undefined') {
        window.location.href = '/login';
      }
      throw new Error('Unauthorized');
    }

    if (!response.ok) {
      const error = await response.json().catch(() => ({ message: 'Request failed' }));
      throw new Error(error.message || error.error || 'Request failed');
    }

    return response.json();
  }

  // Auth
  async register(email: string, password: string, displayName: string) {
    return this.request<AuthResponse>('/auth/register', {
      method: 'POST',
      body: JSON.stringify({ email, password, display_name: displayName }),
    });
  }

  async login(email: string, password: string) {
    return this.request<AuthResponse>('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    });
  }

  async getMe() {
    return this.request<UserResponse>('/auth/me');
  }

  async updateProfile(data: UpdateProfileRequest) {
    return this.request<UserResponse>('/auth/profile', {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  // Paths
  async getPaths(params?: { category?: string; difficulty?: string; search?: string; page?: number }) {
    const searchParams = new URLSearchParams();
    if (params?.category) searchParams.set('category', params.category);
    if (params?.difficulty) searchParams.set('difficulty', params.difficulty);
    if (params?.search) searchParams.set('search', params.search);
    if (params?.page) searchParams.set('page', params.page.toString());
    const query = searchParams.toString();
    return this.request<PathListResponse>(`/paths${query ? `?${query}` : ''}`);
  }

  async getPathBySlug(slug: string) {
    return this.request<PathDetailResponse>(`/paths/${slug}`);
  }

  async getCategories() {
    return this.request<CategoryResponse[]>('/paths/categories');
  }

  // Progress
  async startPath(pathId: number) {
    return this.request<UserPathProgress>(`/paths/${pathId}/start`, { method: 'POST' });
  }

  async getPathProgress(pathId: number) {
    return this.request<UserPathProgress>(`/paths/${pathId}/progress`);
  }

  async completeMilestone(pathId: number, milestoneId: number) {
    return this.request<{ message: string }>(`/paths/${pathId}/milestones/${milestoneId}/complete`, {
      method: 'PUT',
    });
  }

  async completeResource(pathId: number, resourceId: number, rating?: number) {
    return this.request<{ message: string }>(`/paths/${pathId}/resources/${resourceId}/complete`, {
      method: 'PUT',
      body: JSON.stringify({ rating: rating || 0 }),
    });
  }

  async getMyPaths() {
    return this.request<UserPathProgress[]>('/my-paths');
  }

  // Assessments
  async getAssessment(milestoneId: number) {
    return this.request<AssessmentQuestion[]>(`/milestones/${milestoneId}/assessment`);
  }

  async submitAssessment(milestoneId: number, answers: AnswerSubmission[]) {
    return this.request<AssessmentResult>(`/milestones/${milestoneId}/assessment/submit`, {
      method: 'POST',
      body: JSON.stringify({ answers }),
    });
  }

  // Dashboard
  async getDashboard() {
    return this.request<DashboardResponse>('/dashboard');
  }

  // Search
  async search(query: string) {
    return this.request<SearchResponse>(`/search?q=${encodeURIComponent(query)}`);
  }
}

export const api = new ApiClient();

// Types
export interface AuthResponse {
  token: string;
  user: UserResponse;
}

export interface UserResponse {
  id: number;
  email: string;
  display_name: string;
  avatar_url?: string;
  bio?: string;
  github_url?: string;
  linkedin_url?: string;
}

export interface UpdateProfileRequest {
  display_name?: string;
  avatar_url?: string;
  bio?: string;
  github_url?: string;
  linkedin_url?: string;
}

export interface PathListResponse {
  paths: PathSummary[];
  total: number;
  page: number;
  page_size: number;
  total_pages: number;
}

export interface PathSummary {
  id: number;
  title: string;
  slug: string;
  description: string;
  difficulty: string;
  estimated_hours: number;
  category: string;
  icon_url?: string;
  milestone_count: number;
}

export interface PathDetailResponse {
  id: number;
  title: string;
  slug: string;
  description: string;
  difficulty: string;
  estimated_hours: number;
  category: string;
  icon_url?: string;
  is_published: boolean;
  milestones: MilestoneResponse[];
}

export interface MilestoneResponse {
  id: number;
  title: string;
  description: string;
  order_index: number;
  estimated_hours: number;
  resources: ResourceResponse[];
}

export interface ResourceResponse {
  id: number;
  title: string;
  url: string;
  type: string;
  provider?: string;
  is_free: boolean;
  estimated_minutes: number;
  order_index: number;
}

export interface CategoryResponse {
  name: string;
  count: number;
}

export interface UserPathProgress {
  path_id: number;
  path: PathSummary;
  started_at: string;
  completed_at?: string;
  last_activity_at: string;
  progress_pct: number;
  milestones?: MilestoneProgressDetail[];
}

export interface MilestoneProgressDetail {
  milestone_id: number;
  title: string;
  status: string;
  order_index: number;
}

export interface AssessmentQuestion {
  id: number;
  question: string;
  options: string[];
}

export interface AnswerSubmission {
  question_id: number;
  selected_index: number;
}

export interface AssessmentResult {
  score: number;
  total_questions: number;
  passed: boolean;
  details: AnswerDetail[];
}

export interface AnswerDetail {
  question_id: number;
  is_correct: boolean;
  correct_answer: number;
  explanation: string;
}

export interface DashboardResponse {
  paths_in_progress: number;
  paths_completed: number;
  total_hours_learned: number;
  completion_rate: number;
  recent_activity: RecentActivity[];
  skill_radar: SkillRadarPoint[];
  enrolled_paths: UserPathProgress[];
}

export interface RecentActivity {
  type: string;
  title: string;
  path_title: string;
  timestamp: string;
}

export interface SkillRadarPoint {
  category: string;
  score: number;
  max_score: number;
}

export interface SearchResponse {
  paths: PathSummary[];
  resources: ResourceResponse[];
}
