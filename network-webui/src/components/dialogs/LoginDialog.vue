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
                prepend-icon="mdi-account-key"
                @focus="clearErrors" />

              <br/>
              
              <v-text-field
                v-model="password"
                :rules="passwordRules"
                label="Password"
                prepend-icon="mdi-lock"
                type="password"
                @focus="clearErrors" />
            </v-col>
          </v-container>

          <v-card-actions>
            <v-spacer />

            <v-btn
              text="Log In"
              :disabled="hasErrors"
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
import { ref, reactive, computed } from "vue";
import { authenticateUser } from "@/apiHandler";
import { useAppStore } from "@/stores/app";

export default {
  name: "LoginDialog",
  setup() {
    const login = ref('');
    const password = ref('');

    const loginRules = reactive([
      v => !!v || 'Login is required',
      v => (v && v.length >= 8 && v.length <= 16 && !v.includes(' ')) || 'Login must be between 8 and 16 characters',
    ]);

    const passwordRules = reactive([
      v => !!v || 'Password is required',
      v => (v && v.length >= 8 && v.length <= 32) || 'Password must be between 8 and 32 characters',
      v => /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]+$/.test(v) || 'Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character',
    ]);

    const clearErrors = () => {
      loginRules.splice(2);
      passwordRules.splice(3);
    };

    const cancelForm = () => {
      login.value = '';
      password.value = '';
    };

    const hasErrors = computed(() => {
      return loginRules.some(rule => !rule(login.value)) || passwordRules.some(rule => !rule(password.value));
    });

    const submitForm = async () => {
      const appStore = useAppStore();

      const loginFailed = () => {
        appStore.setIsLoading(true);

        setTimeout(() => {
          console.log('login failed');
          const invalidLoginRule = v => false || 'Invalid login or password';

          loginRules.unshift(invalidLoginRule);
          passwordRules.unshift(invalidLoginRule);
          
          appStore.setIsLoading(false);
        }, 2000);
      };

      authenticateUser(login.value, password.value)
        .then(response => {
          console.log(response);

          if (response === true) {
            console.log('login successful');
          } else {
            loginFailed();
          }
        })
        .catch(error => {
          console.log(error);
          loginFailed();
        });
    };

    const onSubmit = () => {
      console.log('submitting form');
      if (!hasErrors.value) {
        submitForm();
      }
    };

    return {
      login,
      password,
      loginRules,
      passwordRules,
      cancelForm,
      onSubmit,
      clearErrors,
      hasErrors
    };
  },
};
</script>
