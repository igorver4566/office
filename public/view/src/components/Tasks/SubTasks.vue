<template>
  <v-card>
    <v-card-title primary-title>
      <h3>Задачи</h3>
      <v-spacer></v-spacer>
      <v-btn 
          flat
          @click="add()"
        >
        <v-icon>add</v-icon></v-btn>
    </v-card-title>
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
          <td class="text-xs-center">
            <v-btn
              flat
              @click="edit(props.item.id)"
            >
            <v-icon>border_color</v-icon></v-btn>
          </td>
        </tr>
      </template>
      <template slot="expand" slot-scope="props">
        <v-card flat>
          <v-card-text><span class="font-weight-bold">Тех. задание</span><br><div class="white-space-pre">{{ props.item.message }}</div></v-card-text>
        </v-card>
      </template>
    </v-data-table>
    <v-dialog v-model="dialog" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">{{this.dialog_name}}</span>
        </v-card-title>
        <v-card-text>
          <v-container grid-list-md>
            <v-layout wrap>
              <v-flex xs12>
                <v-text-field name="name" v-model="name" label="Название задачи" required></v-text-field>
              </v-flex>
              <v-flex xs12 sm12 md4>
                <v-text-field name="price" v-model="price" label="Цена" required></v-text-field>
              </v-flex>
              <v-flex xs12 sm6 md4>
                <v-text-field
                  name="time_dev"
                  v-model="time_dev"
                  label="Время разработчика"
                ></v-text-field>
              </v-flex>
              <v-flex xs12 sm6 md4>
                <v-text-field
                  name="time_manage"
                  v-model="time_manage"
                  label="Время менеджера"
                ></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-textarea 
                  box
                  background-color="#fafafa"
                  name="message" 
                  label="Тех. задание" 
                  auto-grow
                  v-model="message" 
                  rows="4"
              ></v-textarea>
              </v-flex>
              <v-flex xs12>
                <v-select
                  :items="developers"
                  label="Исполнитель"
                  v-model="developer" 
                ></v-select>
              </v-flex>
            </v-layout>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" flat @click="dialog = false">Close</v-btn>
          <v-btn color="blue darken-1" flat @click="onSubmit(this.save)">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-card>
</template>

<script>
function getIdFromArray (array, el) {
  var id = 0
  array.forEach(e => {
    if (e.name === el) {
      id = e.id
    }
  })
  return id
}
export default {
  props: ['id'],
  data () {
    return {
      dialog: false,
      name: '',
      price: '',
      time_dev: '',
      time_manage: '',
      developer: '',
      message: '',
      save: 'new',
      dialog_name: 'Создать задачу',
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
        },
        {
          text: 'Редактировать',
          align: 'center',
          sortable: false,
          value: 'edit'
        }
      ]
    }
  },
  methods: {
    onSubmit (type) {
      var typeDispatch
      if (type === 'new') {
        typeDispatch = 'makeSubTask'
      } else {
        typeDispatch = 'editSubTask'
      }
      var arr = this.$store.getters.form
      const task = {
        name: this.name,
        price: parseInt(this.price),
        time_dev: parseInt(this.time_dev),
        time_manage: parseInt(this.time_manage),
        message: this.message,
        task_id: parseInt(this.id),
        user_id: parseInt(getIdFromArray(arr.developer, this.developer))
      }
      this.$store.dispatch(typeDispatch, task)
          .then(() => {
            this.dialog = false
            const id = this.id
            this.$store.dispatch('getTaskById', id)
          })
          .catch(() => {})
    },
    edit (key) {
      var tasks = this.$store.getters.task
      tasks.items.forEach(element => {
        if (element.id === key) {
          tasks = element
        }
      })
      this.name = tasks.name
      this.price = tasks.price
      this.time_dev = tasks.time_dev
      this.time_manage = tasks.time_manage
      this.developer = tasks.user
      this.message = tasks.message
      this.save = 'edit'
      this.dialog_name = 'Редактировать задачу'
      this.dialog = true
    },
    add () {
      this.dialog = true
      this.name = ''
      this.price = ''
      this.time_dev = ''
      this.time_manage = ''
      this.developer = ''
      this.message = ''
      this.save = 'new'
      this.dialog_name = 'Создать задачу'
    }
  },
  computed: {
    tasks () {
      var task = this.$store.getters.task
      return task.items == null ? [] : task.items
    },
    developers () {
      var arr
      if (this.$store.getters.form == null) {
        this.$store.dispatch('getFormFields')
        .then((r) => {
          arr = this.$store.getters.form
        })
        .catch(() => {})
      } else {
        arr = this.$store.getters.form
      }
      return arr ? arr.developer.map((el) => { el = el.name; return el }) : ['Ошибка при загрузке']
    }
  }
}
</script>
