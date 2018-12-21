<template>
    <v-container fluid fill-height>
        <v-layout align-center justify-center>
        <v-flex xs12 sm8>
            <v-stepper v-model="e1">
                <v-stepper-header>
                <v-stepper-step editable :complete="e1 > 1" step="1">Общее</v-stepper-step>
                <v-divider></v-divider>
                <v-stepper-step editable :complete="e1 > 2" step="2">Заказчик / Исполнитель</v-stepper-step>
                <v-divider></v-divider>
                <v-stepper-step editable :complete="e1 > 3" step="3">Дополнительно</v-stepper-step>
                <v-divider></v-divider>
                <v-stepper-step step="4">ТЗ</v-stepper-step>
                </v-stepper-header>

                <v-stepper-items>
                <v-stepper-content step="1">
                    <v-container fluid fill-height grid-list-sm>
                    <v-layout align-center justify-space-between wrap>
                        <v-flex xs12 md5>
                            <v-text-field 
                                name="name" 
                                :rules="[v => !!v || 'Название задачи обязательное поле']"
                                label="Название задачи" 
                                type="text" 
                                v-model="name" 
                                required
                            ></v-text-field>

                        </v-flex>
                        <v-flex xs12 md5>
                            <v-select
                                :items="managers"
                                label="Менеджер"
                                v-model="manager" 
                            ></v-select>

                        </v-flex>
                        <v-flex xs12>
                            <v-textarea 
                                name="access" 
                                label="Доступы" 
                                auto-grow
                                v-model="access" 
                                rows="3"
                            ></v-textarea>

                        </v-flex>
                    </v-layout>
                    </v-container>
                    <v-btn
                        color="primary"
                        @click="e1 = 2"
                        >
                        Далее
                    </v-btn>

                    
                </v-stepper-content>

                <v-stepper-content step="2">
                    <v-container fluid fill-height grid-list-sm>
                    <v-layout align-center justify-space-between wrap>
                        <v-flex xs12 md5>
                            <v-select
                                :items="owners"
                                label="Заказчик"
                                v-model="owner" 
                            ></v-select>

                        </v-flex>
                        <v-flex xs12 md5>
                            <v-select
                                :items="developers"
                                label="Исполнитель"
                                v-model="developer" 
                            ></v-select>

                        </v-flex>
                    </v-layout>
                    </v-container>

                    <v-btn
                    color="primary"
                    @click="e1 = 3"
                    >
                    Далее
                    </v-btn>

                    
                </v-stepper-content>

                <v-stepper-content step="3">
                    <v-container fluid fill-height grid-list-sm>
                    <v-layout align-center justify-space-between wrap>
                        <v-flex xs12>
                            <v-select
                                v-model="tags"
                                :items="tags_items"
                                attach
                                chips
                                label="Теги"
                                multiple
                            ></v-select>

                        </v-flex>
                        <v-flex xs12>
                            <v-switch
                            :label="'Создать канал в Slack'"
                            v-model="make_slack"
                            ></v-switch>

                        </v-flex>
                    </v-layout>
                    </v-container>

                    <v-btn
                    color="primary"
                    @click="e1 = 4"
                    >
                    Далее
                    </v-btn>

                    
                </v-stepper-content>

                <v-stepper-content step="4">
                    <v-container fluid fill-height grid-list-sm>
                    <v-layout align-center justify-space-between wrap>
                        <v-flex xs12>
                            <v-flex xs12>
                                <v-textarea 
                                    box
                                    background-color="#fafafa"
                                    name="message" 
                                    label="Тех. задание" 
                                    auto-grow
                                    v-model="message" 
                                    rows="6"
                                ></v-textarea>

                            </v-flex>

                        </v-flex>
                    </v-layout>
                    </v-container>

                    <v-btn
                    color="primary"
                    @click="onSubmit()"
                    >
                    Создать
                    </v-btn>

                    
                </v-stepper-content>
                </v-stepper-items>
            </v-stepper>
        </v-flex>
        </v-layout>
      </v-container>
</template>

<script>
  import tag from './tags.json'
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
    data () {
      return {
        e1: 0,
        name: '',
        manager: '',
        access: '',
        owner: '',
        developer: '',
        tags: '',
        make_slack: false,
        message: ''
      }
    },
    methods: {
      onSubmit () {
        var arr = this.$store.getters.form
        const task = {
          id: null,
          name: this.name,
          access: this.access,
          manager_id: parseInt(getIdFromArray(arr.manager, this.manager)),
          owner_id: parseInt(getIdFromArray(arr.owner, this.owner)),
          developer_id: parseInt(getIdFromArray(arr.developer, this.developer)),
          tags: this.tags.join(', '),
          make_slack: this.make_slack ? 1 : 0,
          message: this.message,
          name_slack: '',
          status_id: 0
        }
        this.$store.dispatch('makeTask', task)
            .then(() => {})
            .catch(() => {})
      }
    },
    computed: {
      tags_items () {
        return tag.tags
      },
      managers () {
        var arr = this.$store.getters.form
        return arr ? arr.manager.map((el) => { el = el.name; return el }) : ['Ошибка при загрузке']
      },
      owners () {
        var arr = this.$store.getters.form
        return arr ? arr.owner.map((el) => { el = el.name; return el }) : ['Ошибка при загрузке']
      },
      developers () {
        var arr = this.$store.getters.form
        return arr ? arr.developer.map((el) => { el = el.name; return el }) : ['Ошибка при загрузке']
      }
    },
    created () {
      this.$store.dispatch('getFormFields')
      .then((r) => {})
      .catch(() => {})
    }
  }
</script>
