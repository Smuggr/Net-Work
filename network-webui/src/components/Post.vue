<template>
  <v-card title="Card title" subtitle="Subtitle" text="Amongus Susssus">
    <v-card-actions>
      <v-btn>Click me</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: {
    post: {
      type: Object,
      required: true
    },
  },
  methods: {
    likePost() {
      // Handle like functionality
    },
    commentOnPost() {
      // Handle comment functionality
    },
    sharePost() {
      // Handle share functionality
    }
  }
};
</script>

<style scoped>
.post {
  margin-bottom: 20px;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.user-info {
  display: flex;
  align-items: center;
}

.avatar {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  margin-right: 10px;
}

.username {
  font-weight: bold;
}

.actions {
  margin-top: 10px;
}

.actions button {
  margin-right: 10px;
}
</style>
