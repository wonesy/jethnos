<template>
  <div class="container user-container">
    <div class="field">
      <label class="label">Handle</label>
      <div class="control">
        <input
          class="input"
          type="text"
          placeholder="Enter handle"
          v-model="handle"
          v-bind:class="{'is-success': handleSaved}"
          >
      </div>
    </div>

    <div class="field">
      <label class="label">Email</label>
      <div class="control">
        <input
          class="input"
          type="email"
          placeholder="Email input (optional)"
          >
      </div>
      <!-- <p class="help is-danger">This email is invalid</p> -->
    </div>

    <div class="notification is-warning">
        You are 100% anonymous, which also means no stats will be saved.
        <strong>Create an account by clicking the button below</strong>
    </div>
    <div class="buttons is-centered">
      <button class="button is-link is-right">Create account</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'User',
  props: {
    defaultHandle: {
      type: String
    }
  },
  data () {
    return {
      handle: this.defaultHandle,
      handleSaved: true
    }
  },
  watch: {
    // whenever question changes, this function will run
    handle: function (newHandle, oldHandle) {
      this.handleSaved = false
      this.debouncedGetAnswer()
    }
  },
  created: function () {
    // _.debounce is a function provided by lodash to limit how
    // often a particularly expensive operation can be run.
    // In this case, we want to limit how often we access
    // yesno.wtf/api, waiting until the user has completely
    // finished typing before making the ajax request. To learn
    // more about the _.debounce function (and its cousin
    // _.throttle), visit: https://lodash.com/docs#debounce
    this.debouncedGetAnswer = this._.debounce(this.saveHandle, 800)
  },
  methods: {
    saveHandle: function () {
      this.$store.dispatch('setUserHandle', this.handle)
      this.handleSaved = true
    }
  }
}
</script>

<style scoped>
.user-container {
  margin-top: 20px;
  min-width: 250px;
  max-width: 500px;
}
</style>
