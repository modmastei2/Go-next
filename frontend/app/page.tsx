import Link from "next/link";

export default function Home() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <div className="max-w-7xl mx-auto px-4 py-16">
        <header className="text-center mb-16">
          <h1 className="text-6xl font-bold text-gray-900 mb-4">
            Go-next Shop
          </h1>
          <p className="text-xl text-gray-600 mb-8">
            Full-stack application with Next.js frontend and Golang backend
          </p>
        </header>

        <div className="grid md:grid-cols-2 gap-8 max-w-4xl mx-auto">
          {/* Frontend Card */}
          <div className="bg-white rounded-lg shadow-xl p-8 hover:shadow-2xl transition-shadow">
            <div className="text-4xl mb-4">⚛️</div>
            <h2 className="text-2xl font-bold text-gray-900 mb-4">Frontend</h2>
            <ul className="space-y-2 text-gray-600 mb-6">
              <li>✓ Next.js 15 with App Router</li>
              <li>✓ TypeScript for type safety</li>
              <li>✓ TailwindCSS for styling</li>
              <li>✓ HTTP Interceptor pattern</li>
              <li>✓ Middleware examples</li>
            </ul>
            <Link
              href="/orders"
              className="block w-full bg-blue-600 text-white text-center py-3 rounded-lg hover:bg-blue-700 transition-colors font-semibold"
            >
              Go to Shop →
            </Link>
          </div>

          {/* Backend Card */}
          <div className="bg-white rounded-lg shadow-xl p-8 hover:shadow-2xl transition-shadow">
            <div className="text-4xl mb-4">🚀</div>
            <h2 className="text-2xl font-bold text-gray-900 mb-4">Backend</h2>
            <ul className="space-y-2 text-gray-600 mb-6">
              <li>✓ Golang with Fiber framework</li>
              <li>✓ GORM for database operations</li>
              <li>✓ Clean Architecture</li>
              <li>✓ Dependency Injection</li>
              <li>✓ Custom middleware</li>
            </ul>
            <a
              href="http://localhost:3001/health"
              target="_blank"
              rel="noopener noreferrer"
              className="block w-full bg-green-600 text-white text-center py-3 rounded-lg hover:bg-green-700 transition-colors font-semibold"
            >
              API Health Check →
            </a>
          </div>
        </div>

        <div className="mt-16 max-w-4xl mx-auto bg-white rounded-lg shadow-xl p-8">
          <h2 className="text-2xl font-bold text-gray-900 mb-4">Features Demonstrated</h2>
          <div className="grid md:grid-cols-2 gap-6">
            <div>
              <h3 className="font-semibold text-gray-900 mb-2">Frontend</h3>
              <ul className="space-y-1 text-gray-600 text-sm">
                <li>• Server and Client Components</li>
                <li>• HTTP Client with interceptors</li>
                <li>• Next.js Middleware</li>
                <li>• State management with hooks</li>
                <li>• API integration</li>
                <li>• Form handling</li>
              </ul>
            </div>
            <div>
              <h3 className="font-semibold text-gray-900 mb-2">Backend</h3>
              <ul className="space-y-1 text-gray-600 text-sm">
                <li>• RESTful API routes</li>
                <li>• Clean Architecture layers</li>
                <li>• Repository pattern</li>
                <li>• Business logic separation</li>
                <li>• Custom middleware (CORS, Logger, etc.)</li>
                <li>• Database migrations & seeding</li>
              </ul>
            </div>
          </div>
        </div>

        <footer className="mt-16 text-center text-gray-600">
          <p>Built with ❤️ using Go and Next.js</p>
        </footer>
      </div>
    </div>
  );
}
