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
      ws: null
    }
  },
  watch: {
    hubUUID: {
      handler: function (val, oldVal) {
        console.log('changing from ' + oldVal + ' to ' + val)
        let wsUrl = 'ws://localhost:4444/joinhub'

        if (this.ws !== null) {
          console.log('Closing web socket to hub=' + this.hubUUID)
          this.ws.close()
          delete this.ws
        }

        this.ws = new WebSocket(wsUrl)

        let data = {}
        data['hubUUID'] = this.hubUUID

        // fetch(wsUrl, {
        //   method: 'POST',
        //   headers: {
        //     'Accept': 'application/json',
        //     'Content-Type': 'application/json'
        //   },
        //   body: data.json()
        // })
        //   .then()
        //   .catch()
      }
    }
  },
  methods: {
    send () {
      console.log(this.chatText)
      let listHubsUrl = 'http://localhost:4444/listhubs'
      fetch(listHubsUrl)
        .then(stream => stream.json())
        .then(data => (this.hubs = data))
        .catch(error => console.error(error))
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
