<template>
  <div class="game-container">
    <!-- Body -->
    <div class="body-container is-fullwidth">
      <div class="tile is-ancestor tile-container">
        <div class="tile is-9 is-vertical is-parent">
          <div class="tile is-child box">

            <p class="title">Players</p>
            <div class="tabs is-toggle is-fullwidth is-large">
              <ul>
                <Player
                  v-for="(player,idx) in players"
                  :key="idx"
                  :name="player.name"
                  :isActive="player.isActive"
                />
              </ul>
            </div>

          </div>
          <div class="tile is-child box board-tile">
            <Board/>
          </div>
          <div class="tile is-parent">
            <div class="tile is-child box">
              <p class="title">My Stats</p>
            </div>
            <div class="tile is-child box">
              <p class="title">Game Stats</p>
            </div>
            <!-- <p class="title">My state</p> -->
          </div>
        </div>
        <div class="tile is-parent">
          <div class="tile is-child box">
            <p class="title">Chat</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Player from './Player'
import Board from './Board'

export default {
  name: 'Game',
  components: {
    Player,
    Board
  },
  props: {

  },
  data () {
    return {
      players: [],
      ws: null
    }
  },
  created () {
    for (let i = 0; i < 6; i++) {
      var x = {
        name: 'myName',
        isActive: false
      }
      if (i === 3) {
        x.isActive = true
      }
      this.players.push(x)
    }
    for (let i = 0; i < 6; i++) {
      console.log(this.players[i].name)
      console.log(this.players[i].isActive)
    }
  }
}
</script>

<style scoped>
.game-container {
  height: 100%;
  width: 100vw !important;
}

.body-container {
  height: 100%;
  margin-right: 10px;
}

.tile-container {
  margin: 0 !important;
  margin-top: 10px;
  height: calc(100% - 30px);
}
.tile {
  height: 100%;
}

.board-tile {
  min-height: 55%;
}
</style>
