<template>
  <v-card>
    <v-card-title primary-title><h3>Задачи</h3></v-card-title>
    <v-data-table
      :headers="headers"
      :items="tasks"
      hide-actions
    >
      <template slot="items" slot-scope="props">
        <tr @click="props.expanded = !props.expanded">
          <td>{{ props.item.id }}</td>
          <td class="text-xs-left">{{ props.item.name }}</td>
          <td class="text-xs-center">{{ props.item.user }}</td>
          <td class="text-xs-center">{{ props.item.time_dev + props.item.time_manage }} мин.</td>
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
        </tr>
      </template>
      <template slot="expand" slot-scope="props">
        <v-card flat>
          <v-card-text><span class="font-weight-bold">Тех. задание</span><br>{{ props.item.message }}</v-card-text>
        </v-card>
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
          text: 'Исполнитель',
          align: 'center',
          sortable: false,
          value: 'user'
        },
        {
          text: 'Время',
          align: 'center',
          sortable: true,
          value: 'time_dev'
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
      var task = this.$store.getters.task
      return task.items
    }
  }
}
</script>
