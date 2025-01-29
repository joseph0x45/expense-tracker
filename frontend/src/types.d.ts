type Account = {
  id: string
  label: string
  balance: number
  treshold: number
}

interface AccountProps {
  account: Account
}
