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
    const response = await axios.get("https://api.npms.io/v2/search?q=vue");
    this.totalVuePackages = response.data.total;
  },
};
</script>