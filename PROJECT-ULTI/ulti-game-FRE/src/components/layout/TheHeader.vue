<template>
  <header>
    <nav>
      <h1>ULTI</h1>
      <img src="../../assets/cards.png" alt="cards" width="50" height="50">
    </nav>
    <ul>
      <li v-if="!isLoggedIn"><router-link to="/auth">Log In</router-link></li>
      <li v-else><button @click="logOut">Log Out</button></li>
    </ul>
  </header>
</template>

<script>

export default {
  computed: {
    isLoggedIn() {
      return this.$store.getters.loggedIn;
    }
  },
  methods: {
    logOut() {
      this.$store.dispatch('logout');
    }
  }
}


</script>


<style scoped>
header {
  background-color: #0062ff;
  color: white;
  width: 100%;
  height: 5rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

header nav {
  width: 90%;
  margin: 2rem;
  display: flex;
  justify-self: center;
}

header ul {
  display: flex;
  justify-self: space-around;
  margin: 2rem;
  padding: 0;
}

li {
  list-style: none;
  display: flex;

}

header a ,
header button{
  color: #0062ff;
  background: white;
  padding: 0.25rem 3rem;
  border: 1px solid transparent;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 0.75rem;
  border-radius: 1rem;
  white-space: nowrap;
}

a:hover {
  border: 0.1rem solid #002053;
}
</style>