import { reactive, toRefs } from 'vue';
import { defineStore } from 'pinia';
import { Tabs, DashboardTabs } from '@/navigationHandler';

export const useAppStore = defineStore({
	id: 'app',
	state: () => ({
		isDrawerToggled: true,
		isLoginDialogToggled: true,
		isLoggedIn: false,
		isLoading: false,
		currentTab: Tabs.ABOUT,
		currentDashboardTab: DashboardTabs.DEVICES,
		currentUser: reactive({
			login: 'administrator',
            username: 'Administrator',
        }),
	}),
	actions: {
		setIsDrawerToggled(value) {
			this.isDrawerToggled = value;
		},
		setIsLoginDialogToggled(value) {
			this.isLoginDialogToggled = value;
		},
		setIsLoggedIn(value) {
			this.isLoggedIn = value;
		},
		setIsLoading(value) {
			this.isLoading = value;
		},
		setCurrentTab(tab) {
			this.currentTab = tab;
		},
		setCurrentDashboardTab(tab) {
			this.currentDashboardTab = tab;
		},
	},
})