<template>
  <div>
    <h1>Questions: {{ q.question }}</h1>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      q: null,
    };
  },
  created() {
    axios
      .get("http://localhost:9090/questions/0")
      .then((response) => (this.q = response.data))
      .catch((error) => {
        this.errorMessage = error.message;
        console.error("Error: ", error);
      });
  },
};
</script>