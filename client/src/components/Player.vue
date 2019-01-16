<template>
  <li v-bind:class="{'is-active': isActive}">
    <a>
      <span>{{name}}</span>
    </a>
  </li>
</template>

<script>
export default {
  name: 'Player',
  props: {
    name: {
      type: String
    },
    isActive: {
      type: Boolean
    }
  },
  data () {
    return {}
  }
}
</script>

<style scoped>

</style>
