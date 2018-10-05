<template>
    <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md5>
            <v-card class="elevation-12">
              <v-toolbar dark color="primary">
                <v-toolbar-title>Вход</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <v-form ref="form" v-model="valid" validation>
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
                    prepend-icon="lock" 
                    name="password" 
                    :rules="passRules"
                    label="Пароль" 
                    type="password" 
                    v-model="password" 
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
                >Login</v-btn>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
</template>

<script>
  export default {
    data () {
      return {
        login: '',
        password: '',
        valid: false,
        loginRules: [
          v => !!v || 'Login is required'
        ],
        passRules: [
          v => !!v || 'Password is required',
          v => (v && v.length) >= 6 || 'Password must be equal or more than 6 characters'
        ]
      }
    },
    methods: {
      onSubmit () {
        if (this.$refs.form.validate()) {
          const user = {
            login: this.login,
            password: this.password
          }
          console.log(user)
        }
      }
    }
  }
</script>
