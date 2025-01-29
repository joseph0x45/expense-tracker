import { useState } from "react"
import Account from "../components/Account"

export default function Home() {
  let [accounts, setAccounts] = useState<Array<Account>>([])
  return (
    <main className="mx-auto w-1/2 h-screen ">
      <div className="py-8">
        <h1 className="text-4xl font-bold">My Accounts</h1>
        <h1 className="text-gray-500">Track your account balances and transactions</h1>
      </div>
      <div className="flex flex-col gap-8">
        <form className="flex gap-2 ">
          <input className="flex-[4] rounded-md border-1 border-gray-300 p-2 focus:outline-none" placeholder="Enter account name" />
          <button className="flex justify-center items-center gap-2 flex-[1] bg-black text-white rounded-md p-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" className="size-5">
              <path d="M10.75 4.75a.75.75 0 0 0-1.5 0v4.5h-4.5a.75.75 0 0 0 0 1.5h4.5v4.5a.75.75 0 0 0 1.5 0v-4.5h4.5a.75.75 0 0 0 0-1.5h-4.5v-4.5Z" />
            </svg>
            <h1>Add Account</h1>
          </button>
        </form>
        <div className="flex flex-col gap-4">
          {
            accounts.map((account) => {
              return (
                <Account key={account.id} account={account} />
              )
            })
          }
        </div>
        <footer>Expense Tracker v1.0 - Made my <a className="text-blue-400 underline" href="https://github.com/joseph0x45" target="_blank">Me :)</a></footer>

      </div>
    </main>
  )
}
