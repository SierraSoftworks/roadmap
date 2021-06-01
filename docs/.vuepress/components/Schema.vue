<template>
    <pre><code>{{ schema }}</code></pre>
</template>

<script lang="ts">
import {defineComponent, ref, computed} from 'vue'

export default defineComponent({
    props: {
        url: {
            type: String,
            required: true
        }
    },

    setup(props) {
        const schema = ref(null)
        const error = ref(null)

        fetch(props.url)
            .then(res => res.json())
            .then(res => (schema.value = res))
            .catch(err => {
                error.value = err
            })

        return {
            schema,
            error
        }
    }
})
</script>

<style scoped>
   
</style>