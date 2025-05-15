export default {
    logout(context) {
        context.commit('setUserLogout');
    },
    async login(context, payload) {
        // const url = "http://localhost:3000/auth";

        const res = await fetch('/api/auth', {
            method: 'POST',
            body: JSON.stringify({
                email: payload.email,
                password: payload.password,
                state: "login"
            })
        })

        const resData = await res.json();

        if (!res.ok) {
            const error = new Error(resData.message || 'Failed to authenticate!');
            throw error;
        }

        localStorage.setItem('token', resData.idToken);
        localStorage.setItem('userId', resData.localId);

        context.commit('setUser', {
            token: resData.idToken,
            userId: resData.localId,
        });

    },
    async signup(context, payload) {

        // const url = 'http://localhost:5173/auth';

        const res = await fetch('/api/auth', {
            method: 'POST',
            body: JSON.stringify({
                name: payload.name,
                email: payload.email,
                password: payload.password,
                state: "signup"
            })
        })
    }
}

