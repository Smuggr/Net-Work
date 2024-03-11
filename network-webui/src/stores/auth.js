import { defineStore } from "pinia";

export const useAuthStore = defineStore({
	id: "auth",
	state: () => ({
		isAuthenticated: false,
		jwtToken: "",
	}),
	actions: {
		setJWTToken(token) {
			this.jwtToken = token;
			this.isAuthenticated = true;
		},
		clearJWTToken() {
			this.jwtToken = "";
			this.isAuthenticated = false;
		},
	},
});