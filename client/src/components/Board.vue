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
      board: null
    }
  },
  created () {
    paper.install(window)
  },
  mounted () {
    var img = require('./../assets/slovakia.png')
    this.board = new Board('board', img)
    // this.board.createRealms()
  }
}
</script>

<style scoped>
canvas {
  height: 100%;
  width: 100%;
}
</style>
