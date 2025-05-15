<template>
    <div class="border">
        <router-link to="/auth">
            <h3>PLAY</h3>
        </router-link>
    </div>
</template>


<style scoped>
a {
    text-decoration: none;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: large;
    color: #0062ff;
    
}

a:hover {
    color: white;
    background-color:  #0062ff;;
    border-radius: 2rem;

}

.border {
    border-radius: 2rem;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.26);
    padding: 0;
    margin: 2rem auto;
    max-width: 15rem;
}
</style>