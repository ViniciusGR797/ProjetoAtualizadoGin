<template>
  <q-page class="flex flex-center" style="background-image: linear-gradient(to right, #ff6d0f, #8222ff);">
    <q-card style="width: 300px; height: auto; display: flex; flex-direction: column; align-items: center;">

      <q-img src="../assets/logo.png" ratio="1" width="150px" style="margin-top: 15px;"/>

      <q-card-section class="q-pt-none">
        <div class="text-h5">Login</div>
      </q-card-section>

      <q-card-section class="q-pt-none full-width">
        <q-input outlined v-model="username" @keyup="handleKeypress" label="Usuário" />
      </q-card-section>

      <q-card-section class="q-pt-none full-width">
        <q-input v-model="password" @keyup="handleKeypress" outlined :type="isPwd ? 'password' : 'text'" label="Senha">
        <template v-slot:append>
          <q-icon
            :name="isPwd ? 'visibility_off' : 'visibility'"
            class="cursor-pointer"
            @click="showHide"
          />
        </template>
      </q-input>
      </q-card-section>

      <q-card-section>
        <q-btn @click="handleLogin" color="primary" icon="login" label="Entrar" />
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useQuasar } from 'quasar'
import { postData } from '../boot/axios'

export default defineComponent({
  name: 'LoginComponent',

  setup () {
    const $q = useQuasar()
    const username = ref('')
    const password = ref('')
    const router = useRouter()
    const isPwd = ref(true)

    function showHide () {
      isPwd.value = !isPwd.value
      return isPwd.value
    }

    function handleKeypress (e) {
      if (e.key === 'Enter') {
        if (username.value.length > 3 && password.value.length >= 6) {
          handleLogin()
        } else {
          $q.notify({
            position: 'top-right',
            type: 'warning',
            message: 'É necessário informa [usuário com no minimo 3 digitos] e [senha com no minimo 6 digitos] para fazer login.'
          })
        }
      }
    }

    function handleLogin () {
      const userCredentials = { username: username.value, password: password.value }
      const url = 'api/v1/user/login'
      postData(url, userCredentials, {}).then((response) => {
        localStorage.setItem('autenticated', JSON.stringify(response.data))
        router.push({ name: 'Index' })
        $q.notify({
          position: 'top-right',
          type: 'positive',
          message: `Bem vindo, ${username.value}`
        })
      }).catch((error) => {
        console.log('Login', error)
        $q.notify({
          position: 'top-right',
          type: 'negative',
          message: `Erro ao fazer login, verifique o usuário e senha digitados. [${error}]`
        })
      })
    }

    return {
      isPwd,
      username,
      password,
      showHide,
      handleKeypress,
      handleLogin
    }
  }

})
</script>
