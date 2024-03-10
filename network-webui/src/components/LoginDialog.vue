<script setup>
const model = defineModel()
</script>

<template>
  <v-dialog v-model="model" max-width="500">
    <template v-slot:default="{ isActive }">
      <v-card title="Log in to your account">
        <v-form @submit.prevent>
        <v-container>
          <v-col>
            <v-text-field
              v-model="username"
              :rules="userRules"
              label="Username"
              prepend-icon="mdi-account" />
            
            <v-text-field
              v-model="login"
              :rules="loginRules"
              label="Login"
              prepend-icon="mdi-account-key" />

            <v-text-field
              v-model="password"
              :rules="passwordRules"
              label="Password"
              prepend-icon="mdi-lock"
              type="password" />
          </v-col>
        </v-container>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn
            text="Log In"
            type="submit"
            @click="isActive.value = false" />
          <v-btn
            text="Close"
            @click="isActive.value = false" />
        </v-card-actions>
        </v-form>
      </v-card>
    </template>
  </v-dialog>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      login: '',
      password: '',

      usernameRules: [
        v => !!v || 'Username is required',
        v => (v && v.length >= 8 && v.length <= 32) || 'Username must be between 8 and 32 characters',
      ],
      loginRules: [
        v => !!v || 'Login is required',
        v => (v && v.length >= 8 && v.length <= 16 && !v.includes(' ')) || 'Login must be between 8 and 16 characters and cannot contain spaces',
      ],
      passwordRules: [
        v => !!v || 'Password is required',
        v => (v && v.length >= 8 && v.length <= 32) || 'Password must be between 8 and 32 characters',
        v => /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]+$/.test(v) || 'Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character',
      ],
    };
  },
};
</script>
