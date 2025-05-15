import mutations from "./mutations";
import actions from "./actions";
import getters from "./getters";



export default {
    state() {
        return {
            loggedIn: false,
            token: null,
            userId: null
        };
    },
    getters,
    actions,
    mutations
}