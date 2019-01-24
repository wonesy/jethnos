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
      <div class="column is-one-third no-top-margin">
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
        {
          'uuid': 'aaaa-aaaa',
          'name': 'game0',
          'numClients': 1,
          'isStarted': false,
          'params': {
            'tribes': [],
            'mode': 0,
            'leader': {
              'uuid': 'dddd-aaaa-cccc-eeee'
            }
          }
        }
      ],
      selectedGame: null
    }
  },
  created () {
    // this.$store.dispatch('fetchToken')
    this.$store.dispatch('setWebsocket')
    this.getGameList()
  },
  watch: {
    joinedGame: function (val, oldVal) {
      if (val === null) {
        return
      }
      this.joinGame(val)
    }
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
    },
    getGameList: function () {
      this.$http.get('http://localhost:4444/game/list')
        .then(stream => stream.json())
        .then(data => (this.gamesList = data))
    },
    joinGame: function (gameUUID) {
      // mark that shit as joined
      // set the screen accordingly
      console.log('we have joined a brand new game')
      this.$router.go({name: 'game', params: {gameID: gameUUID}})
    }
  },
  computed: {
    tabData: function () {
      if (this.currentTabComponent === GameDetails) {
        return this.selectedGame
      }
    },
    joinedGame: function () {
      return this.$store.getters.gameUUID
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
  max-height: calc(100vh - 230px);
}

.stats-col {
  margin-right: 10px;
  background: #fafafa;
}

.btn-newgame {
  margin-bottom: 10px;
}

.no-top-margin {
  margin-top: 0 !important;
  padding-top: 0;
}
</style>
