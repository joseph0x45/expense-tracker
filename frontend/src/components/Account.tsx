export default function Account({ account }: AccountProps) {
  return (
    <a href={`/accounts/${account.id}/transactions`} className="shadow-sm flex justify-between w-full px-2 py-4 rounded-md border-1 border-gray-300">
      <h1 className="font-semibold">{account.label}</h1>
      <h1 className={`${account.balance > account.treshold ? "text-green-500" : "text-red-500"}`}>
        {account.balance} XOF
      </h1>
    </a>
  )
}

