<template>
<div>
  <h1><strong>Create new game</strong></h1>
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
    <tribe-choose></tribe-choose>
  </div>

  <div class="field is-grouped">
    <div class="control">
      <button class="button is-link">Submit</button>
    </div>
    <div class="control">
      <button class="button is-text">Cancel</button>
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
      selected: 'Democracy'
    }
  }
}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
