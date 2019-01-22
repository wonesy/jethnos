<template>
<div class="container choose-container">
  <progress id="tribeProgress" class="progress" 
    :class="{'is-success': numSelectedTribes===6}" value="0" max="6"></progress>
  <div class="field is-grouped is-grouped-multiline">
    <div
      class="control"
      v-for="(tribe,idx) in tribes"
      :key="idx"
      @click="inspectTribe(tribe)">
        <div class="tags has-addons">
        <span class="tag tag-name is-dark">{{tribe.name}}</span>
        <span class="tag"
          :class="{'is-text': !tribe.selected, 'is-success': tribe.selected}"
          @click="toggleSelectTribe(tribe)">
          <span class="icon">
            <i v-if="tribe.selected" class="fas fa-check"></i>
            <i v-else class="fas fa-chevron-up"></i>
          </span>
        </span>
      </div>
    </div>
  </div>
  <div v-if="errorMsg!==''" class="notification is-danger">
    {{errorMsg}}
  </div>
  <div class="inspection box">
    <p>{{inspectedPower}}</p>
  </div>
</div>
</template>

<script>
import { JethnosTribes } from './../game_rules/tribes.js'

export default {
  name: 'TribeChoose',
  data () {
    return {
      tribes: JethnosTribes,
      selectedTribes: [],
      inspectedTribe: null,
      numSelectedTribes: 0,
      errorMsg: ''
    }
  },
  watch: {
    numSelectedTribes: function (val, oldVal) {
      this.errorMsg = ''
    }
  },
  computed: {
    inspectedPower: function () {
      if (this.inspectedTribe === null) {
        return 'Click on a tribe to see their power'
      }
      return this.inspectedTribe.power
    }
  },
  methods: {
    inspectTribe: function (tribe) {
      this.inspectedTribe = tribe
    },
    toggleSelectTribe: function (tribe) {
      if (tribe.selected) {
        if (this.numSelectedTribes > 0) {
          this.numSelectedTribes -= 1
        }
        tribe.selected = false
      } else {
        if (this.numSelectedTribes === 6) {
          this.errorMsg = 'You have already selected 6 tribes'
          return
        }
        this.numSelectedTribes += 1
        tribe.selected = true
      }
      this.updateProgressBar()
    },
    updateProgressBar: function () {
      var elem = document.getElementById('tribeProgress')
      if (elem === null) {
        return
      }
      elem.value = this.numSelectedTribes
    }
  }
}
</script>

<style scoped>
.choose-container {
  /* width: 100%; */
  max-width: 60vw;
  text-align: center;
}

.control :hover {
  cursor: pointer;
}

.inspection {
  margin-bottom: 10px;
}
</style>
