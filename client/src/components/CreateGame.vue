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
      <button class="button is-link" :disabled="!gameValid">Submit</button>
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
      selectedTribes: []
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
      console.log('yo')
      this.selectedTribes = tribes
    }
  }
}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
