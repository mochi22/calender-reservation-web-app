// import NextAuth from "next-auth"
// import Google from 'next-auth/providers/google'
// import { PrismaAdapter } from "@next-auth/prisma-adapter"
// import { PrismaClient } from "@prisma/client"

// const prisma = new PrismaClient()

// export default NextAuth({
//   adapter: PrismaAdapter(prisma),
//   providers: [
//     Google({
//       clientId: process.env.GOOGLE_CLIENT_ID!,
//       clientSecret: process.env.GOOGLE_CLIENT_SECRET!,
//       authorization: {
//         params: {
//           prompt: "consent",
//           access_type: "offline",
//           response_type: "code"
//         }
//       }
//     }),
//   ],
//   callbacks: {
//     async signIn({ user, account, profile, email, credentials }) {
//       const allowedEmails = ['admin1@example.com', 'admin2@example.com', 'jdhdbemd@gmail.com'];
//       if (!allowedEmails.includes(user.email!)) {
//         return false; // 許可されていないメールアドレスの場合、ログインを拒否
//       }
//       console.log("ttt", user.email!);

//       try {
//         // Google IDを使用して管理者を検索
//         let admin = await prisma.administrator.findUnique({
//           where: { googleId: user.id },
//         })
        
//         if (!admin) {
//           // 管理者が存在しない場合、新しく作成
//           admin = await prisma.administrator.create({
//             data: {
//               email: user.email!,
//               name: user.name!,
//               googleId: user.id,
//               role: 'admin', // デフォルトの役割
//               lastLogin: new Date(),
//             },
//           })
//         } else {
//           // 既存の管理者の場合、最終ログイン時間を更新
//           await prisma.administrator.update({
//             where: { id: admin.id },
//             data: { lastLogin: new Date() },
//           })
//         }
//         return true
//       } catch (error) {
//         console.error("Error in signIn callback:", error)
//         return false
//       }
//     },
//     async session({ session, token, user }) {
//       if (session.user) {
//         session.user.image = token.picture as string;
        
//         // セッションに管理者情報を追加
//         try {
//           const admin = await prisma.administrator.findUnique({
//             where: { email: session.user.email! },
//           })
//           if (admin) {
//             session.user.role = admin.role;
//             session.user.adminId = admin.id;
//           }
//         } catch (error) {
//           console.error("Error in session callback:", error)
//         }
//       }
//       return session;
//     },
//     async jwt({ token, account }) {
//       if (account) {
//         token.accessToken = account.access_token
//       }
//       return token
//     }
//   },
// })





import NextAuth from "next-auth"
import Google from 'next-auth/providers/google'
// import GoogleProvider from "@auth/google-provider"

export default NextAuth({
  providers: [
    Google({
      clientId: process.env.GOOGLE_CLIENT_ID!,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET!,
      authorization: {
        params: {
          prompt: "consent",
          access_type: "offline",
          response_type: "code"
        }
      }
    }),
  ],
  callbacks: {
    async signIn({ user, account, profile, email, credentials }) {
      const allowedEmails = ['admin1@example.com', 'admin2@example.com', 'jdhdbemd@gmail.com'];
      return allowedEmails.includes(user.email!);
    },
    async session({ session, token }) {
      if (session.user) {
        session.user.image = token.picture as string;
      }
      return session;
    },
    async jwt({ token, account }) {
      if (account) {
        token.accessToken = account.access_token
      }
      return token
    }
  },
})
