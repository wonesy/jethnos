<template>
<div class="chat-container box">
  <div class="messages box"></div>
  <!--  -->
  <footer class="footers">
    <div class="field has-addons">
      <p class="control is-expanded">
        <input v-model="chatText" class="input" type="text" placeholder="Type to chat">
      </p>
      <p class="control">
        <a v-on:click="send" class="button is-primary">
          Send
        </a>
      </p>
    </div>
  </footer>
</div>
</template>

<script>
export default {
  name: 'Chat',
  props: {
    hubUUID: {
      type: String
    }
  },
  data () {
    return {
      ws: null,
      chatText: ''
    }
  },
  created: function () {
    let wsUrl = 'ws://localhost:4444/ws'
    if (this.ws === null) {
      this.ws = new WebSocket(wsUrl)
    }
  },
  watch: {
    hubUUID: {
      immediate: false,
      handler: function (val, oldVal) {
        console.log('changing from ' + oldVal + ' to ' + val)
        let message = {
          'type': 'join',
          'data': {
            'hubUUID': val
          }
        }

        let data = JSON.stringify(message)
        this.ws.send(data)
      }
    }
  },
  methods: {
    send () {
      console.log(this.chatText)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.chat-container {
  display: flex;
  /* min-height: 98vh; */
  height: 100%;
  flex-direction: column;
}

.messages {
  flex: 1;
}
</style>
