<template>
    <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md5>
            <v-card class="elevation-12">
              <v-toolbar dark color="primary">
                <v-toolbar-title>Вход</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <span class="red-span" >{{this.submit}}</span>
                <v-form ref="form" v-model="valid" lazy-validation>
                  <v-text-field 
                    prepend-icon="person" 
                    name="login" 
                    :rules="loginRules"
                    label="Логин" 
                    type="text" 
                    v-model="login" 
                    required
                  ></v-text-field>
                  <v-text-field 
                    prepend-icon="person" 
                    name="email" 
                    :rules="emailRules"
                    label="Эл. почта" 
                    type="text" 
                    v-model="email" 
                    required
                  ></v-text-field>
                  <v-text-field 
                    prepend-icon="lock" 
                    name="password" 
                    :rules="passRules"
                    label="Пароль" 
                    type="password" 
                    v-model="password" 
                    required
                  ></v-text-field>
                  <v-text-field 
                    id="password" 
                    prepend-icon="lock" 
                    name="confirmPassword" 
                    :rules="confirmPasswordRules"
                    label="Подтвердите пароль" 
                    type="password" 
                    v-model="confirmPassword" 
                    required
                  ></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn 
                  color="primary" 
                  @click="onSubmit"
                  :disabled="!valid"
                >Зарегистрироваться</v-btn>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
</template>

<script>
  import {Registration} from '@/api/hostsetting.js'
  export default {
    data () {
      return {
        login: '',
        password: '',
        confirmPassword: '',
        email: '',
        valid: false,
        submit: '',
        loginRules: [
          v => !!v || 'Login is required'
        ],
        passRules: [
          v => !!v || 'Password is required',
          v => (v && v.length) >= 6 || 'Пароль должен иметь больше 6 символов'
        ],
        confirmPasswordRules: [
          v => !!v || 'Password is required',
          v => v === this.password || 'Пароли должны совпадать'
        ],
        emailRules: [
          v => !!v || 'E-mail is required',
          v => /.+@.+/.test(v) || 'E-mail must be valid'
        ]
      }
    },
    methods: {
      onSubmit () {
        if (this.$refs.form.validate()) {
          const user = {
            login: this.login,
            password: this.password,
            email: this.email
          }
          Registration(user.login, user.password, user.email)
          .then((res) => {
            this.submit = res
          })
        }
      }
    }
  }
</script>

<style>
  .red-span {
    color: red;
  }
</style>
