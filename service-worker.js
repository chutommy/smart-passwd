importScripts('https://storage.googleapis.com/workbox-cdn/releases/6.1.5/workbox-sw.js');

// Cache CSS, JS, and Web Worker requests with a Stale While Revalidate strategy
workbox.routing.registerRoute(
  // Check to see if the request's destination is style for stylesheets, script for JavaScript, or worker for web worker
  ({request}) =>
    request.destination === 'style' ||
    request.destination === 'script' ||
    request.destination === 'worker',
  // Use a Stale While Revalidate caching strategy
  new workbox.strategies.StaleWhileRevalidate({
    // Put all cached files in a cache named 'assets'
    cacheName: 'assets',
  }),
);

// Cache images with a Cache First strategy
workbox.routing.registerRoute(
  // Check to see if the request's destination is style for an image
  ({request}) => request.destination === 'image',
  // Use a Cache First caching strategy
  new workbox.strategies.CacheFirst({
    // Put all cached files in a cache named 'images'
    cacheName: 'images',
    plugins: [
      // Don't cache more than 50 items, and expire them after 30 days
      new workbox.expiration.ExpirationPlugin({
        maxEntries: 50,
        maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // 3 Days
      }),
    ],
  }),
);

// Catch routing errors, like if the user is offline
workbox.routing.setCatchHandler(async ({event}) => {
  // Return the precached offline page if a document is being requested
  if (event.request.destination === 'document') {
    return workbox.precaching.matchPrecache('/offline.html');
  }

  return Response.error();
});