<template>
  <v-container grid-list-xl fluid>
    <v-layout row wrap>
      <v-flex xs12>
        <v-card>
          <v-img height="200px" src="https://images.unsplash.com/photo-1508615070457-7baeba4003ab?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6c58016a8848bf447257e8d9ea378f92&auto=format&fit=crop&w=1500&q=80" >
            <v-container fill-height>
              <v-layout align-center>
                <v-flex text-xs-center>
                  <h3 class="display-3 white--text">{{ task.task.name }}</h3>
                </v-flex>
              </v-layout>
            </v-container>
          </v-img>

          <v-card-title primary-title>
            <v-container grid-list-xl fluid>
              <v-layout row wrap>
                <v-flex xs6 lg2>
                  <p class="title font-weight-bold">Доступы</p>
                  <p class="white-space-pre">{{ task.task.access }}</p>
                </v-flex>
                <v-flex xs6 lg2>
                  <p class="title font-weight-bold">Менеджер</p>
                  <p>{{ task.task.manager }}</p>
                </v-flex>
                <v-flex xs6 lg2>
                  <p class="title font-weight-bold">Заказчик</p>
                  <p>{{ task.task.owner }}</p>
                </v-flex>
                <v-flex xs6 lg2>
                  <p class="title font-weight-bold">Исполнитель</p>
                  <p>{{ task.task.developer }}</p>
                </v-flex>
                <v-flex xs6 lg2>
                  <p class="title font-weight-bold">Теги</p>
                  <p>{{ task.task.tags }}</p>
                </v-flex>
                <v-flex xs6 lg2>
                  <p class="title font-weight-bold">Статус</p>
                  <p>{{ task.task.status }}</p>
                </v-flex>
                <v-flex xs12><v-divider></v-divider></v-flex>
                <v-flex xs12>
                  <p class="title font-weight-bold">Описание</p>
                  <p class="white-space-pre">{{ task.task.message }}</p>
                </v-flex>
              </v-layout>
            </v-container>
          </v-card-title>
        </v-card>
      </v-flex>
      <v-flex xs12>
        <SubTasks :id="this.id" />
      </v-flex>
      <v-flex xs12>
        <Chat />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import SubTasks from './SubTasks'
import Chat from './Chat'
export default {
  components: {
    SubTasks,
    Chat
  },
  props: ['id'],
  data () {
    return {}
  },
  computed: {
    task () {
      return this.$store.getters.task
    }
  },
  created () {
    const id = this.id
    this.$store.dispatch('getTaskById', id)
  }
}
</script>

<style>
  .white-space-pre {
    white-space: pre-wrap;
  }
</style>

