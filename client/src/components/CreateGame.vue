<template>
  <div class="container main">
    <div class="field">
      <label class="label">Name</label>
      <div class="control">
        <input class="input" type="text" placeholder="Enter game name">
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

    <div v-if="gameMode=='Democratic'" class="notification is-warning">
        You are a trusting leader. A vote for tribes will be held.
    </div>
    <div v-else class="choose-dictator-cards">
      <keep-alive>
        <TribeSelect/>
      </keep-alive>
    </div>

    <div class="field">
      <label class="label">Message</label>
      <div class="control">
        <textarea class="textarea" placeholder="Textarea"></textarea>
      </div>
    </div>

    <div class="field">
      <div class="control">
        <label class="checkbox">
          <input type="checkbox">
          I agree to the <a href="#">terms and conditions</a>
        </label>
      </div>
    </div>

    <div class="field">
      <div class="control">
        <label class="radio">
          <input type="radio" name="question">
          Yes
        </label>
        <label class="radio">
          <input type="radio" name="question">
          No
        </label>
      </div>
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
import TribeSelect from './TribeSelect'

export default {
  name: 'CreateGame',
  components: {
    TribeSelect
  },
  data () {
    return {
      modeOptions: ['Democratic', 'Dictatorship'],
      gameMode: 'Democratic'
    }
  },
  methods: {
    changeGameMode: function () {
      console.log(this.gameMode)
    }
  }
}
</script>

<style scoped>
.main {
  height: 100%;
  min-width: 500px;
  max-width: 500px;
}
</style>
