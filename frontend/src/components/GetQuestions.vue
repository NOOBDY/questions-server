<template>
  <div class="container">
    <div
      v-for="(question, index) in questions"
      :key="index"
      class="container m-1 border border-1"
    >
      <p>Question {{ index + 1 }}: {{ question.question }}</p>
      <ul class="list-group-flush">
        <li class="list-group-item">A: {{ question.option_a }}</li>
        <li class="list-group-item">B: {{ question.option_b }}</li>
        <li class="list-group-item">C: {{ question.option_c }}</li>
        <li class="list-group-item">D: {{ question.option_d }}</li>
      </ul>
      <button
        class="btn btn-primary"
        v-on:click="question.show = !question.show"
      >
        Show Answer
      </button>
      <p>{{ questions.show }}</p>
      <p v-if="question.show">{{ question.answer }}</p>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      questions: null,
    };
  },
  created() {
    axios
      .get("http://localhost:9090/questions")
      .then((response) => (this.questions = response.data))
      .catch((error) => {
        this.errorMessage = error.message;
        console.error("Error: ", error);
      });
  },
};
</script>