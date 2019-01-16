<template>
  <div class=" main-container is-fullheight">
    <div class="banner">
      <Banner v-on:change-tab="changeTab"/>
    </div>
    <keep-alive>
      <component v-bind:is="currentTabComponent" :defaultHandle="anonHandle"></component>
    </keep-alive>
  </div>
</template>

<script>
import Banner from './Banner'
import Lobby from './Lobby'
import User from './User'
import CreateGame from './CreateGame'

export default {
  name: 'Main',
  components: {
    Banner,
    Lobby,
    User,
    CreateGame
  },
  data () {
    return {
      currentTabComponent: Lobby,
      anonHandle: ''
    }
  },
  created: function () {
    this.anonHandle = 'anon' + (Math.floor(Math.random() * 10000)).toString()
    this.$store.dispatch('setUserHandle', this.anonHandle)
  },
  methods: {
    changeTab: function (tab) {
      if (tab === 'lobby') {
        this.currentTabComponent = Lobby
      } else if (tab === 'user') {
        this.currentTabComponent = User
      } else if (tab === 'create') {
        this.currentTabComponent = CreateGame
      }
    }
  }
}
</script>

<style scoped>
.is-fullheight {
  height: 100%;
}

.main-container {
  display: flex;
  height: 100%;
  /* min-height: 98vh; */
  flex-direction: column;
}

.lobby {
  flex: 1;
  overflow: hidden;
  overflow-y: scroll;
}
</style>
