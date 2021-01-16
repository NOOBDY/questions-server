<template>
  <div class="text">
    <p>Question {{ index + 1 }}: {{ question.question }}</p>
    <ul class="list-group">
      <li class="btn option">A: {{ question.option_a }}</li>
      <li class="btn option">B: {{ question.option_b }}</li>
      <li class="btn option">C: {{ question.option_c }}</li>
      <li class="btn option">D: {{ question.option_d }}</li>
    </ul>
    <button class="btn" id="answer-btn" v-on:click="show = !show">
      Show Answer
    </button>
    <p v-if="show">{{ question.answer }}</p>
  </div>
</template>

<script>
export default {
  props: ["question", "index"],
  data() {
    return {
      show: false,
    };
  },
};
</script>

<style>
.option {
  text-align: left;
  margin: 5px;
  background-color: #1c1c2e;
}
.text {
  padding: 20px;
  margin: 20px;
  color: #b2b5cc;
}
.btn {
  color: #b2b5cc;
  background-color: transparent;
  border-color: turquoise;
  border-radius: 20px;
  transition: 30ms;
}
.btn:hover {
  color: white;
  background-color: turquoise;
}
#answer-btn {
  margin: 10px;
}
</style>