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

    <div class="choose-dictator-cards"></div>

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
export default {
  name: 'CreateGame',
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
  min-width: 250px;
  max-width: 500px;
}
</style>
