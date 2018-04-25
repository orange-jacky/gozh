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
    changeLeftSideBar (state, payload) {
      state.isEditProfile = payload
    },
  },
  actions: {

  }
})

export default store
