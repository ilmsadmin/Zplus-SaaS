import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'Zplus SaaS - Multi-Tenant Platform',
  description: 'Modern multi-tenant SaaS platform with modular architecture',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className="font-sans">{children}</body>
    </html>
  )
}