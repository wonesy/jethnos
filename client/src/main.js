// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuex from 'vuex'
import VueLodash from 'vue-lodash'
import App from './App'
import router from './router'
import './../node_modules/bulma/css/bulma.css'
import './../node_modules/lodash'
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
    }
  },
  mutations: {
    SET_TOKEN (state, token) {
      state.token = token
    },
    SET_USER_HANDLE (state, handle) {
      state.user.handle = handle
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
    }
  },
  getters: {
    token: state => {
      return state.token
    },
    userHandle: state => {
      return state.user.handle
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

export function authHeader () {
  // return authorization header with jwt token
  let token = store.getters.token

  if (token) {
    return { 'Authorization': 'Bearer ' + token }
  } else {
    return {}
  }
}
