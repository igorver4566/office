<template>
  <v-container grid-list-xl fluid>
    <v-layout row wrap>
      <v-flex xs12 sm6 md4>
        <v-menu
          ref="menu_start"
          :close-on-content-click="false"
          v-model="menu_start"
          :nudge-right="40"
          :return-value.sync="date_start"
          lazy
          transition="scale-transition"
          offset-y
          full-width
          min-width="290px"
        >
          <v-text-field
            slot="activator"
            v-model="date_start"
            label="С"
            prepend-icon="event"
            readonly
          ></v-text-field>
          <v-date-picker v-model="date_start" no-title scrollable>
            <v-spacer></v-spacer>
            <v-btn flat color="primary" @click="menu_start = false">Отмена</v-btn>
            <v-btn flat color="primary" @click="$refs.menu_start.save(date_start)">OK</v-btn>
          </v-date-picker>
        </v-menu>
      </v-flex>
      <v-flex xs12 sm6 md4>
        <v-menu
          ref="menu_end"
          :close-on-content-click="false"
          v-model="menu_end"
          :nudge-right="40"
          :return-value.sync="date_end"
          lazy
          transition="scale-transition"
          offset-y
          full-width
          min-width="290px"
        >
          <v-text-field
            slot="activator"
            v-model="date_end"
            label="По"
            prepend-icon="event"
            readonly
          ></v-text-field>
          <v-date-picker v-model="date_end" no-title scrollable>
            <v-spacer></v-spacer>
            <v-btn flat color="primary" @click="menu_end = false">Отмена</v-btn>
            <v-btn flat color="primary" @click="$refs.menu_end.save(date_end)">OK</v-btn>
          </v-date-picker>
        </v-menu>
      </v-flex>
      <v-flex xs12 sm4 md2>
        <v-btn flat color="primary" @click="sort()">Фильтр</v-btn>
      </v-flex>
      <v-flex xs12>
        <v-data-table
          :headers="headers"
          :items="workers"
          hide-actions
        >
          <template slot="items" slot-scope="props">
            <tr class="tasks_go">
              <td>{{ props.item.id }}</td>
              <td class="text-xs-left">{{ props.item.name }}</td>
              <td class="text-xs-center">{{ props.item.time}} мин.</td>
              <td class="text-xs-center">{{ props.item.price }} руб.</td>
              <td class="text-xs-center">{{ props.item.true_time }}</td>
            </tr>
          </template>
        </v-data-table>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  data () {
    return {
      date_start: new Date().toISOString().substr(0, 10),
      date_end: new Date().toISOString().substr(0, 10),
      menu_end: false,
      menu_start: false,
      headers: [
        {
          text: 'ID',
          align: 'left',
          sortable: true,
          value: 'id'
        },
        {
          text: 'Имя',
          align: 'left',
          sortable: false,
          value: 'name'
        },
        {
          text: 'Время работы',
          align: 'center',
          sortable: false,
          value: 'time'
        },
        {
          text: 'Заработано',
          align: 'center',
          sortable: false,
          value: 'price'
        },
        {
          text: 'Фактическое время',
          align: 'center',
          sortable: false,
          value: 'true_time'
        }
      ]
    }
  },
  computed: {
    workers () {
      var workers = this.$store.getters.workers
      return workers
    }
  },
  methods: {
    sort () {
      const date = {
        dt_start: this.date_start,
        dt_end: this.date_end
      }
      this.$store.dispatch('GetAllWorkers', date)
    }
  },
  created () {
    var now = new Date()
    now.setMonth(now.getMonth() - 1)
    this.date_start = now.toISOString().substr(0, 10)
    this.date_end = new Date().toISOString().substr(0, 10)
    const date = {
      dt_start: now.toISOString().substr(0, 10),
      dt_end: new Date().toISOString().substr(0, 10)
    }
    this.$store.dispatch('GetAllWorkers', date)
  }
}
</script>
