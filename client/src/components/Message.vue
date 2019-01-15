<template>
  <div v-bind:class="{'is-me': msg.isMe, 'is-sys': msg.name==='System'}">
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

.is-me p {
  text-align: left;
}

.is-sys p {
  text-align: center;
  color: dimgray;
}
</style>
