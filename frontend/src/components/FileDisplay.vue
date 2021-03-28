<template>
    <div class="filedisplay">
        {{ name }}
    </div>
</template>

<script>
export default {
    name: "FileDisplay",
    setup() {
        
    },
    props: ['name']
}
</script>

<style scoped>

.filedisplay {
    margin-top: 5px;
    padding: 10px;
    background-color: white;
}

</style>