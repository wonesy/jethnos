<template>
  <div v-bind:class="{isme: msg.isMe}">
    <p>{{msg.name}}: {{msg.text}}</p>
  </div>
</template>

<script>
export default {
  name: 'Message',
  props: {
    msg: {
      name: {
        type: String
      },
      text: {
        type: String
      },
      isMe: {
        type: Boolean
      }
    }
  },
  data () {
    return {}
  }
}
</script>

<style scoped>
p {
  text-align: right;
}

.isme p {
  text-align: left;
}
</style>
