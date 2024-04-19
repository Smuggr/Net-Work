<template>
  <v-dialog max-width="500px" persistent>
    <template v-slot:default="{ isActive }">
      <v-card title="Log in to your account">
        <v-progress-linear
          :active="isLoading"
          :indeterminate="true"
          location="top"
          color="deep-purple-accent-4"
          absolute
          bottom
        ></v-progress-linear>

        <v-form @submit.prevent="onSubmit">
          <v-container>
            <v-col>
              <v-text-field
                v-model="login"
                :error-messages="loginErrorMessages"
                label="Login"
                prepend-icon="mdi-account-key"
                clearable
              />

              <br />

              <v-text-field
                v-model="password"
                :error-messages="passwordErrorMessages"
                label="Password"
                prepend-icon="mdi-lock"
                type="password"
                clearable
              />
            </v-col>
          </v-container>

          <v-card-actions>
            <v-spacer />

            <v-btn text="Log In" type="submit" />
            <v-btn text="Close" @click="isActive.value = false;" />
          </v-card-actions>
        </v-form>
      </v-card>
    </template>
  </v-dialog>
</template>

<script>
import { ref, reactive, computed, watch } from "vue";
import { authenticateUser } from "@/apiHandler";
import { useAppStore } from "@/stores/app";

export default {
  name: "LoginDialog",
  setup() {
    const appStore = useAppStore();

    const login = ref("");
    const password = ref("");

    const isLoading = ref(false);

    const loginErrorMessages = ref("");
    const passwordErrorMessages = ref("");

    const invalidCredentialsMessage = 'Invalid login or password';

    const clearErrors = () => {
      loginErrorMessages.value = "";
      passwordErrorMessages.value = "";
    };

    const clearForm = () => {
      login.value = "";
      password.value = "";
      clearErrors();
    };

    const hasErrors = computed(() => !!loginErrorMessages.value || !!passwordErrorMessages.value);

    const submitForm = async () => {
      isLoading.value = true;

      const loginFailed = () => {
        setTimeout(() => {
          console.log("login failed");

          loginErrorMessages.value = invalidCredentialsMessage;
          passwordErrorMessages.value = invalidCredentialsMessage;

          isLoading.value = false;
        }, 2000);
      };

      try {
        const response = await authenticateUser(login.value, password.value);
        console.log(response);

        if (response === true) {
          console.log("login successful");
          isLoading.value = false;
        } else {
          loginFailed();
        }
      } catch (error) {
        console.log(error);
        loginFailed();
      }
    };

    const onSubmit = () => {
      console.log("submitting form");
      submitForm();
      clearErrors();
    };

    watch([login, password], () => {
      hasErrors.value;
    });

    watch(
      () => appStore.$state.isLoginDialogToggled,
      (newValue, oldValue) => {
        clearErrors();
        clearForm();
      }
    );

    return {
      login,
      password,
      isLoading,
      loginErrorMessages,
      passwordErrorMessages,
      hasErrors,
      onSubmit,
    };
  },
};
</script>
