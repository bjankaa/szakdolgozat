import { createStore } from "vuex";

import authModule from './auth/auth_idx.js'

const store = createStore({
    modules: {
        auth: authModule
    }
});

export default store;