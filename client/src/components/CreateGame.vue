<template>
<div>
  <div class="notification is-warning">
    Create a new game
  </div>
  <div class="field">
    <label class="label">Game name</label>
    <div class="control">
      <input class="input" type="text" placeholder="Game name" v-model="gameName">
    </div>
  </div>

  <div class="field">
    <label class="label">Game mode</label>
    <div class="control">
      <div class="select">
        <select v-model="selected">>
          <option
            v-for="(mode,idx) in gameModes"
            :key="idx">
            {{mode}}
            </option>
        </select>
      </div>
    </div>
  </div>

  <div v-if="selected==='Democracy'" class="notification is-primary">
    You are a trusting leader. There will be a vote among game participants before
    the game starts
  </div>
  <div v-else>
    <tribe-choose v-on:update-tribes="updateTribes"></tribe-choose>
  </div>

  <div class="field is-grouped">
    <div class="control">
      <button class="button is-link" :disabled="!gameValid" @click="createNewGame">Submit</button>
    </div>
  </div>
</div>
</template>

<script>
import TribeChoose from './TribeChoose'

export default {
  name: 'CreateGame',
  components: {
    TribeChoose
  },
  data () {
    return {
      gameName: '',
      gameModes: ['Democracy', 'Dictatorship'],
      selected: 'Democracy',
      submitIsDisabled: true,
      selectedTribes: [],
      wtf: null
    }
  },
  computed: {
    gameValid: function () {
      let nameValid = (this.gameName !== '')
      let modeValid = (this.selected === 'Democracy' || this.selectedTribes.length === 6)

      if (nameValid && modeValid) {
        return true
      }
      return false
    }
  },
  methods: {
    updateTribes: function (tribes) {
      this.selectedTribes = tribes
    },
    createNewGame: function () {
      let postData = {
        'name': this.gameName,
        'uuid': this.$store.getters.userUUID,
        'tribes': this.selectedTribes,
        'mode': this.selected
      }

      this.$http.post('http://localhost:4444/game/new', postData)
        .then(stream => stream.json())
        .then(data => (this.$store.dispatch('setJoinedGame', data)))
    }
  }
}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
