<template>
<div class="chat-container box">
  <div id="chatbox" class="messages box"></div>
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
    // create a new websocket
    let wsUrl = 'ws://localhost:4444/ws'
    if (this.ws === null) {
      this.ws = new WebSocket(wsUrl)
    }

    // create an event listener as well to receieve messages
    this.ws.addEventListener('message', function (e) {
      console.log('herererere')
      console.log(e)
      var msg = JSON.parse(e.data.trim())
      var chatContent = '<p>' + msg.name + ': ' + msg.text + '</p>'
      console.log(chatContent)
      var element = document.getElementById('chatbox')
      element.insertAdjacentHTML('beforeend', chatContent)
      console.log(element)
      element.scrollTop = element.scrollHeight // Auto scroll to the bottom
    })
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
      let chatMessage = {
        'type': 'chat',
        'data': {
          'name': 'someone',
          'text': this.chatText
        }
      }

      let data = JSON.stringify(chatMessage)
      this.ws.send(data)
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
