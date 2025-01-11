// pages/index.tsx
import Head from 'next/head'

export default function Home() {
  return (
    <div>
      <Head>
        <title>予約管理アプリ - ホーム</title>
        <meta name="description" content="予約管理アプリのホームページ" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="container mx-auto mt-8">
        <h1 className="text-3xl font-bold">予約管理アプリへようこそ</h1>
        {/* ここにホームページのコンテンツを追加 */}
      </main>
    </div>
  )
}