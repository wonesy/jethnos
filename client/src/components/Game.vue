<template>
  <div class="xx game-container">
    <div class="xx players-bar">players</div>
    <div class="xx game-board"><board></board></div>
    <div class="chat-box"><chat></chat></div>
    <div class="xx game-info">info</div>
  </div>
</template>

<script>
import Chat from './Chat'
import Board from './Board'

export default {
  name: 'Game',
  components: {
    Chat,
    Board
  },
  data () {
    return {}
  },
  methods: {
    getGameParameters: function () {}
  }
}
</script>

<style scoped>
.xx {
  border: 1px solid black;
}

.game-container {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  grid-gap: 10px;
  grid-auto-rows: minmax(100px, auto);
  min-height: 100vh;
  padding: 20px;
  max-width: 1750px;
  margin: auto;
}

.players-bar {
  grid-column: 1 / 5;
  grid-row: 1;
}

.game-board {
  grid-column: 1 / 5;
  grid-row: 2 / 5;
}

.chat-box {
  grid-column: 1 / 3;
  grid-row: 5 / 8;
}

.game-info {
  grid-column: 3 / 5;
  grid-row: 5 / 8;
}
</style>
