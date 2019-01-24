/* eslint-disable */

import Vue from 'vue'
import Router from 'vue-router'
import Lobby from '@/components/Lobby'
import Game from '@/components/Game'

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'lobby',
      component: Lobby
    },
    {
      path: '/game/:gameID',
      name: 'game',
      component: Game
    }
  ]
})
