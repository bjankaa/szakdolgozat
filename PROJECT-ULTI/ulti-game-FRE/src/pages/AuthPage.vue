<template>
    <div>
        <base-card>
            <form @submit.prevent="submitForm">
                <div v-if=singupInfo class="form.control" :class="{ invalid: !isFormValid }">
                    <label for="name">Name</label>
                    <input type="text" id="name" v-model.trim="name">
                </div>
                <div class="form.control" :class="{ invalid: !isFormValid }">
                    <label for="email">E-mail</label>
                    <input type="email" id="email" v-model.trim="email">
                </div>
                <div class="form.control" :class="{ invalid: !isFormValid }">
                    <label for="password">Password</label>
                    <input type="password" id="password" v-model.trim="password">
                </div>
                <p v-if="!isFormValid">Please fill all the required areas</p>
                <button>{{ submitButtonCaption }}</button>
                <button @click.prevent="switchButtonFunc" class="link">{{ otherButtonCaption }}</button>
            </form>
        </base-card>
    </div>
</template>

<script>

export default {

    data() {
        return {
            name: "",
            email: "",
            password: "",
            mode: 'login',
            isFormValid: true
        }
    },
    computed: {
        singupInfo() {
            if (this.mode === 'signup') {
                return true;
            } else {
                return false;
            }
        },
        submitButtonCaption() {
            if (this.mode === 'login') {
                return 'Log in';
            } else {
                return 'Sign up';
            }
        },
        otherButtonCaption() {
            if (this.mode === 'login') {
                return 'Sign up here';
            } else {
                return 'Log in here';
            }
        }
    },
    methods: {
        switchButtonFunc() {
            if (this.mode === 'login') {
                console.log(this.mode);
                return this.mode = 'signup';
            } else {
                console.log(this.mode);
                return this.mode = 'login';
            }

        },
        validation() {
            this.isFormValid = true;
            console.log(this.isFormValid);
            if (this.mode === 'signup') {
                if (this.name === '' || this.email === '' || !this.email.includes('@') || this.password.lenght < 6) {
                    this.isFormValid = false;
                    console.log(this.isFormValid);
                    return false;
                } else {
                    return true;
                }
            } else {
                if (this.email === '' || !this.email.includes('@') || this.password.lenght === '') {
                    this.isFormValid = false;
                    console.log(this.isFormValid);
                    return false;
                } else {
                    return true;
                }
            }
        },
        async submitForm() {
            if (!this.validation()) {
                return;
            }

            const loginPayload = {
                email: this.email,
                password: this.password
            };
            const signupPayload = {
                name: this.name,
                email: this.email,
                password: this.password
            };

            try {
                if (this.mode === 'login') {
                    await this.$store.dispatch('login', loginPayload);
                } else {
                    await this.$store.dispatch('signup', signupPayload);
                }
                const redirectUrl = '/' + (this.$route.query.redirect || 'frontpage');
                this.$router.replace(redirectUrl);
            } catch (err) {
                this.error = err.message || 'Failed to authenticate.';
            }

        }
    }

}


</script>

<style scoped>
form {
    margin: 1rem;
    padding: 1rem;
}

.form-control {
    margin: 0.5rem 0;
}

label {
    font-weight: bold;
    margin-bottom: 0.5rem;
    display: block;
}

input,
textarea {
    display: block;
    width: 100%;
    font: inherit;
    border: 1px solid #ccc;
    padding: 0.15rem;
    border-radius: 2rem;
}

input:focus,
textarea:focus {
    border-color: #0062ff;
    background-color: transparent;
    outline: none;
    border-radius: 2rem;
}

a,
button {
    text-decoration: none;
    font-size: large;
    color: white;
    border: 1px solid transparent;
    border-color: #0062ff;
    margin-top: 0.5rem;
    border-radius: 2rem;
    background: #0062ff;
    max-width: 15rem;
    display: inline-block;

}

a:hover,
button:hover,
button:active {
    color: #0062ff;
    background-color: white;
    border-color: transparent;
    border-radius: 2rem;

}

.link {
    margin-left: 1rem;
    background-color: transparent;
    border-color: transparent;
    color: #0062ff
}

.link:hover {
    color: #0062ff;
    border-color: #0062ff;
}

.invalid label {
    color: red;
}

.invalid input,
.invalid textarea {
    border: 1px solid red;
}
</style>