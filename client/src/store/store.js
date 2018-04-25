import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new  Vuex.Store({
  state: {
    isEditProfile : false
  },
  getters: {
    isEditP (state) {
      return state.isEditProfile
    }
  },
  mutations: {

  },
  actions: {
    changeLeftSideBar (state, payload) {
      state.isEditProfile = payload
    },
  }
})

export default store
