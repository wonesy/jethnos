// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuex from 'vuex'
import App from './App'
import router from './router'
import './../node_modules/bulma/css/bulma.css'

Vue.use(Vuex)
Vue.config.productionTip = false

// eslint-disable-next-line
const store = new Vuex.Store({
  state: {
    token: ''
  },
  mutations: {
    SET_TOKEN (state, token) {
      state.token = token
    }
  },
  actions: {
    setToken (context, token) {
      context.commit('SET_TOKEN', token)
    }
  },
  getters: {
    token: state => state.token
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  router,
  components: { App },
  template: '<App/>'
})
