<template>
  <div class="container main">
    <div class="field">
      <label class="label">Name</label>
      <div class="control">
        <input class="input" type="text" v-model="gameName" placeholder="Enter game name">
      </div>
    </div>

    <div class="field">
      <label class="label">Game mode</label>
      <div class="control">
        <div class="select">
          <select v-model="gameMode" @change="changeGameMode()">
            <option v-for="mode in modeOptions" :key="mode">{{mode}}</option>
          </select>
        </div>
      </div>
    </div>

    <div v-if="gameMode=='Democracy'" class="notification is-warning">
        You are a trusting leader. A vote for tribes will be held.
    </div>
    <div v-else class="choose-dictator-cards">
      <keep-alive>
        <TribeSelect v-bind:submit="submit" v-on:is-valid="broadcastNewGame"/>
      </keep-alive>
    </div>

    <div class="field is-grouped">
      <div class="control">
        <button class="button is-link" v-on:click="submitGame">Submit</button>
      </div>
      <div class="control">
        <button class="button is-text">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script>
import TribeSelect from './TribeSelect'

export default {
  name: 'CreateGame',
  components: {
    TribeSelect
  },
  data () {
    return {
      gameName: '',
      modeOptions: ['Democracy', 'Dictatorship'],
      gameMode: 'Democracy',
      submit: false,
      newHubUUID: ''
    }
  },
  methods: {
    changeGameMode: function () {
      console.log(this.gameMode)
    },
    submitGame: function () {
      if (this.gameMode === 'Democracy') {
        // send game data to the store
        console.log('submitting democratic game')
        this.broadcastNewGame(true, this.gameName, null)
      } else {
        this.submit = true
      }
    },
    broadcastNewGame: function (isValid, name, tribeList) {
      // reset the submit value if the init settings are invalid
      if (!isValid) {
        this.submit = false
        return
      }

      if (!name || name === '') {
        this.submit = false
        return
      }

      // if we're here, then the settings are valid and we can save the information to the store
      // let requestOptions = {
      //   headers: authHeader()
      // }

      let newHubUrl = 'http://localhost:4444/createhub'
      this.$http.post(newHubUrl, {name: name})
        .then(stream => stream.json())
        .then(data => (this.newHubUUID = data))
        .catch(error => console.log(error))
    }
  }
}
</script>

<style scoped>
.main {
  margin-top: 20px;
  height: 100%;
  min-width: 250px;
  max-width: 1200px;
}
</style>
