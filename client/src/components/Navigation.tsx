import Link from 'next/link'
import { useRouter } from 'next/router'
import { useSession, signOut, signIn } from "next-auth/react"
import Image from 'next/image'

const Navigation = () => {
  const router = useRouter()
  const { data: session } = useSession()

  return (
    <nav className="bg-gray-800 text-white p-4">
      <div className="container mx-auto flex justify-between items-center">
        <ul className="flex space-x-4 items-center">
          <li>
            <Link href="/" className={router.pathname === "/" ? "font-bold" : ""}>
              ホーム
            </Link>
          </li>
          <li>
            <Link href="/calendar" className={router.pathname === "/calendar" ? "font-bold" : ""}>
              カレンダー
            </Link>
          </li>
          {session && (
            <li>
              <Link href="/admin" className={router.pathname === "/admin" ? "font-bold" : ""}>
                管理者ページ
              </Link>
            </li>
          )}
        </ul>
        <div className="flex items-center space-x-4">
          {session ? (
            <>
              {session.user?.image && (
                <Image
                  src={session.user.image}
                  alt="User profile"
                  width={32}
                  height={32}
                  className="rounded-full"
                />
              )}
              <span className="hidden md:inline">{session.user?.name}</span>
              <button onClick={() => signOut()} className="bg-red-500 hover:bg-red-600 px-3 py-1 rounded text-sm">
                ログアウト
              </button>
            </>
          ) : (
             <button onClick={() => signIn('google')} className="bg-blue-500 hover:bg-blue-600 px-3 py-1 rounded">
              ログイン
              </button>
          )}
        </div>
      </div>
    </nav>
  )
}

export default Navigation


// // components/Navigation.tsx
// import Link from 'next/link'
// import { useRouter } from 'next/router'
// import { useSession, signIn, signOut } from "next-auth/react"
// import Image from 'next/image'

// const Navigation = () => {
//   const router = useRouter()
//   const { data: session } = useSession()

//   return (
//     <nav className="bg-gray-800 text-white p-4">
//         <ul className="flex space-x-4 items-center">
//         <li>
//           <Link href="/" className={router.pathname === "/" ? "font-bold" : ""}>
//             ホーム
//           </Link>
//         </li>
//         <li>
//           <Link href="/calendar" className={router.pathname === "/calendar" ? "font-bold" : ""}>
//             カレンダー
//           </Link>
//         </li>
//         <li>
//           <Link href="/admin" className={router.pathname === "/admin" ? "font-bold" : ""}>
//             管理者ページ
//           </Link>
//         </li>
//       </ul>
//       {session ? (
//           <div className="flex items-center space-x-4">
//             {session.user?.image && (
//               <Image
//                 src={session.user.image}
//                 alt="User profile"
//                 width={32}
//                 height={32}
//                 className="rounded-full"
//               />
//             )}
//             <span>{session.user?.name}</span>
//             <button onClick={() => signOut()} className="bg-red-500 hover:bg-red-600 px-3 py-1 rounded">
//               ログアウト
//             </button>
//           </div>
//         ) : (
//           // <Link href="/admin" className="bg-blue-500 hover:bg-blue-600 px-3 py-1 rounded">
//              <button onClick={() => signIn('google')} className="bg-blue-500 hover:bg-blue-600 px-3 py-1 rounded">
//               ログイン
//               </button>
//           // </Link>
//         )}
//     </nav>
//   )
// }

// export default Navigation
