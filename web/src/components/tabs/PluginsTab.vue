<template>
  <v-window-item :value="tab">
    <feed>
      <v-table>
        <thead>
          <tr>
            <th class="text-left">
              Name
            </th>
            <th class="text-left">
              Author
            </th>
            <th class="text-left">
              Version
            </th>
            <th class="text-left">
              API Version
            </th>
            <th class="text-left">
              Source
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="item in desserts"
            :key="item.name"
          >
            <td>{{ item.name }}</td>
            <td>{{ item.author }}</td>
            <td>{{ item.version }}</td>
            <td>{{ item.api_version }}</td>
            <td>
              <a
                :href="item.source"
                target="_blank"
              >
                github.com/Smuggr/Schedule-Keepr
              </a>
            </td>
          </tr>
        </tbody>
      </v-table>
    </feed>
  </v-window-item>
</template>


<script>
export default {
  name: 'PluginsTab',
  props: {
    tab: {
      type: Object,
      required: true,
    }
  },
  data() {
    return {
      desserts: [
        {
          name: 'Schedule-Keepr',
          author: 'Smuggr',
          version: '1.0.0',
          api_version: '1.0.0',
          source: 'https://github.com/Smuggr/Schedule-Keepr',
        },
      ],
    }
  }
}
</script>
