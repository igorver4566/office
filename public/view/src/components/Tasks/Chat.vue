<template>
  <v-card class="chat-room"> 
    <v-toolbar card dense flat class="white chat-room--toolbar" light>
      <v-spacer></v-spacer>
      <v-toolbar-title> <h4>Чат</h4></v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>    
    <vue-perfect-scrollbar class="chat-room--scrollbar grey lighten-5">
      <v-card-text class="chat-room--list pa-3">
        <template v-for="(item, index) in chat.messages">
          <div v-bind:class="[ index % 2 == 0 ? 'reverse' : '']" class="messaging-item layout row my-4" :key="index">
            <div class="messaging--body layout column mx-2">
              <p :value="true" v-bind:class="[ index % 2 == 0 ? 'primary white--text' : 'white']" class="pa-2">
                {{item.text}}
              </p>
              <div class="caption px-2 text--secondary">20:10</div>
            </div>
            <v-spacer></v-spacer>
          </div>
        </template>
      </v-card-text>  
    </vue-perfect-scrollbar>
    <v-card-actions>
      <v-text-field 
        full-width
        flat
        clearable
        solo
        append-icon="send"
        label="Type some message here">
        <v-icon slot="append-icon">send</v-icon>
        <v-icon slot="append-icon" class="mx-2">photo</v-icon>
        <v-icon slot="append-icon">face</v-icon>
      </v-text-field>
    </v-card-actions>
  </v-card>
</template>

<script>
import VuePerfectScrollbar from 'vue-perfect-scrollbar'
export default {
  components: {
    VuePerfectScrollbar
  },
  computed: {
    chat () {
      let chatOrigin = {
        title: 'Chat',
        users: [],
        messages: [
          {
            text: 'text'
          },
          {
            text: 'text2'
          },
          {
            text: 'Сообщение'
          },
          {
            text: 'Ответ'
          },
          {
            text: 'Сообщение'
          }
        ]
      }
      return chatOrigin
    }
  }
}
</script>

<style>
  .chat-room--scrollbar {
    height: 400px;
  }
</style>
