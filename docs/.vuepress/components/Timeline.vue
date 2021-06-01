<template>
  <div class="timeline">
    <div class="timeline-item" v-for="item in sorted" :key="item.date" :data-date="item.date">
      <h4 class="timeline-item__title">{{ item.title }}</h4>
      <div class="timeline-item__description">{{ item.description }}</div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, toRef, computed } from "vue";

export default defineComponent({
  props: {
    timeline: {
      type: Array,
      required: true,
    },
  },

  setup(props) {
    const timeline = toRef(props, 'timeline');
    const sorted = computed(() => {
      const arr = timeline.value.splice(0)
      arr.sort((a: any,b: any) => a.date.localeCompare(b.date))
      return arr
    })
    return {
      timeline,
      sorted,
    };
  },
});
</script>

<style scoped>
.timeline {
  border-left: 4px solid rgb(128,128,128);
  border-bottom-right-radius: 4px;
  border-top-right-radius: 4px;
  background: rgba(200, 200, 200, 0.1);
  margin: 2rem auto;
  padding: 1rem 2rem;
  position: relative;

  text-align: center;
  margin-left: 7rem;
}

.timeline-item {
  text-align: left;
  position: relative;
  padding-bottom: 1rem;
  margin-bottom: 1rem;
}

.timeline-item::before,
.timeline-item::after {
  position: absolute;
  display: block;
  top: 0;
}

.timeline-item::before {
  left: -10rem;
  content: attr(data-date);
  text-align: right;
  font-size: 0.9rem;
  font-weight: bold;
  opacity: 0.7;
  min-width: 6rem;
}

.timeline-item::after {
  box-shadow: 0 0 0 4px rgba(128, 128, 128, 1);
  left: -2.4rem;
  background: #444;
  border-radius: 50%;
  height: 11px;
  width: 11px;
  content: "";
  top: 5px;
}

.timeline-item__title {
  margin-bottom: 0.6rem;
}
</style>