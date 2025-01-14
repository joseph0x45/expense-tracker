<script setup>
import Account from "../components/Account.vue"
import { ref } from "vue"
let accountLabel = ref("")
let accounts = ref([])

async function createAccount() {
  try {
    const response = await fetch(
      "/api/accounts", {
      method: "POST",
      body: JSON.stringify({ label: accountLabel })
    })
    switch (response.status) {
      case 201:
        alert("Created")
        accounts.value = [
          ...accounts.value,
          { label: accountLabel.value, balance: 0}
        ]
        break
      case 409:
        alert("An account with the same label already exist")
        break
      default:
        alert("An error occured. Please check the logs")
    }
  } catch (error) {
    console.log(error)
  }
  accountLabel.value = ""
}
</script>

<template>
  <main class="mx-auto w-[50%] flex-col items-center h-screen flex gap-2">
    <div class=" w-full flex flex-col gap-2">
      <h1 class="text-2xl font-bold pt-4">My Accounts</h1>
      <p>Track your accounts balances and transactions</p>
    </div>
    <div class="flex justify-between w-full gap-2" action="">
      <input v-model="accountLabel" class="focus:outline-none border w-full rounded-md tex-gray-500 p-2"
        placeholder="Enter account label" type="text">
      <button @click="createAccount" type="button" class="bg-black text-white rounded-md w-[30%]">+ Add Account</button>
    </div>
    <ul class="w-full flex flex-col gap-2">
      <Account v-for="account in accounts" :key="account.label" :href="'/accounts/' + account.id" :label="account.label"
        :balance="account.balance" />
    </ul>
  </main>
</template>
