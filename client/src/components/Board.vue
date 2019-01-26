<template>
  <canvas id="board"></canvas>
</template>

<script>
// import './../game_rules/animation.js'
import { Board } from './../game_rules/animation.js'

export default {
  name: 'Board',
  data () {
    return {
      board: null
    }
  },
  mounted () {
    // paper.install(window)
    this.board = new Board('board')
    this.board.setupCountryPoints()
    // this.board.anotherLine()
  },
  methods: {
  }
}
</script>

<style scoped>
canvas {
  height: 100%;
  width: 100%;
}
</style>
