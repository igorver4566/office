<template>
    <v-container fluid fill-height>
        <v-layout align-center justify-center>
        <v-flex xs8>
            <v-stepper v-model="e1">
                <v-stepper-header>
                <v-stepper-step editable :complete="e1 > 1" step="1">Общее</v-stepper-step>
                <v-divider></v-divider>
                <v-stepper-step editable :complete="e1 > 2" step="2">Заказчик</v-stepper-step>
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
                        <v-flex xs5 md3>
                            <v-text-field 
                                name="time_manage" 
                                label="Время менеджера" 
                                type="text" 
                                v-model="time_manage"
                            ></v-text-field>

                        </v-flex>
                        <v-flex xs5 md3>
                            <v-text-field 
                                name="time_dev" 
                                label="Время Исполнителя" 
                                type="text" 
                                v-model="time_dev"
                            ></v-text-field>

                        </v-flex>
                        <v-flex xs12 md5>
                            <v-text-field 
                                name="price" 
                                label="Цена" 
                                type="text" 
                                v-model="price"
                            ></v-text-field>

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
  function getIdFromArray (el) {
    return (e) => {
      if (e.name === el) {
        return e.id
      }
    }
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
        time_manage: '',
        time_dev: '',
        price: '',
        tags: '',
        make_slack: false,
        message: '',
        formFields: null
      }
    },
    methods: {
      onSubmit () {
        const task = {
          id: null,
          name: this.name,
          access: this.access,
          manager_id: parseInt(this.formFields.manager.map(getIdFromArray(this.manager)).toString()),
          time_dev: parseInt(this.time_dev),
          time_manage: parseInt(this.time_manage),
          owner_id: parseInt(this.formFields.owner.map(getIdFromArray(this.owner)).toString()),
          developer_id: parseInt(this.formFields.developer.map(getIdFromArray(this.developer)).toString()),
          price: parseInt(this.price),
          tags: this.tags.join(','),
          make_slack: this.make_slack ? 1 : 0,
          message: this.message
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
        return this.formFields ? this.formFields.manager.map((el) => { el = el.name; return el }) : ['Ошибка при загрузке']
      },
      owners () {
        return this.formFields ? this.formFields.owner.map((el) => { el = el.name; return el }) : ['Ошибка при загрузке']
      },
      developers () {
        return this.formFields ? this.formFields.developer.map((el) => { el = el.name; return el }) : ['Ошибка при загрузке']
      }
    },
    created () {
      this.$store.dispatch('getFormFields')
      .then((r) => {
        this.formFields = this.$store.getters.form
        console.log()
      })
      .catch(() => {})
    }
  }
</script>
