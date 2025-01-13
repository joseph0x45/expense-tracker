<script setup>
const props = defineProps(['id', 'label', 'balance'])
</script>

<template>
  <a class="hover:bg-gray-50 p-2 w-full flex justify-between gap-2 border rounded-md">
    <h1>{{ props.label }}</h1>
    <h1>{{ props.balance }}</h1>
  </a>
</template>
