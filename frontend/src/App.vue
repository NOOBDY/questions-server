<template>
  <div id="app">
    <div class="container flex-container">
      <div
        v-for="(question, index) in questions"
        :key="index"
        class="container m-1 question"
      >
        <question :question="question" :index="index" />
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Question from "./components/Question.vue";

export default {
  name: "App",
  components: {
    Question,
  },
  data() {
    return {
      questions: null,
    };
  },
  async created() {
    this.questions = await this.getQuestions();
  },
  methods: {
    async getQuestions() {
      const response = await axios.get("http://localhost:9090/questions");
      return response.data;
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  margin-top: 60px;
}
.question {
  width: 300px;
  border-radius: 20px;
  background-color: #1c1c2e;
  box-shadow: -1rem 0 3rem #000;
  margin: 30px;
}
.question:hover {
  transform: translate(-0.5rem, -0.5rem);
  transition: 100ms;
}
.flex-container {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}
</style>
