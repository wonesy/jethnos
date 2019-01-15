<template>
<div class="chat-container box">
  <div id="chatbox" class="messages box">
    <Message
      v-for="(msg,index) in messages"
      :key="index"
      :msg="msg"/>
  </div>
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
import Message from './Message.vue'

export default {
  name: 'Chat',
  components: {
    Message
  },
  props: {
    hubUUID: {
      type: String
    }
  },
  data () {
    return {
      ws: null,
      chatText: '',
      clientUUID: null,
      messages: []
    }
  },
  created: function () {
    // save the "Vue" instance for closures
    // let vm = this

    // create a new websocket
    let wsUrl = 'ws://localhost:4444/ws'
    if (this.ws === null) {
      this.ws = new WebSocket(wsUrl)
    }

    // as soon as the websocket is ready, call whoami to get uuid
    this.ws.onopen = function (e) {
      let message = {
        'type': 'whoami',
        'data': null
      }
      e.explicitOriginalTarget.send(JSON.stringify(message))
    }

    // create an event listener as well to receieve messages
    this.ws.onmessage = this.handleMessages
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
          'text': this.chatText,
          'sender': this.clientUUID
        }
      }

      let data = JSON.stringify(chatMessage)
      this.ws.send(data)
    },
    handleMessages (e) {
      var msg = JSON.parse(e.data.trim())

      // if it is not a chat message, handle it here
      if (msg.hasOwnProperty('type')) {
        if (msg.type === 'whoami') {
          this.clientUUID = msg.uuid
        }
        return
      }

      // requried details for a Message component
      let msgProps = {
        'name': msg.name,
        'text': msg.text,
        'isMe': (msg.sender === this.clientUUID)
      }

      console.log(msg)
      this.messages.push(msgProps)

      this.$nextTick(() => {
        let e = document.getElementById('chatbox')
        e.scrollTop = e.scrollHeight
      })
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
  overflow: hidden;
  overflow-y: scroll;
}
</style>
