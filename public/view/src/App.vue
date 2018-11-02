<template>
  <v-app>
  <v-navigation-drawer 
    app
    v-model="drawer"
  >
  <v-toolbar flat>
      <v-list>
        <v-list-tile>
          <v-list-tile-title class="title">
            Меню
          </v-list-tile-title>
        </v-list-tile>
      </v-list>
    </v-toolbar>

    <v-divider></v-divider>
  <v-list dense>
      <v-list-tile
        v-for="item in items"
        :key="item.title"
        :to="item.url"
        @click="drawer = !drawer"
      >
        <v-list-tile-content>
          <v-list-tile-title>{{ item.title }}</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </v-navigation-drawer>
  <v-toolbar app>
      <v-toolbar-side-icon
        @click="drawer = !drawer"
      ></v-toolbar-side-icon>
      <v-toolbar-title>
        <router-link to="/" tag="span" class="pointer">SoftSetters</router-link>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <v-btn 
          v-for="link in links"
          :key="link.title"
          :to="link.url"
          flat
        >
        <v-icon>{{link.title}}</v-icon></v-btn>
      </v-toolbar-items>
  </v-toolbar>
  <v-content>
      <router-view></router-view>
  </v-content>
  <v-snackbar
      v-model="snackbar"
      :color="color"
      top
    >
      {{ error }}
      {{ ok }}
      <v-btn
        dark
        flat
        @click="snackbar = false"
      >
        Закрыть
      </v-btn>
    </v-snackbar>
</v-app>
</template>

<script>
export default {
  data () {
    return {
      snackbar: false,
      drawer: false,
      color: 'error',
      links: [
        {title: 'exit_to_app', url: '/login'},
        {title: 'face', url: '/register'}
      ],
      items: [
        {title: 'Создать задачу', url: '/make_task'},
        {title: 'Все задачи', url: '/tasks'}
      ]
    }
  },
  computed: {
    error () {
      if (this.$store.getters.error != null) {
        this.color = 'error'
        this.snackbar = false
        this.snackbar = true
        return this.$store.getters.error
      }
    },
    ok () {
      if (this.$store.getters.ok != null) {
        this.color = 'success'
        this.snackbar = false
        this.snackbar = true
        return this.$store.getters.ok
      }
    }
  }
}
</script>

<style scoped>
  .pointer {
    cursor: pointer;
  }
</style>


