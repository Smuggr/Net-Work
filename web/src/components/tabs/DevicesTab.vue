<template>
  <v-window-item :value="tab">
    <feed>
      <v-card>
        <v-card-title>
          <span class="headline">Schedule Keepr</span>
        </v-card-title>
        <v-card-text>
          <div class="rtc-container">
            <v-icon class="rtc-icon">mdi-clock</v-icon>
            <span class="rtc-time">{{ rtcTime }}</span>
          </div>
          
          <br/>

          <div class="gpio-container">
            <v-icon class="gpio-icon">{{ gpioIconClass }}</v-icon>
            <span class="gpio-status">{{ gpioStatus }}</span>
          </div>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn @click="turnOff">Relay Off</v-btn>
          <v-btn @click="turnOn">Relay On</v-btn>
        </v-card-actions>
      </v-card>
    </feed>
  </v-window-item>
</template>


<script>
import { ref, onMounted, computed } from 'vue';
import { getDevices, removeDevices, registerDevice, toggleDevice, getRtcTime, getGpioStatus } from '@/apiHandler';

export default {
  name: 'DevicesTab',
  props: {
    tab: {
      type: Object,
      required: true,
    }
  },
  setup() {
    const rtcTime = ref('');
    const gpioStatus = ref('');

    const loadDevices = async () => {
      try {
        const { devices } = await getDevices();
        if (Array.isArray(devices)) {
          items.value = devices;
        } else {
          console.error('Devices is not an array:', devices);
        }
      } catch (error) {
        console.error('Error fetching devices:', error);
      }
    };

    const removeDevices = async () => {
      try {
        await removeSelectedDevices(selected.value);
        await loadDevices();
      } catch (error) {
        console.error('Error removing devices:', error);
      }
    };

    const refreshDevices = async () => {
      await loadDevices();
    };

    const addDevice = async () => {
      try {
        await addNewDevice();
        await loadDevices();
      } catch (error) {
        console.error('Error adding device:', error);
      }
    };

    const turnOff = () => {
      toggleDevice(false);
    };

    const turnOn = () => {
      toggleDevice(true);
    };

    const updateRtcTime = async () => {
      try {
        const time = await getRtcTime();
        rtcTime.value = time;
      } catch (error) {
        console.error('Error fetching RTC time:', error);
      }
    };

    const updateGpioStatus = async () => {
      try {
        const status = await getGpioStatus();
        gpioStatus.value = status;
      } catch (error) {
        console.error('Error fetching GPIO status:', error);
      }
    };

    onMounted(() => {
      setInterval(updateRtcTime, 1000);
      setInterval(updateGpioStatus, 1000);
    });

    const gpioIconClass = computed(() => {
      return gpioStatus.value === '1' ? 'mdi-led-on' : 'mdi-led-off';
    });

    return {
      rtcTime,
      gpioStatus,
      turnOff,
      turnOn,
      gpioIconClass,
    };
  },
}
</script>

<style scoped>
.post {
  max-width: 100%;
  max-height: 380px;
  margin-bottom: 64px;
}

.gpio-container,
.rtc-container {
  display: flex;
  align-items: center;
}

.gpio-icon,
.rtc-icon {
  margin-right: 8px;
}

</style>
