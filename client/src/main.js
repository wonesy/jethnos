// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuex from 'vuex'
import VueLodash from 'vue-lodash'
import App from './App'
import router from './router'
import './../node_modules/bulma/css/bulma.css'
import VueResource from 'vue-resource'

Vue.use(VueResource)
Vue.use(Vuex)
Vue.use(VueLodash)

Vue.config.productionTip = false

// eslint-disable-next-line
const store = new Vuex.Store({
  state: {
    token: '',
    user: {
      handle: ''
    },
    game: {
      name: '',
      tribes: []
    }
  },
  mutations: {
    SET_TOKEN (state, token) {
      state.token = token
    },
    SET_USER_HANDLE (state, handle) {
      state.user.handle = handle
    },
    ADD_GAME_TRIBE (state, tribe) {
      state.game.tribes.push(tribe)
    }
  },
  actions: {
    fetchToken ({commit}) {
      return new Promise((resolve, reject) => {
        Vue.http.get('http://localhost:4444/gettoken')
          .then((response) => {
            commit('SET_TOKEN', response.body.token)
            resolve()
          })
          .catch(error => {
            console.log(error.statusText)
            reject(error.statusText)
          })
      })
    },
    setUserHandle ({commit}, handle) {
      commit('SET_USER_HANDLE', handle)
    },
    addTribe ({commit}, tribe) {
      commit('ADD_GAME_TRIBE', tribe)
    }
  },
  getters: {
    token: state => {
      return state.token
    },
    userHandle: state => {
      return state.user.handle
    },
    gameTibes: state => {
      return state.game.tribes
    }
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

Vue.http.interceptors.push(function (request) {
  let token = store.getters.token

  if (token) {
    // modify headers
    request.headers.set('Authorization', 'Bearer ' + token)
  }
})
