<template>
  <div>
    <p :class="{'is-me': isMe}">{{message.uuid}}: {{message.text}}</p>
  </div>
</template>

<script>
export default {
  name: 'Dialogue',
  props: {
    message: {
      type: Object
    }
  },
  data () {
    return {}
  },
  computed: {
    isMe: function () {
      return this.message.uuid === this.$store.getters.userUUID
    }
  }
}
</script>

<style scoped>
p {
  text-align: right;
}
.is-me {
  text-align: left;
}
</style>
