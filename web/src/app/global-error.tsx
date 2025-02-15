'use client' // Error boundaries must be Client Components

export default function GlobalError({
  error,
}: {
  error: Error & { digest?: string }
}) {
  return (
    // global-error must include html and body tags
    <html>
      <body>
        <h2>Something went wrong!</h2>
        <h2>{error.message}</h2>
      </body>
    </html>
  )
}
