<script setup>
import Account from "../components/Account.vue"
let accountLabel = ''
let accounts = [
  {
    "id": "1",
    "label": "bruh",
    "balance": 400
  },
  {
    "id": "2",
    "label": "bruh2",
    "balance": 500
  },
  {
    "id": "3",
    "label": "bruh3",
    "balance": 600
  },
]

async function createAccount(){
}
</script>

<template>
  <main class="mx-auto w-[50%] flex-col items-center h-screen flex gap-2">
    <div class=" w-full flex flex-col gap-2">
      <h1 class="text-2xl font-bold pt-4">My Accounts</h1>
      <p>Track your accounts balances and transactions</p>
    </div>
    <div class="flex justify-between w-full gap-2" action="">
      <input v-model="accountLabel" class="focus:outline-none border w-full rounded-md tex-gray-500 p-2" placeholder="Enter account label"
        type="text">
      <button @click="createAccount" type="button" class="bg-black text-white rounded-md w-[30%]">+ Add Account</button>
    </div>
    <ul class="w-full flex flex-col gap-2">
      <Account v-for="account in accounts" :href="'/accounts/' + account.id" :label="account.label"
        :balance="account.balance" />
    </ul>
  </main>
</template>
