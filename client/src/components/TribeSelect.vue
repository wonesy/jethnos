<template>
  <div class="tribe-container">
    <progress :class="{'is-success': numSelected===6}" id="selectedProgress" class="progress" value="0" max="6"></progress>
    <div class="tabs is-boxed">
      <ul>
        <li
          :class="{'is-active': tribe===activeTribe, glow: inSelectedList(tribe)}"
          v-for="tribe in allTribes"
          :key="tribe.name"
          v-on:click="activeTribe=tribe">
          <a>
            <span>{{tribe.name}}</span>
          </a>
        </li>
      </ul>
    </div>
    <div v-if="activeTribe!==null" class="tribe-description box">
      <p>{{activeTribe.power}}</p>
      <br>
      <a v-if="isSelected"
        v-on:click="removeTribe"
        class="button is-danger">Remove Tribe</a>
      <a v-else
        v-on:click="addTribe"
        class="button is-primary"
        :disabled="maxSelected">Select Tribe</a>
    </div>
  </div>
</template>

<script>
import {JethnosTribes} from '../jethnos_rules/tribes.js'

export default {
  name: 'TribeSelect',
  props: {
    submit: {
      type: Boolean
    }
  },
  data () {
    return {
      allTribes: JethnosTribes,
      activeTribe: null,
      selectedTribes: [],
      errorMsg: null,
      warningMsg: null
    }
  },
  computed: {
    isSelected: function () {
      var index = this.selectedTribes.indexOf(this.activeTribe)
      if (index > -1) {
        return true
      }
      return false
    },
    numSelected: function () {
      var num = this.selectedTribes.length
      return num
    },
    maxSelected: function () {
      return this.numSelected >= 6
    }
  },
  watch: {
    submit: {
      immediate: false,
      handler: function (val, oldVal) {
        const MAX_TRIBES = 6
        if (val === false) {
          return
        }

        // user has requested to submit a new game with these settings
        if (this.numSelected > MAX_TRIBES) {
          this.errorMsg = 'Too many tribes are chosen. Required = ' + MAX_TRIBES
          this.$emit('is-valid', false, null, null)
          return
        }

        if (this.numSelected < MAX_TRIBES) {
          this.errorMsg = 'Too few tribes are chosen. Required = ' + MAX_TRIBES
          this.$emit('is-valid', false, null, null)
          return
        }

        this.$emit('is-valid', true, this.name, this.selectedTribes)
      }
    }
  },
  methods: {
    removeTribe: function () {
      var index = this.selectedTribes.indexOf(this.activeTribe)
      if (index > -1) {
        this.selectedTribes.splice(index, 1)
        this.updateProgressBar()
      }
    },
    addTribe: function () {
      var index = this.selectedTribes.indexOf(this.activeTribe)
      if (index === -1 && this.numSelected < 6) {
        this.selectedTribes.push(this.activeTribe)
        this.updateProgressBar()
      }
    },
    inSelectedList: function (tribe) {
      var index = this.selectedTribes.indexOf(tribe)
      if (index > -1) {
        return true
      }
      return false
    },
    updateProgressBar: function () {
      var elem = document.getElementById('selectedProgress')
      elem.value = this.numSelected.toString()
    }
  }
}
</script>

<style scoped>
.tribe-container {
  min-height: 100px;
}

.glow a {
  color:goldenrod !important;
}

.tribe-description {
  margin-bottom: 20px;
}
</style>
