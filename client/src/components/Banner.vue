<template>
  <div class="hero is-primary banner">
    <div class="tabs is-centered">
      <ul>
        <li class="is-active">
          <a v-on:click="$emit('change-tab', 'lobby')">Lobby</a>
        </li>
        <li>
          <a v-on:click="$emit('change-tab', 'user')">User</a>
        </li>
      </ul>
    </div>
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
