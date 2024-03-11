import axios from 'axios';
import { useAuthStore } from './stores/auth';
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

	try {
		const response = await api.request(options);
		console.log('authentication successful:', response.data);
		
		const jwtToken = response.data.token;
		useAuthStore().setJWTToken(jwtToken);
		
		return true;
	} catch (error) {
		console.error('authentication failed:', error);
		return false;
	}
};

export { authenticateUser };
