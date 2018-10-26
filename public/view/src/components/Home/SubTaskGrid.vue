<template>
  <v-card>
    <v-card-title primary-title><h3>Задачи</h3></v-card-title>
    <v-expansion-panel
      v-model="open"
      expand
    >
      <v-expansion-panel-content
        v-for="item in items"
        :key="item.id"
      >
        <div slot="header" class="sub_panel-title">{{ item.name }}</div>
        <v-list two-line>
          <template v-for="(subItem, index) in item.items">
            <v-divider
              v-if="index > 0"
              :key="index"
            ></v-divider>
            <v-list-tile
                :key="subItem.id"
                class="grey lighten-3"
              >
                <v-list-tile-content>
                  <v-list-tile-title class="sub-title">
                    {{ subItem.name }}
                  </v-list-tile-title>
                  <v-list-tile-sub-title class="sub-title_sub">
                    <b>Цена: </b>{{ subItem.price }} руб.&nbsp;&nbsp;
                    <b>Время: </b>{{ subItem.time_dev + subItem.time_manage }} мин.&nbsp;&nbsp;
                    <b>Исполнитель: </b>{{ subItem.user }}&nbsp;&nbsp;
                    <b>Статус: </b>{{ subItem.status }}
                  </v-list-tile-sub-title>
                </v-list-tile-content>
              </v-list-tile>
          </template>
        </v-list>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-card>
</template>

<script>
export default {
  data () {
    return {
      open: [true]
    }
  },
  computed: {
    items () {
      var tasks = this.$store.getters.tasks
      var subs = this.$store.getters.subTasks
      return tasks.map((val, key, arr) => {
        var id = val.id
        val.items = []
        subs.forEach(element => {
          if (element.task_id === id) {
            val.items.push(element)
          }
        })
        return val
      })
    }
  }
}
</script>

<style>
  .sub_panel-title {
    font-size: 18px;
  }
  .sub-title {
    font-size: 15px;
  }
  .sub-title_sub {
    font-size: 13px;
  }
</style>

