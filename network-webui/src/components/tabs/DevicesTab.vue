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
import { ref, onMounted } from 'vue';
import { getDevices, removeDevices, registerDevice, toggleDevice, getRtcTime } from '@/apiHandler';

export default {
  name: 'DevicesTab',
  props: {
    tab: {
      type: Object,
      required: true,
    }
  },
  setup() {
    const items = ref([]);
    const selected = ref([]);
    const rtcTime = ref(''); // Define rtcTime as a ref object

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
        // After removal, refresh the devices list
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
        await addNewDevice(); // You can pass any necessary parameters here
        // After addition, refresh the devices list
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
        rtcTime.value = time; // Assign the value to rtcTime.value
      } catch (error) {
        console.error('Error fetching RTC time:', error);
      }
    };

    onMounted(() => {
      //loadDevices();
      setInterval(updateRtcTime, 1000);
    });

    return {
      items,
      selected,
      rtcTime,
      removeDevices,
      refreshDevices,
      addDevice,
      turnOff,
      turnOn
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

.rtc-container {
  display: flex;
  align-items: center; /* Center vertically */
}

.rtc-icon {
  margin-right: 8px; /* Adjust margin on the right */
}
</style>