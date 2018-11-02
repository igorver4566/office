<template>
  <v-container grid-list-xl fluid>
    <v-layout row wrap>
      <v-flex xs12>
        <v-data-table
          :headers="headers"
          :items="tasks"
          hide-actions
        >
          <template slot="items" slot-scope="props">
            <tr class="tasks_go" @click="toTask(props.item.id)">
              <td>{{ props.item.id }}</td>
              <td class="text-xs-left">{{ props.item.name }}</td>
              <td class="text-xs-center">{{ props.item.owner }}</td>
              <td class="text-xs-center">{{ props.item.manager }}</td>
              <td class="text-xs-center">{{ props.item.developer }}</td>
              <td class="text-xs-center">{{ props.item.time}} мин.</td>
              <td class="text-xs-center">{{ props.item.price }} руб.</td>
              <td class="text-xs-center">{{ props.item.subs }}</td>
              <td class="text-xs-center">
                <v-chip
                    class="white--text ml-0"
                    color="green"
                    label
                    small
                  >
                    {{ props.item.status }}
                  </v-chip>
              </td>
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
      headers: [
        {
          text: 'ID',
          align: 'left',
          sortable: true,
          value: 'id'
        },
        {
          text: 'Название задачи',
          align: 'left',
          sortable: false,
          value: 'name'
        },
        {
          text: 'Заказчик',
          align: 'center',
          sortable: false,
          value: 'owner'
        },
        {
          text: 'Менеджер',
          align: 'center',
          sortable: false,
          value: 'manager'
        },
        {
          text: 'Разработчик',
          align: 'center',
          sortable: false,
          value: 'developer'
        },
        {
          text: 'Время',
          align: 'center',
          sortable: true,
          value: 'time'
        },
        {
          text: 'Стоимость',
          align: 'center',
          sortable: true,
          value: 'price'
        },
        {
          text: 'Кол-во. задач',
          align: 'center',
          sortable: true,
          value: 'subs'
        },
        {
          text: 'Статус',
          align: 'center',
          sortable: false,
          value: 'status'
        }
      ]
    }
  },
  methods: {
    toTask (id) {
      this.$router.push('/task/' + id)
    }
  },
  computed: {
    tasks () {
      var tasks = this.$store.getters.tasks
      var subs = this.$store.getters.subTasks
      return tasks.map((val) => {
        var id = val.id
        val.time = 0
        val.price = 0
        val.subs = 0
        subs.forEach(element => {
          if (element.task_id === id) {
            val.time = val.time + element.time_dev + element.time_manage
            val.price = val.price + element.price
            val.subs += 1
          }
        })
        return val
      })
    }
  },
  created () {
    this.$store.dispatch('getAllSubTasks')
    this.$store.dispatch('getAllTasks')
  }
}
</script>

<style>
  .tasks_go {
    cursor: pointer;
  }
</style>
