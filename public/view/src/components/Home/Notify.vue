<template>
  <v-card>
    <v-card-title primary-title><h3>Оповещения</h3></v-card-title>
    <v-list two-line class="scroll-y notify">
    <template v-for="(item, index) in items">
      <v-divider
        v-if="index > 0"
        :key="index"
      ></v-divider>

      <v-list-tile
        :key="item.title"
      >
      <v-list-tile-content>
          <v-list-tile-title v-html="item.title"></v-list-tile-title>
          <v-list-tile-sub-title v-html="item.subtitle"></v-list-tile-sub-title>
      </v-list-tile-content>
      </v-list-tile>
    </template>
    </v-list>
  </v-card>
</template>

<script>
export default {
  data () {
    return {
      items: [
        {
          title: 'Оповещение 1',
          subtitle: 'Какой-то текст'
        },
        {
          title: 'Оповещение 2',
          subtitle: 'Какой-то текст'
        },
        {
          title: 'Оповещение 3',
          subtitle: 'Какой-то текст'
        },
        {
          title: 'Оповещение 22',
          subtitle: 'Какой-то текст'
        },
        {
          title: 'Оповещение 32',
          subtitle: 'Какой-то текст'
        }
      ]
    }
  }
}
</script>

<style>
  .notify {
    max-height: 300px;
  }
</style>
