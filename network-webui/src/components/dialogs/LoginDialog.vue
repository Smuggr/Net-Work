<template>
  <v-dialog max-width="500px">
    <template v-slot:default="{ isActive }">
      <v-card title="Log in to your account">
        <v-form @submit.prevent="onSubmit">
          <v-container>
            <v-col>
              <v-text-field
                v-model="login"
                :rules="loginRules"
                label="Login"
                prepend-icon="mdi-account-key" />

              <br/>
              
              <v-text-field
                v-model="password"
                :rules="passwordRules"
                label="Password"
                prepend-icon="mdi-lock"
                type="password" />
            </v-col>
          </v-container>

          <v-card-actions>
            <v-spacer />

            <v-btn
              text="Log In"
              type="submit" />
            <v-btn
              text="Close"
              @click="isActive.value = false; cancelForm();" />
          </v-card-actions>
        </v-form>
      </v-card>
    </template>
  </v-dialog>
</template>

<script>
import { ref } from "vue";
import { authenticateUser } from "@/apiHandler";

const login = ref('');
const password = ref('');

const loginRules = [
  v => !!v || 'Login is required',
  v => (v && v.length >= 8 && v.length <= 16 && !v.includes(' ')) || 'Login must be between 8 and 16 characters',
];

const passwordRules = [
  v => !!v || 'Password is required',
  v => (v && v.length >= 8 && v.length <= 32) || 'Password must be between 8 and 32 characters',
  v => /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]+$/.test(v) || 'Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character',
];

const cancelForm = () => {
  login.value = '';
  password.value = '';
};

const submitForm = async () => {
  authenticateUser(login.value, password.value)
    .then(response => {
      console.log(response);
    })
    .catch(error => {
      console.log(error);
    });
};

export default {
  name: "LoginDialog",
  setup() {
    return {
      login,
      password,
      loginRules,
      passwordRules,
      cancelForm,
      submitForm
    };
  },
  emits: ['form-on-submit'],
  methods: {
    onSubmit() {
      console.log('submitting form');
      this.$emit('form-on-submit');
      submitForm();
    },
  },
};
</script>
