<template>
  <div class="chat-container">
    <div id="messagebox" class="chat-message-display">
      <dialogue
        v-for="(msg,idx) in chatMessages"
        :key="idx"
        v-bind:message="msg"/>
    </div>
    <div class="intput-bar field has-addons">
      <div class="control is-expanded">
        <input id="chatId" class="input" type="text" placeholder="Type to chat"
          v-model="chatText" @keyup.enter="sendChatMessage">
      </div>
      <div class="control">
        <a class="button is-info" @click="sendChatMessage">
          Send
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import Dialogue from './Dialogue'
export default {
  name: 'Chat',
  components: {
    Dialogue
  },
  data () {
    return {
      chatMessages: [],
      chatText: ''
    }
  },
  created () {
    this.setSocketHandler()
  },
  watch: {
    newWebsocket: function (val, oldVal) {
      this.setSocketHandler()
    }
  },
  computed: {
    newWebsocket: function () {
      return this.$store.getters.websocket
    }
  },
  methods: {
    // TODO handle exceptions better here
    setSocketHandler: function () {
      if (this.$store.getters.websocket === null) {
        this.$store.dispatch('setWebsocket')
        // maybe implement a watch and return here
      }
      this.$store.getters.websocket.onmessage = this.messageHandler
    },
    messageHandler: function (e) {
      console.log('handling a message')
      let data = JSON.parse(e.data)

      if (data.type === 'chat') {
        this.chatHandler(data)
      }
    },
    chatHandler: function (data) {
      this.chatMessages.push(data)
      this.$nextTick(function () {
        var msgDiv = document.getElementById('messagebox')
        msgDiv.scrollTop = msgDiv.scrollHeight
      })
    },
    sendChatMessage: function () {
      let txt = this._.trim(this.chatText)
      this.chatText = ''

      if (txt === '') {
        return
      }

      let message = {
        'type': 'chat',
        'uuid': this.$store.getters.userUUID,
        'text': txt
      }

      this.$store.getters.websocket.send(JSON.stringify(message))
    }
  }
}
</script>

<style scoped>

.chat-container {
  display: grid;
  grid-template-rows: auto 30px;
  height: 100%;
  max-height: 100%;
}

.chat-message-display {
  grid-row: 1;
  max-height: 100%;
  margin-bottom: 5px;
  overflow-y: scroll;
}

.input-bar {
  grid-row: 2;
  margin: auto;
  padding: 50px;
}
</style>
