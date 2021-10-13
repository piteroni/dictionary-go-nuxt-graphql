<template>
  <header class="flex bg-white border-b border-gray-200 h-20 items-center header-shadow">
    <div class="px-6 w-full flex justify-between">
      <slot />
    </div>
  </header>
</template>

<style scss scoped>
.header-shadow {
  margin-bottom: 5px;
  box-shadow: 0 3px 6px 0 rgba(99, 78, 78, 0.2);
}
</style>