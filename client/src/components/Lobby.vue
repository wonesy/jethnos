<template>
  <div class="lobby-container">
    <!-- Header -->
    <div class="hero">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            Hero title
          </h1>
          <h2 class="subtitle">
            Subtitle
          </h2>
        </div>
      </div>
    </div>

    <!-- Body -->
    <div class="lobby-body columns">
      <div class="column is-one-third">
        <a class="button is-primary is-medium is-fullwidth btn-newgame"
          @click="setCurrentTab('createGame')">
          New Game</a>
        <div class="games-list-col">
          <GameListItem
            v-for="(game,idx) in gamesList"
            :game="game"
            :selected="game===selectedGame"
            @click="selectGame(game)"
            :key="idx" />
        </div>
      </div>
      <div class="column stats-col box">
        <div class="tabs is-large is-fullwidth">
          <ul>
            <li
              :class="{'is-active': isCurrentTab('gameDetails')}"
              @click="setCurrentTab('gameDetails')">
              <a>Game Details</a>
            </li>
            <li
              :class="{'is-active': isCurrentTab('accountDetails')}"
              @click="setCurrentTab('accountDetails')">
              <a>Account</a>
            </li>
          </ul>
        </div>
        <keep-alive>
          <component
            v-bind:is="currentTabComponent"
            :data="tabData">
          </component>
        </keep-alive>
      </div>
    </div>
  </div>
</template>

<script>
import GameListItem from './GameListItem'
import GameDetails from './GameDetails.vue'
import AccountDetails from './AccountDetails.vue'
import CreateGame from './CreateGame.vue'

export default {
  name: 'Lobby',
  components: {
    GameListItem,
    GameDetails,
    AccountDetails,
    CreateGame
  },
  data () {
    return {
      tabComponentTypes: {
        'gameDetails': GameDetails,
        'accountDetails': AccountDetails,
        'createGame': CreateGame
      },
      currentTabComponent: GameDetails,
      gamesList: [
        {'uuid': 'aaaa-aaaa', 'name': 'game0', 'numClients': 1, 'isStarted': false, 'mode': 0},
        {'uuid': 'bbbb-bbbb', 'name': 'game1', 'numClients': 1, 'isStarted': false, 'mode': 1},
        {'uuid': 'cccc-cccc', 'name': 'game2', 'numClients': 7, 'isStarted': false, 'mode': 2},
        {'uuid': 'dddd-dddd', 'name': 'game3', 'numClients': 5, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-dd5d', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-dd6d', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-dd7d', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-dd8d', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-dd9d', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-dd2d', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-dd3d', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-4ddd', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-5ddd', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-7ddd', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-9ddd', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0},
        {'uuid': 'dddd-2ddd', 'name': 'game3', 'numClients': 0, 'isStarted': false, 'mode': 0}
      ],
      selectedGame: null
    }
  },
  created () {
    // this.$store.dispatch('fetchToken')
  },
  methods: {
    selectGame: function (game) {
      this.selectedGame = game
      this.currentTabComponent = GameDetails
    },
    isCurrentTab: function (key) {
      let val = this.tabComponentTypes[key]
      return val === this.currentTabComponent
    },
    setCurrentTab: function (key) {
      let val = this.tabComponentTypes[key]
      if (val !== undefined) {
        this.currentTabComponent = val
      }
    }
  },
  computed: {
    tabData: function () {
      if (this.currentTabComponent === GameDetails) {
        return this.selectedGame
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.hero {
  max-height: 160px;
  min-height: 160px;
}

.lobby-container {
  /* height: 100%; */
  display: flex;
  flex-flow: column;
  /* overflow: hidden; */
}

.lobby-body {
  flex: 1;
  margin: 0;
}

.column {
  margin: 0
}

.games-list-col {
  overflow: auto;
  /* hero (160) + newgamebtn (10+30) */
  max-height: calc(100vh - 240px);
}

.stats-col {
  margin-right: 10px;
  background: #fafafa;
}

.btn-newgame {
  margin-bottom: 10px;
}
</style>
