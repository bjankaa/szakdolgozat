export default {
    setUser(state, payload) {
        state.token = payload.token;
        state.userId = payload.userId;
        state.loggedIn = true;
    },
    setUserLogout() {
        state.token = null;
        state.userId = null
        state.loggedIn = false;
    }

}