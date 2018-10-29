<template>
  <v-container grid-list-xl fluid>
    <v-layout row wrap>
      <v-flex xs12 md6>
        <div>
          <span>{{ weather }} &#176;</span>
        </div>
        <br>
        <h2>{{ hello }}</h2>
      </v-flex>
      <v-flex xs12 md6>
        <Notify />
      </v-flex>
      <v-flex xs12 md7>
        <TaskGrid />
      </v-flex>
      <v-flex xs12 md5>
        <SubTaskGrid />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import Notify from './Notify'
import TaskGrid from './TaskGrid'
import SubTaskGrid from './SubTaskGrid'

export default {
  components: {
    Notify,
    TaskGrid,
    SubTaskGrid
  },
  data () {
    return {}
  },
  computed: {
    hello () {
      var date = new Date()
      var hour = date.getHours()
      var helloText = ''
      if (hour >= 6 && hour < 12) {
        helloText = 'Доброе утро'
      } else if (hour >= 12 && hour < 17) {
        helloText = 'Добрый день'
      } else if (hour >= 17 && hour <= 23) {
        helloText = 'Добрый вечер'
      } else if (hour >= 0 && hour < 6) {
        helloText = 'Доброй ночи'
      }

      return helloText + ', ' + this.$store.getters.userName
    },
    weather () {
      return this.$store.getters.location + ', ' + this.$store.getters.weather
    }
  },
  created () {
    this.$store.dispatch('getSubTasks')
    this.$store.dispatch('getTasks')
  }
}
</script>



