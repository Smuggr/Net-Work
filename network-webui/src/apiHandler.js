import axios from 'axios';
import { useAuthStore } from './stores/auth';
import { useAppStore } from './stores/app';
import config from './config.json';

const api = axios.create({
	baseURL: config.baseURL,
	headers: {
		'Content-Type': 'application/json'
	}
});

const authenticateUser = async (login, password) => {
	const data = {
		login: login,
		password: password
	};

	const options = {
		method: 'POST',
		url: '/api/v1/user/authenticate',
		headers: {
			'content-type': 'application/json',
		},
		data: data
	};

	const appStore = useAppStore();
	const authStore = useAuthStore();

	try {
		appStore.setIsLoading(true);

		const response = await api.request(options);
		console.log('authentication successful ', response.data);
		
		const jwtToken = response.data.token;
		console.log('token ', jwtToken);

		authStore.setJWTToken(jwtToken);
		appStore.setIsLoggedIn(true);
		
		setTimeout(() => {
			appStore.setIsLoading(false);
			appStore.setIsLoginDialogToggled(false);
		}, 1000);
		
		return true;
	} catch (error) {
		console.error('authentication failed:', error);

		appStore.setIsLoading(false);
		return false;
	}
};

export { authenticateUser };
