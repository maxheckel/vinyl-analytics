<template>
  <div class="home">
    <Listen/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import Listen from "../components/Listen";

export default {
  name: 'home',
  components: {
    Listen,
    HelloWorld
  }
}
</script>
