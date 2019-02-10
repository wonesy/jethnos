<template>
  <canvas id="board"></canvas>
</template>

<script>
/* eslint-disable */
import { paper, Path, Point } from 'paper'
import  { Board } from './../game_rules/animation.js'

export default {
  name: 'Board',
  data () {
    return {
      board: null,
      gamestate: null
    }
  },
  created () {
    paper.install(window)
  },
  mounted () {
<<<<<<< HEAD
    var img = require('./../assets/slovakia.png')
    console.log(img)
    this.board = new Board('board', img)
    // this.board.createRealms()
=======
    var map = require('../assets/slovakia.png')
    this.board = new Board('board', map)
    this.gamestate = board.gamestate
>>>>>>> working on map
  }
}
</script>

<style scoped>
canvas {
  height: 100%;
  width: 100%;
}

.pointer-cursor {
  cursor: pointer;
}
</style>
