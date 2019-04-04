<template>
    <div class="big-button">
        <h1>{{ title }}</h1>
    </div>
</template>


<script>
    export default {
        name: 'BigButton',
        props: {
            title: String
        }
    }
</script>