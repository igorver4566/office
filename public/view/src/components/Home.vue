<template>
    <v-container>
        <v-layout row>
            <v-flex xs12>
                <h1>Hello</h1>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
export default {
  data () {
    return {}
  },
  computed: {
    user () {
      return this.$store.getters.user
    }
  }
}
</script>
