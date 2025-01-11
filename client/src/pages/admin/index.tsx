

import { useState } from 'react'
import { useSession } from 'next-auth/react'
import GoogleAuthButton from '../../components/GoogleAuthButton'

export default function AdminPage() {
  const { data: session } = useSession()
  const [isBasicAuth, setIsBasicAuth] = useState(false)

  if (session || isBasicAuth) {
    return (
      <div>
        <h1>Admin Dashboard</h1>
        {/* 管理者機能をここに実装 */}
      </div>
    )
  }

  return (
    <div>
      <h1>Admin Login</h1>
      <GoogleAuthButton />
      <hr />
      {/* <BasicAuthForm onAuthenticated={() => setIsBasicAuth(true)} /> */}
    </div>
  )
}


// import { useSession } from 'next-auth/react'
// import { useState, useEffect } from 'react'

// export default function AdminDashboard() {
//   const { data: session } = useSession()
//   const [users, setUsers] = useState([])

//   useEffect(() => {
//     if (session?.user.role) {
//       fetchUsers()
//     }
//   }, [session])

//   const fetchUsers = async () => {
//     const res = await fetch('/api/admin/users')
//     const data = await res.json()
//     setUsers(data)
//   }

//   if (!session || !session.user.role) {
//     return <div>Access Denied</div>
//   }

//   return (
//     <div>
//       <h1>Admin Dashboard</h1>
//       <p>Welcome, {session.user.name} ({session.user.role})</p>
//       {/* ユーザー管理UI */}
//     </div>
//   )
// }
