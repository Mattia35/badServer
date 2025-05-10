<script>
export default{
    data: function(){
        return{
            errormsg: null,
            password: "",
            username: "",
        };
    },
    methods: {
        async login() {
            this.errormsg = null;
            try {
                // send the username to the server, if the user doesnt't exist, the server will create a new user
                let response = await this.$axios.put('/login', {
                    username: this.username,
                    password: this.password,
                });
                
                // save the user data to the local storage
                sessionStorage.username = response.data.username;
                sessionStorage.token = response.data.token;
                sessionStorage.session = response.data.session;

                // redirect to the home page
                this.$router.push(`/${sessionStorage.username}`);

                // emit the login success event
                this.$emit('login-success');

            } catch (e) {
                this.errormsg = e.toString();
            };
        },
    },
}
</script>

<template>
    <div class="body">
        <div class = "login-container">
            <ErrorMsg v-if="errormsg" :msg="errormsg"></Errormsg>
            <form @submit.prevent="login">
                <h1>Login</h1>
                <div class="input-container">
                    <input type="text" v-model="username" placeholder="Write your username">
                    <input type="text" v-model="password" placeholder="Write your password">
                    <button type="submit">Login</button>
                </div>
            </form>
        </div>
    </div>
</template>

<style>
    .body {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
    }

    .login-container {
        background: rgb(248 249 250);
        padding: 50px;
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.6);
        width: 500px;
        height: 200px;
        text-align: center;
    }

    .input-container {
        position: relative;
        margin-top: 20px;
    }

    .input-container input {
        width: 100%;
        padding: 10px;
        font-size: 16px;
        border: 1px solid #ccc;
        border-radius: 5px;
        outline: none;
    }

    .input-container label {
        position: absolute;
        top: 50%;
        left: 10px;
        transform: translateY(-50%);
        color: #aaa;
        font-size: 16px;
        pointer-events: none;
        transition: all 0.3s ease;
    }

    .input-container input:focus + label,
    .input-container input:not(:placeholder-shown) + label {
        top: 5px;
        left: 10px;
        font-size: 12px;
        color: #333;
    }

    .input-container button {
        width: 50%;
        padding: 10px;
        margin-top: 20px;
        background: #12B886;
        color: #fff;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        font-size: 16px;
        transition: all 0.3s ease;
    }

    .input-container button:hover {
        background: #555;
    }
</style>