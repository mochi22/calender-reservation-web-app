// pages/index.tsx
import Head from 'next/head'
import CalendarComponent from '../../components/Calendar'

export default function Home() {
  return (
    <div>
      <Head>
        <title>予約管理アプリ - カレンダー</title>
        <meta name="description" content="予約管理アプリのカレンダーページ" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="container mx-auto mt-8">
        <h1 className="text-3xl font-bold mb-4">予約カレンダー</h1>
        <CalendarComponent />
      </main>
    </div>
  )
}
