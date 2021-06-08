<template>
  <div class="milestone">
    <div class="milestone-header" :data-progress="progressLabel" :class="[progressLabel]">
      <h3>{{ milestone.title }}</h3>
      <Markdown :value="milestone.description" />
    </div>

    <Deliverables :deliverables="milestone.deliverables || []" />
  </div>
</template>

<script lang="ts">
import { defineComponent, toRef, computed } from "vue";

export default defineComponent({
  props: {
    milestone: {
      type: Object,
      required: true,
    },
  },

  setup(props) {
    const milestone = toRef(props, "milestone");
    const breakdown = computed(() => milestone.value?.deliverables?.reduce((counts, d) => ({
        ...counts,
        [d.state]: (counts[d.state] || 0) + 1,
        total: (counts.total || 0) + 1
    }), {}) || [])
    
    const progressLabel = computed(() => {
        if (!breakdown.value.total || breakdown.value.TODO === breakdown.value.total) return "TODO"
        if (breakdown.value.SKIP === breakdown.value.total) return "SKIP"
        if (breakdown.value.DONE === breakdown.value.total) return "DONE"
        return "DOING"
    })

    return {
      milestone,
      progressLabel,
    };
  },
});
</script>

<style scoped>
.milestone {
  position: relative;
  margin: 2rem auto;
}

.milestone-header {
  border-left: 4px solid rgb(200, 200, 200);
  border-bottom-right-radius: 4px;
  border-top-right-radius: 4px;
  background: rgba(200, 200, 200, 0.1);
  margin-bottom: -2rem;
  padding: 1rem 2rem;

  margin-left: 7rem;
  text-align: left;
  position: relative;
}

.milestone-header > h3 {
    margin-top: 0;
}

.milestone-header::before,
.milestone-header::after {
  position: absolute;
  display: block;
  top: 0;
}

.milestone-header::before {
  left: -8rem;
  content: attr(data-progress);
  text-align: right;
  font-size: 0.9rem;
  font-weight: bold;
  opacity: 0.7;
  min-width: 6rem;
  top: 18px;
}

.milestone-header::after {
  box-shadow: 0 0 0 6px rgba(200, 200, 200, 1);
  left: -15px;
  background: #666;
  border-radius: 50%;
  height: 25px;
  width: 25px;
  content: "";
  top: 18px;
}
.milestone-header.DOING::after {
  box-shadow: 0 0 0 6px rgb(48, 179, 255);
}

.milestone-header.DONE::after {
  box-shadow: 0 0 0 6px rgb(36, 153, 0);
}

.milestone-header.SKIP::after {
  box-shadow: 0 0 0 6px rgb(197, 112, 0);
}

</style>