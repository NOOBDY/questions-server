<template>
  <div>Total Value Packages: {{ totalVuePackages }}</div>
</template>

<script>
import axios from "axios";
export default {
  name: "GetQuestions",
  data() {
    return {
      totalVuePackages: null,
    };
  },
  async created() {
    const response = await axios.get("localhost:8080/questions/0");
    this.totalVuePackages = response.data.question;
  },
};
</script>