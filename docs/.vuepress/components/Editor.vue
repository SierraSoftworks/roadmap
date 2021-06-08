<template>
  <div class="editor-container">
    <div class="editor-control">
      <v-ace-editor
        v-model:value="roadmap"
        placeholder="Enter your Roadmap definition here to begin visualizing it."
        lang="yaml"
        theme="tomorrow_night_bright"
        style="height: 90vh"
      >
      </v-ace-editor>
    </div>
    <div class="editor-preview">
      <Viewer v-if="!error" :roadmap="parsed" />
      <div class="editor-error" v-if="error">
        <h2>&#9888; Error</h2>
        <p>{{ error.message }}</p>
        <pre><code>{{ error.stack }}</code></pre>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from "vue";
import { VAceEditor } from "vue3-ace-editor";
import "ace-builds/src-noconflict/mode-yaml";
import "ace-builds/src-noconflict/theme-tomorrow_night_bright";
import yaml from "js-yaml";

const defaultRoadmap = `
title: Example Road Map
description: |
    This is an example of what a road map might look like. It can include **Markdown** if you
    wish.

authors:
  - name: Benjamin Pannell
    contact: contact@sierrasoftworks.com

timeline:
  - date: 2021-04-21
    title: Project Start
    description: This is when we will start working on the project, get the team ready!

milestones:
  - title: Build the Team
    description: We don't yet have anyone, that's not gonna work...
    deliverables:
      - title: Team Lead
        state: DONE
        requirement: MUST
        description: This person needs to know enough about this domain to be able to run with the project.

      - title: Senior Engineer No.1
        state: DOING
        requirement: SHOULD

      - title: Junior Engineer No.1
        state: TODO
        requirement: SHOULD

      - title: Barista
        state: SKIP
        requirement: MAY

  - title: Finish the Project
    description: We don't need other milestones, do we?
    deliverables:
      - title: MVP
        description: Who needs a polished product? Let's just ship the MVP and call it done.
      - title: Marketing
      - title: VC Funding
      - title: Yacht
        reference: https://lmgtfy.app/?q=yacht&t=i
`;

export default defineComponent({
  props: {},

  components: {
    VAceEditor,
  },

  setup(props) {
    const roadmap = ref(defaultRoadmap);
    const error = ref(null);
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
      roadmap,
      parsed,
      error,
    };
  },
});
</script>

<style scoped>
.editor-container {
  position: relative;
  height: calc(100vh - var(--navbar-height));
  overflow: hidden;
  margin: 0 -6rem;
  margin-bottom: -6rem;
}

.editor-control,
.editor-preview {
  position: absolute;
  width: 49%;
  max-height: 100%;
  overflow: hidden;
  overflow-y: auto;
}

.editor-preview {
  left: 51%;
  padding: 0 0.5rem;
}
</style>