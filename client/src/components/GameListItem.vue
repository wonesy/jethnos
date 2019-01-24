<template>
  <div class="message is-dark" v-on="inputListeners">
    <div class="message-body" :class="{'is-selected': selected}">
      <div class="level">
        <div class="level-left">
          <div>
            <p><strong>{{game.name}}</strong></p>
            <p>Players: {{game.numClients}}</p>
          </div>
        </div>
        <div class="level-right">
          <div class="field is-grouped is-grouped-multiline">
            <div class="control">
              <div class="tags has-addons">
                <span class="tag is-dark">mode</span>
                <span class="tag show-mode"
                  :class="{
                    'is-success': game.params.mode===0,
                    'is-warning': game.params.mode===1
                    }">
                  {{modeString}}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'GameListItem',
  data () {
    return {}
  },
  props: {
    game: Object,
    selected: Boolean
  },
  computed: {
    inputListeners: function () {
      var vm = this
      // `Object.assign` merges objects together to form a new object
      return Object.assign({},
        // We add all the listeners from the parent
        this.$listeners,
        // Then we can add custom listeners or override the
        // behavior of some listeners.
        {
          // This ensures that the component works with v-model
          click: function (event) {
            vm.$emit('click', event.target.value)
          }
        }
      )
    },
    modeString: function () {
      let mode2String = {
        0: 'Democracy',
        1: 'Dictatorship'
      }
      let str = mode2String[this.game.params.mode]
      if (str === undefined) {
        return 'Unknown'
      }
      return str
    }
  }
}
</script>

<style scoped>
.message {
  margin-bottom: 0.25rem;
  border-left-width: 150px;
}

.message :hover {
  cursor: pointer;
}

.is-selected {
  color: goldenrod !important;
  border-color: goldenrod !important;
}

.show-mode {
  min-width: 90px;
}
</style>
