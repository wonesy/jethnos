<template>
  <div class="lobby-container is-fullheight">
    <div class="columns is-fullheight">
      <div class="column is-one-third lobby-view">
        <div class="buttons">
          <a
            v-for="hub in hubs"
            :key="hub.uuid"
            class="join-chat-btn button is-fullwidth"
            v-on:click='activeHub=hub.uuid'
            v-bind:class="{active: activeHub===hub.uuid}"
          >
            {{hub.uuid}}
          </a>
        </div>
      </div>
      <div class="column chat-room">
        <Chat :hubUUID="activeHub"/>
      </div>
    </div>
  </div>
</template>

<script>
import Chat from './Chat'
import { mapGetters } from 'vuex'

export default {
  name: 'Lobby',
  components: {
    Chat
  },
  data () {
    return {
      activeHub: '',
      hubs: []
    }
  },
  computed: {
    ...mapGetters([
      'token'
    ])
  },
  created () {
    this.$store.dispatch('fetchToken').then(() => {
      console.log('finished with the token fetching')
      this.listHubs()
    })
  },
  methods: {
    listHubs () {
      let token = this.$store.getters.token
      let requestOptions = {
        headers: { 'Authorization': 'Bearer ' + token }
      }
      console.log(requestOptions)
      let listHubsUrl = 'http://localhost:4444/listhubs'
      this.$http.get(listHubsUrl, requestOptions)
        .then(stream => stream.json())
        .then(data => (this.hubs = data))
        .catch(error => console.log(error))
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.lobby-container {
  height: 100%;
}

.lobby-view {
  height: 100%;
  background-color: rgb(0, 209, 178);
}

.chat-room {
  background-color: rgb(223, 223, 223);
  margin-top: 12px;
}

.active {
  color:blueviolet;
}

.is-fullheight {
  height: 100%;
}

.columns {
  margin-left: 0px;
  margin-right: 0px;
}

</style>
