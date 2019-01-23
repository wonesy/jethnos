// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuex from 'vuex'
import VueLodash from 'vue-lodash'
import App from './App'
import router from './router'
import './../node_modules/bulma/css/bulma.css'
import './../node_modules/@fortawesome/fontawesome-free/css/all.css'
import VueResource from 'vue-resource'

Vue.use(VueResource)
Vue.use(Vuex)
Vue.use(VueLodash)

Vue.config.productionTip = false

// eslint-disable-next-line
const store = new Vuex.Store({
  state: {
    token: '',
    websocket: null,
    user: {
      clientUUID: '',
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
    SET_USER_UUID (state, uuid) {
      state.user.clientUUID = uuid
    },
    SET_WEBSOCKET (state, ws) {
      state.websocket = ws
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
    setWebsocket ({commit}) {
      return new Promise((resolve, reject) => {
        let ws = new WebSocket('ws://localhost:4444/ws')
        ws.onopen = function (e) {
          let msg = {
            'type': 'whoami',
            'data': ''
          }
          ws.send(JSON.stringify(msg))
        }

        ws.onmessage = function (e) {
          let data = JSON.parse(e.data)
          if (data.hasOwnProperty('uuid')) {
            commit('SET_USER_UUID', data.uuid)
          } else {
            console.log('Error reading uuid in whoami request')
          }
        }
        commit('SET_WEBSOCKET', ws)
        resolve()
      })
    }
  },
  getters: {
    token: state => {
      return state.token
    },
    userHandle: state => {
      return state.user.handle
    },
    userUUID: state => {
      return state.user.clientUUID
    },
    gameTibes: state => {
      return state.game.tribes
    },
    websocket: state => {
      return state.websocket
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
