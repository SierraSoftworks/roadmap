<template>
  <div class="viewer">
    <h1 class="viewer-title">{{ roadmap.title }}</h1>

    <div class="viewer-authors">
      <div class="viewer-authors__title">Authored by</div>
      <div
        class="viewer-author"
        v-for="author in roadmap.authors"
        :key="author.name"
      >
        <h5>{{ author.name }}</h5>
        <p>{{ author.contact }}</p>
      </div>
    </div>

    <Markdown class="viewer-description" :value="roadmap.description" />

    <div class="viewer-timeline">
      <h2>Important Dates</h2>
      <Timeline :timeline="roadmap.timeline" />
    </div>

    <div class="viewer-milestones">
      <h2>Key Milestones</h2>

      <Milestone
        :milestone="milestone"
        v-for="milestone in roadmap.milestones"
        :key="milestone.title"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, toRef, computed } from "vue";

export default defineComponent({
  props: {
    roadmap: {
      type: Object,
      required: true,
    },
  },

  setup(props) {
    const roadmap = toRef(props, "roadmap");
    return {
      roadmap,
    };
  },
});
</script>

<style scoped>
.viewer-authors {
  display: flex;
  flex-direction: row;

  background: rgba(200, 200, 200, 0.15);
  padding: 1rem;
  border-radius: 6px;
}

.viewer-authors__title {
  vertical-align: middle;
  font-size: 0.7rem;
  line-height: 1.8rem;
  opacity: 0.6;
}

.viewer-author {
  display: block;
  margin-left: 1rem;
}

.viewer-author:not(:first-child) {
  margin-left: 1rem;
  padding-left: 1rem;
  border-left: 1px solid rgba(200, 200, 200, 0.5);
}

.viewer-author > h5 {
  font-size: 0.8rem;
  margin: 0;
}

.viewer-author > p {
  font-size: 0.6rem;
  margin: 0;
}

.viewer-milestones {
  position: relative;
}

.viewer-milestones::before {
  background: linear-gradient(
    0deg,
    rgba(200, 200, 200, 0) 0%,
    #c8c8c8 1%,
    #c8c8c8 99%,
    rgba(200, 200, 200, 0) 100%
  );
  content: "";
  display: block;
  width: 4px;
  position: absolute;
  left: 7rem;
  top: 3rem;
  bottom: -1rem;
}
</style>