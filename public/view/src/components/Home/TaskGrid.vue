<template>
  <v-card>
    <v-card-title primary-title><h3>Проекты</h3></v-card-title>
    <v-data-table
      :headers="headers"
      :items="tasks"
      hide-actions
    >
      <template slot="items" slot-scope="props">
        <td>{{ props.item.id }}</td>
        <td class="text-xs-left">{{ props.item.name }}</td>
        <td class="text-xs-center">{{ props.item.manager }}</td>
        <td class="text-xs-center">{{ props.item.time}} мин.</td>
        <td class="text-xs-center">{{ props.item.price }} руб.</td>
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
      </template>
    </v-data-table>
  </v-card>
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
          text: 'Менеджер',
          align: 'center',
          sortable: false,
          value: 'manager'
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
          text: 'Статус',
          align: 'center',
          sortable: false,
          value: 'status'
        }
      ]
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
        subs.forEach(element => {
          if (element.task_id === id) {
            val.time = val.time + element.time_dev + element.time_manage
            val.price = val.price + element.price
          }
        })
        return val
      })
    }
  },
  created () {
    this.$store.dispatch('getTasks')
  }
}
</script>


