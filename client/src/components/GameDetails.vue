<template>
  <div v-if="data!==null" class="card">
    <header class="card-header">
      <p class="card-header-title">{{data.name}}</p>
      <a href="#" class="card-header-icon" aria-label="more options">
        <span class="icon">
          <i class="fas fa-angle-down" aria-hidden="true"></i>
        </span>
      </a>
    </header>
    <div class="card-content">
      <div class="content">
        <p><strong>Game ID:</strong> {{data.uuid}}</p>
        <p><strong>Players in lobby:</strong> {{data.numClients}}</p>
        <p><strong>Game mode:</strong> {{modeString}}</p>
      </div>
    </div>
    <footer class="card-footer">
      <a href="#" class="card-footer-item">Join</a>
    </footer>
  </div>
  <header v-else class="card-header">
      <p class="card-header-title">No game is selected</p>
      <a href="#" class="card-header-icon" aria-label="more options">
        <span class="icon">
          <i class="fas fa-angle-down" aria-hidden="true"></i>
        </span>
      </a>
    </header>
</template>

<script>
export default {
  name: 'GameDetails',
  data () {
    return {
    }
  },
  props: {
    data: Object
  },
  computed: {
    modeString: function () {
      let mode2String = {
        0: 'Democracy',
        1: 'Dictatorship'
      }
      let str = mode2String[this.data.mode]
      if (str === undefined) {
        return 'Unknown'
      }
      return str
    }
  }
}
</script>

<style scoped>

</style>
