<template>
    <base-card>
        <h2>EZ A FRONTPAGE</h2>
    </base-card>
    <div v-if="!isLoggedIn">
        <logged-out></logged-out>
    </div>
    <div v-else>
        <logged-in></logged-in>
    </div>
</template>


<script>

import LoggedIn from '../components/frontpage/LoggedIn.vue';
import LoggedOut from '../components/frontpage/LoggedOut.vue';

export default {
    components: {
        LoggedOut,
        LoggedIn
    },
    computed: {
        isLoggedIn() {
            const login = this.$store.getters.loggedIn;
            return login;
        }
    }
}

</script>

<style scoped>

</style>