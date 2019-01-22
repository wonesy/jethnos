/* eslint-disable */

import Vue from 'vue'
import Router from 'vue-router'
import Lobby from '@/components/Lobby'

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Lobby',
      component: Lobby
    }
  ]
})
