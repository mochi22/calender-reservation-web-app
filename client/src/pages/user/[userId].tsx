// pages/[userId]/index.tsx

import { useRouter } from 'next/router'
import { useState, useEffect } from 'react'
import CalendarGfg from '../calendar/CalendarGfg'

export default function UserPage() {
  const router = useRouter()
  const { userId } = router.query
  const [userName, setUserName] = useState('')

  useEffect(() => {
    // ここでユーザー名を取得する処理を実装
    // 例: APIからユーザー情報を取得する
    const fetchUserName = async () => {
      // 仮のユーザー名取得処理
      setUserName(`User ${userId}`)
    }
    if (userId) {
      fetchUserName()
    }
  }, [userId])

  if (!userId) return <div>Loading...</div>

  return (
    <div>
      <h1>Welcome, {userName}!</h1>
      <CalendarGfg userId={userId as string} />
    </div>
  )
}
