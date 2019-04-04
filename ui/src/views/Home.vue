<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <BigButton title="Listen"/>
    <BigButton title="+ Album"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import BigButton from "../components/BigButton";

export default {
  name: 'home',
  components: {
    BigButton,
    HelloWorld
  }
}
</script>
