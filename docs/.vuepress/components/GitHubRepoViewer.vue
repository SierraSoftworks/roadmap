<template>
  <div>
    <Viewer :roadmap="parsed" v-if="parsed" />
    <div class="no-content" v-else>
      <h2>No Roadmap Found</h2>
      <p>Enter the name of a repository you wish to view the road map for.</p>

      <label>
        <input type="text" v-model="repo" placeholder="SierraSoftworks/roadmap">
      </label>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref, watch } from "vue";
import yaml from "js-yaml";

export default defineComponent({
  setup() {
    const repo = ref(null as string|null);
    const roadmap = ref(null as string|null);
    const error = ref(null as Error|null);

    watch(
      () => window.location.hash,
      (hash) => {
        console.log(`Hash changed: ${hash}`)
        repo.value = hash && hash.substring(1);
        error.value = null;

        fetch(
          `https://api.github.com/repos/${repo.value}/contents/roadmap.yml`,
          {
            headers: {
              Accept: "application/vnd.github.v3+json",
            },
          }
        )
          .then((res) => res.json())
          .then((res) => {
            if (res.type === "file") {
              roadmap.value = atob(res.content);
            } else {
              roadmap.value = null;
              error.value = new Error(JSON.stringify(res));
            }
          })
          .catch((err) => {
            error.value = err;
          });
      },
      {
        immediate: true
      }
    );

    const parsed = computed(() => {
      try {
        return (
          roadmap.value &&
          yaml.load(roadmap.value, {
            schema: yaml.JSON_SCHEMA,
          })
        );
      } catch (err) {
        error.value = err;
        return null;
      }
    });

    return {
        repo,
        roadmap,
        parsed,
        error
    }
  },
});
</script>

<style scoped>
.no-content {
  text-align: center;
}

input[type="text"] {
  padding: 0.5rem;
  min-width: 40vw;
  text-align: center;
  font-size: 1.4rem;
}
</style>