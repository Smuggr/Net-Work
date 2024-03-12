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

const deviceApi = axios.create({
	baseURL: 'http://192.168.1.30:5000',
	headers: {
		'Content-Type': 'application/json'
	}
});


const authenticateUser = async (login, password) => {
	const appStore = useAppStore();
	const authStore = useAuthStore();

	const data = {
		login: login,
		password: password
	};

	const options = {
		method: 'POST',
		url: '/user/authenticate',
		headers: {
			'content-type': 'application/json',
		},
		data: data
	};

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

const registerUser = async (login, username, password, permissionLevel) => {
	const appStore = useAppStore();
	const authStore = useAuthStore();

	const data = {
		login: login,
		username: username,
		password: password,
		permissionLevel: permissionLevel,
	};

	const options = {
		method: 'POST',
		url: '/user/register',
		headers: {
			'content-type': 'application/json',
			'Authorization': `Bearer ${authStore.jwtToken}`
		},
		data: data
	};

	try {
		appStore.setIsLoading(true);

		const response = await api.request(options);
		console.log('registering user successful ', response.data);
		
		return true;
	} catch (error) {
		console.error('registering user failed:', error);

		return false;
	}
};

const registerDevice = async (clientId, username, password, plugin) => {
	const authStore = useAuthStore();

	const data = {
		client_id: clientId,
		usernam: username,
		password: password,
		plugin: plugin,
	};

	const options = {
		method: 'POST',
		url: '/device/register',
		headers: {
			'content-type': 'application/json',
			'Authorization': `Bearer ${authStore.jwtToken}`
		},
		data: data
	};

	try {
		const response = await api.request(options);
		console.log('registering device successful ', response.data);

		return true;
	} catch (error) {
		console.error('registering device failed:', error);

		return false;
	}
};

const getDevices = async () => {
	const authStore = useAuthStore();

	const options = {
		method: 'GET',
		url: '/devices/all',
		headers: {
			'content-type': 'application/json',
			'Authorization': `Bearer ${authStore.jwtToken}`
		},
	};

	try {
		const response = await api.request(options);
		console.log('getting all devices successful ', response.data);

		const devices = response.data.devices;
		// const message = response.data.message;
		console.log('devices ', devices);

		return { devices };
	} catch (error) {
		console.error('getting all devices failed:', error);

		return;
	}
};

const removeDevice = async (deviceId) => {

}

const removeDevices = async (deviceId) => {

}


const toggleDevice = async (boolValue) => {
	const url = (boolValue ? '/on' : '/off');
	const options = {
		method: 'POST',
		url: url,
		headers: {
			'content-type': 'application/json',
		}
	};
	try {
		const response = await deviceApi.request(options);
		console.log('Toggling device successful:', response.data);
		return true;
	} catch (error) {
		console.error('Toggling device failed:', error);
		return false;
	}
};

const getRtcTime = async () => {
	const url = '/time';
	const options = {
		method: 'GET',
		url: url,
		headers: {
			'content-type': 'application/json',
		}
	};
	try {
		const response = await deviceApi.request(options);
		const currentTime = response.data.current_time;
		console.log('Getting RTC time successful:', currentTime);

		return currentTime;
	} catch (error) {
		console.error('Getting RTC time failed:', error);

		return null;
	}
};

export {
	authenticateUser,
	registerUser,
	registerDevice,
	getDevices,
	removeDevice,
	removeDevices,
	toggleDevice,
	getRtcTime
};
