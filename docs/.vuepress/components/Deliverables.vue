<template>
<div class="deliverables">
  <div class="deliverable" v-for="deliverable in deliverables" :key="deliverable.title" :class="[deliverable.state || 'TODO']" :data-state="deliverable?.state || 'TODO'">
    <h5 class="deliverable__title">
      {{ deliverable.title }}

      <Badge :text="deliverable.requirement" :type="getRequirementStyle(deliverable?.requirement)" v-if="deliverable?.requirement"/>
    </h5>
    <Markdown :value="deliverable.description" />

    <a :href="deliverable.reference" v-if="deliverable.reference">More info &rarr;</a>
  </div></div>
</template>

<script lang="ts">
import { defineComponent, ref, toRef, computed } from "vue";

export default defineComponent({
  props: {
    deliverables: {
      type: Array,
      required: true,
    },
  },

  setup(props) {
    const deliverables = toRef(props, 'deliverables');

    const getRequirementStyle = ref(req => ({
      'MUST': 'danger',
      'SHOULD': 'warning',
      'MAY': 'tip'
    }[req]))

    return {
      deliverables,
      getRequirementStyle,
    };
  },
});
</script>

<style scoped>
.deliverables {
  border-left: 4px solid rgb(200, 200, 200);
  border-bottom-right-radius: 4px;
  border-top-right-radius: 4px;
  background: rgba(200, 200, 200, 0.1);
  margin: 2rem auto;
  padding: 1rem 2rem;
  position: relative;

  text-align: center;
  margin-left: 7rem;
}

.deliverable {
  text-align: left;
  position: relative;
  padding-bottom: 1rem;
  margin-bottom: 1rem;
}

.deliverable::before,
.deliverable::after {
  position: absolute;
  display: block;
  top: 0;
}

.deliverable::before {
  left: -10rem;
  content: attr(data-state);
  text-align: right;
  font-size: 0.9rem;
  font-weight: bold;
  opacity: 0.7;
  min-width: 6rem;
}

.deliverable::after {
  box-shadow: 0 0 0 4px rgba(200, 200, 200, 1);
  left: -2.5rem;
  background: #444;
  border-radius: 50%;
  height: 11px;
  width: 11px;
  content: "";
  top: 5px;
}

.deliverable__title {
  margin-bottom: 0.6rem;
}

.deliverable.DOING::after {
  box-shadow: 0 0 0 4px rgb(48, 179, 255);
}

.deliverable.DONE::after {
  box-shadow: 0 0 0 4px rgb(36, 153, 0);
}

.deliverable.SKIP::after {
  box-shadow: 0 0 0 4px rgb(197, 112, 0);
}
</style>