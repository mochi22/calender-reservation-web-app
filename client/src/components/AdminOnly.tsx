import { useSession } from "next-auth/react"
import { useRouter } from "next/router"

const AdminOnly = ({ children }: { children: React.ReactNode }) => {
  const { data: session, status } = useSession()
  const router = useRouter()

  if (status === "loading") {
    return <p>Loading...</p>
  }

  if (!session) {
    router.push("/")
    return null
  }

  // ここで管理者チェックを行うこともできます
  // if (!session.user.isAdmin) {
  //   router.push("/")
  //   return null
  // }

  return <>{children}</>
}

export default AdminOnly
