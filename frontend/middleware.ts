import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

/**
 * Next.js Middleware
 * This middleware runs before each request is completed
 * Useful for authentication, logging, redirects, etc.
 */

export function middleware(request: NextRequest) {
  // Log incoming requests
  console.log(`[Middleware] ${request.method} ${request.url}`);

  // Example: Add custom headers
  const requestHeaders = new Headers(request.headers);
  requestHeaders.set('x-request-timestamp', Date.now().toString());

  // Example: Authentication check (commented out)
  // const token = request.cookies.get('token')?.value;
  // if (!token && request.nextUrl.pathname.startsWith('/orders')) {
  //   return NextResponse.redirect(new URL('/login', request.url));
  // }

  // Example: Rate limiting or IP blocking could be implemented here
  const ip = request.headers.get('x-forwarded-for') || request.headers.get('x-real-ip') || 'unknown';
  console.log(`[Middleware] Request from IP: ${ip}`);

  // Continue with the request
  const response = NextResponse.next({
    request: {
      headers: requestHeaders,
    },
  });

  // Add custom response headers
  response.headers.set('x-request-id', crypto.randomUUID());
  response.headers.set('x-powered-by', 'Next.js Shop App');

  return response;
}

// Configure which paths the middleware runs on
export const config = {
  // Match all paths except static files and API routes
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     */
    '/((?!api|_next/static|_next/image|favicon.ico).*)',
  ],
};
