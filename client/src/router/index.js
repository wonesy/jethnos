/* eslint-disable */

import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/components/Main'
import Game from '@/components/Game'

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/game',
      name: 'Game',
      component: Game
    },
    {
      path: '/',
      name: 'Main',
      component: Main
    }
  ]
})
