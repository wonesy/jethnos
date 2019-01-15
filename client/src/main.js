// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuex from 'vuex'
import App from './App'
import router from './router'
import './../node_modules/bulma/css/bulma.css'
import VueResource from 'vue-resource'

Vue.use(VueResource)
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
    fetchToken ({commit}) {
      return new Promise((resolve, reject) => {
        Vue.http.get('http://localhost:4444/gettoken')
          .then((response) => {
            commit('SET_TOKEN', response.body['token'])
            resolve()
          })
          .catch(error => {
            console.log(error.statusText)
            reject(error.statusText)
          })
      })
    }
  },
  getters: {
    token: state => {
      return state.token
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

  console.log('yoooooooooooo')
  console.log(token)

  if (token) {
    return { 'Authorization': 'Bearer ' + token }
  } else {
    return {}
  }
}
