<template>
  <v-card>
    <v-card-title primary-title><h3>Задачи</h3></v-card-title>
    <v-data-table
      :headers="headers"
      :items="desserts"
      hide-actions
    >
      <template slot="items" slot-scope="props">
        <td>{{ props.item.id }}</td>
        <td class="text-xs-left">{{ props.item.name }}</td>
        <td class="text-xs-center">{{ props.item.manager }}</td>
        <td class="text-xs-center">{{ props.item.price }} руб.</td>
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
          text: 'Стоимость',
          align: 'center',
          sortable: true,
          value: 'price'
        }
      ],
      desserts: [
        {
          value: false,
          id: '1',
          name: 'Название задачи 1',
          manager: 'Администратор',
          price: 1000
        },
        {
          value: false,
          id: '2',
          name: 'Название задачи 2',
          manager: 'Администратор',
          price: 2000
        }
      ]
    }
  }
}
</script>


