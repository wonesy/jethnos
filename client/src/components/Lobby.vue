<template>
<div>
  <div class="columns">
    <div class="column is-one-third lobby-view">
      yo
    </div>
    <div class="column chat-room">
      <Chat/>
    </div>
  </div>
</div>
</template>

<script>
import Chat from './Chat'

export default {
  name: 'Lobby',
  components: {
    Chat
  },
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.lobby-view {
  background-color: aquamarine;
}

.chat-room {
  background-color: rgb(223, 223, 223);
  border: 1px black solid;
}
</style>
