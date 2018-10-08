<template>
  <v-app>
  <v-navigation-drawer 
    app
    v-model="drawer"
  ></v-navigation-drawer>
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
</v-app>
</template>

<script>
export default {
  data () {
    return {
      drawer: false,
      links: [
        {title: 'exit_to_app', url: '/login'},
        {title: 'face', url: '/register'}
      ]
    }
  }
}
</script>

<style scoped>
  .pointer {
    cursor: pointer;
  }
</style>


