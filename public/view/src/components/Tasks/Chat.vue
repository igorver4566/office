<template>
  <v-card class="chat-room"> 
    <v-toolbar card dense flat class="white chat-room--toolbar" light>
      <v-spacer></v-spacer>
      <v-toolbar-title> <h4>Чат</h4></v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>    
    <vue-perfect-scrollbar ref="ps" class="chat-room--scrollbar grey lighten-5">
      <v-card-text class="chat-room--list pa-3">
        <template v-for="(item, index) in chat">
          <div class="messaging-item sm12 row my-4" :key="index">
            <div class="messaging--body sm12 column mx-2">
              <p :value="true" class="pa-3 white">
                <b>{{item.username ? item.username : 'Slack'}}</b>:<br>
                {{item.text}}
              </p>
              <div class="caption px-2 text--secondary">{{item.ts}}</div>
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
        v-model="msg"
        label="Сообщение...">
      </v-text-field>
      <v-btn flat @click="submitMsg()"><v-icon>send</v-icon></v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import VuePerfectScrollbar from 'vue-perfect-scrollbar'
function onlineChat (store) {
  setTimeout(function run () {
    var chatName = store.getters.task.task.name_slack
    store.dispatch('getMessages', chatName)
    setTimeout(run, 5000)
  }, 5000)
}
export default {
  data () {
    return {
      msg: ''
    }
  },
  components: {
    VuePerfectScrollbar
  },
  methods: {
    submitMsg () {
      const message = {
        chat: this.$store.getters.task.task.name_slack,
        msg: {
          msg: this.msg,
          name: this.$store.getters.userName
        }
      }
      this.$store.dispatch('sendMsg', message).then(() => {
        var chatName = this.$store.getters.task.task.name_slack
        this.$store.dispatch('getMessages', chatName)
        this.msg = ''
      })
    }
  },
  computed: {
    chat () {
      var chatName = this.$store.getters.task.task.name_slack
      var messages = []
      if (chatName !== '' && chatName !== undefined) {
        this.$store.dispatch('getMessages', chatName)
        messages = this.$store.getters.messages
        messages = messages.map((val) => {
          var ts = val.ts
          var date = new Date(ts * 1000)
          var day = date.getDate()
          var mounth = date.getMonth()
          var hours = date.getHours()
          var minutes = '0' + date.getMinutes()
          val.ts = day + '.' + mounth + ' ' + hours + ':' + minutes.substr(-2)
          return val
        })
      }
      return messages
    }
  },
  created () {
    onlineChat(this.$store)
  }
}
</script>

<style>
  .chat-room--scrollbar {
    height: 400px;
  }
</style>
