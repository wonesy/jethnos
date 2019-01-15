<template>
  <div class=" main-container is-fullheight">
    <div class="banner">
      <Banner v-on:change-tab="changeTab"/>
    </div>
    <keep-alive>
      <component v-bind:is="currentTabComponent"></component>
    </keep-alive>
  </div>
</template>

<script>
import Banner from './Banner'
import Lobby from './Lobby'

export default {
  name: 'Main',
  components: {
    Banner,
    Lobby
  },
  data () {
    return {
      currentTabComponent: Lobby
    }
  },
  methods: {
    changeTab: function (tab) {
      if (tab === 'lobby') {
        this.currentTabComponent = Lobby
      } else if (tab === 'banner') {
        this.currentTabComponent = Banner
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
