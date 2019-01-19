<template>
  <v-card>
    <v-card-title primary-title>
      <h3>Задачи</h3>
      <v-spacer></v-spacer>
      <span v-show="this.$store.getters.userName === 'Администратор'">Сумма за неделю: {{weakSum}} руб.</span>
      <v-btn 
          flat
          @click="add()"
        >
        <v-icon>add</v-icon></v-btn>
    </v-card-title>
    <v-data-table
      :headers="headers"
      :items="tasks"
      class="elevation-1"
    >
      <template slot="items" slot-scope="props">
        <tr :key="props.item.id">
          <td @click="props.expanded = !props.expanded">{{ props.item.id }}</td>
          <td  @click="props.expanded = !props.expanded">{{ props.item.priority }}</td>
          <td class="text-xs-center" @click="props.expanded = !props.expanded">{{ props.item.dt_create.replace(/[A-Z]/g,' ') }}</td>
          <td class="text-xs-left" @click="props.expanded = !props.expanded">{{ props.item.name }}</td>
          <td class="text-xs-center">{{ props.item.user }}</td>
          <td class="text-xs-center">{{ props.item.time_dev + props.item.time_manage }} мин.</td>
          <td class="text-xs-center">{{ props.item.price }} руб.</td>
          <td class="text-xs-center">
            <v-btn
              flat
              @click="show(props.item.id)"
            >
              <v-chip
                  class="white--text ml-0"
                  :color="greenStatus(props.item.status)"
                  label
                  small
                >
                {{ props.item.status }}
              </v-chip>
            </v-btn>
            <v-select
                :items="statuses"
                label="Статус"
                :satatus_id="props.item.id"
                :value="props.item.status"
                v-show="status_show === props.item.id"
                @change="changeStatus"
              ></v-select>
          </td>
          <td>
            <v-btn
              flat
              @click="showAddTime(props.item.id)"
            >
              {{ props.item.true_time }}
            </v-btn>
            <v-layout align-center justify-center v-show="time_show === props.item.id">
              <v-flex sm6>
                <v-text-field
                  name="true_time"
                  v-model="true_time"
                ></v-text-field>
              </v-flex>
              <v-flex sm6>
                 <v-icon @click="addTime()">add_box</v-icon>
              </v-flex>
            </v-layout>
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
                  @input="devChange"
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
                <v-subheader class="pl-0">Приоритет</v-subheader>
                <v-slider
                  v-model="priority"
                  thumb-label="always"
                  max="5"
                ></v-slider>
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
          <v-btn color="blue darken-1" flat @click="dialog = false">Закрыть</v-btn>
          <v-btn color="blue darken-1" flat @click="onSubmit()">Сохранить</v-btn>
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
      task_id: 0,
      dialog: false,
      status: '',
      status_show: false,
      time_show: false,
      true_time: 0,
      name: '',
      price: '',
      time_dev: '',
      priority: 0,
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
          text: 'Приоритет',
          align: 'left',
          sortable: true,
          value: 'priority'
        },
        {
          text: 'Дата',
          align: 'center',
          sortable: true,
          value: 'dt_create'
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
          text: 'Факт. время',
          align: 'center',
          sortable: false,
          value: 'true_time'
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
    onSubmit () {
      var type = this.save
      var id = this.task_id
      var typeDispatch
      if (type === 'new') {
        typeDispatch = 'makeSubTask'
      } else {
        typeDispatch = 'editSubTask'
      }
      var arr = this.$store.getters.form
      const task = {
        id: id === 0 ? 0 : id,
        name: this.name,
        price: parseInt(this.price),
        time_dev: parseInt(this.time_dev),
        time_manage: parseInt(this.time_manage),
        message: this.message,
        task_id: parseInt(this.id),
        user_id: parseInt(getIdFromArray(arr.developer, this.developer)),
        priority: this.priority
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
      this.task_id = key
      this.name = tasks.name
      this.price = tasks.price
      this.time_dev = tasks.time_dev
      this.time_manage = tasks.time_manage
      this.developer = tasks.user
      this.message = tasks.message
      this.priority = tasks.priority
      this.save = 'edit'
      this.dialog_name = 'Редактировать задачу'
      this.dialog = true
    },
    add () {
      this.dialog = true
      this.task_id = 0
      this.priority = 0
      this.name = ''
      this.price = ''
      this.time_dev = ''
      this.time_manage = ''
      this.developer = ''
      this.message = ''
      this.save = 'new'
      this.dialog_name = 'Создать задачу'
    },
    changeStatus (e) {
      var arr = this.$store.getters.form
      const task = {
        id: parseInt(this.status_show),
        status_id: parseInt(getIdFromArray(arr.status, e))
      }
      this.$store.dispatch('changeStatusSubTask', task)
          .then(() => {
            this.dialog = false
            const id = this.id
            this.status_show = 0
            this.$store.dispatch('getTaskById', id)
          })
          .catch(() => {})
    },
    addTime () {
      const task = {
        id: parseInt(this.time_show),
        true_time: parseInt(this.true_time)
      }
      this.$store.dispatch('addTrueTimeSubTask', task)
          .then(() => {
            this.dialog = false
            const id = this.id
            this.time_show = 0
            this.$store.dispatch('getTaskById', id)
          })
          .catch(() => {})
    },
    greenStatus (key) {
      switch (key) {
        case 'Новая':
          return 'yellow'
        case 'В работе':
          return 'blue'
        case 'Отменена':
          return 'red'
        case 'Выполнена':
          return 'green'
      }
    },
    show (key) {
      if (this.status_show === key) {
        this.status_show = 0
      } else {
        this.status_show = key
      }
    },
    showAddTime (key) {
      if (this.time_show === key) {
        this.time_show = 0
      } else {
        this.time_show = key
      }
    },
    devChange (e) {
      this.time_manage = e * 0.25
      this.price = Math.round(e / 60 * 1100 + this.time_manage / 60 * 800)
    }
  },
  computed: {
    tasks () {
      var task = this.$store.getters.task
      var uName = this.$store.getters.userName
      if (task.items !== null && uName !== '') {
        if (uName === 'Администратор') {
          task = task.items
        } else {
          task = task.items.filter(el => el.user === uName)
        }
      } else {
        task = []
      }
      return task
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
    },
    statuses () {
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
      return arr ? arr.status.map((el) => { el = el.name; return el }) : ['Ошибка при загрузке']
    },
    weakSum () {
      var task = this.$store.getters.task
      var sum = 0
      if (task.items) {
        task.items.map((el) => {
          var dt = new Date(el.dt_create)
          var now = new Date()
          now.getDay() !== 0 ? now.setDate(now.getDate() - (now.getDay() - 1)) : now.setDate(now.getDate() - 7)
          if (dt >= now) {
            sum += el.price
          }
        })
      }
      return sum
    }
  }
}
</script>
