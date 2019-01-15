<template>
  <div class="hero is-primary banner">
    <button v-on:click="$emit('change-tab', 'lobby')">
      Lobby
    </button>
    <button v-on:click="$emit('change-tab', 'banner')">
      Banner
    </button>
  </div>
</template>

<script>
export default {
  name: 'Banner',
  data () {
    return {}
  }
}
</script>

<style scoped>
.banner {
  min-height: 100px;
}
</style>
