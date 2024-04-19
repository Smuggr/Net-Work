<template>
  <v-window-item :value="tab">
    <feed>
      <v-data-table
        :headers="headers"
        :items="desserts"
        :sort-by="[{ key: 'username', order: 'asc' }]"
      >
        <template v-slot:top>
          <v-toolbar flat>
            <v-toolbar-title>Manage Devices</v-toolbar-title>
            <v-divider class="mx-4" inset vertical></v-divider>
            <v-spacer></v-spacer>
            <v-dialog v-model="dialog" max-width="500px">
              <template v-slot:activator="{ props }">
                <v-btn class="mb-2" color="primary" dark v-bind="props">
                  New Device
                </v-btn>
              </template>
              <v-card>
                <v-card-title>
                  <span class="text-h5">{{ formTitle }}</span>
                </v-card-title>

                <v-card-text>
                  <v-container>
                    <v-row>
                      <v-col cols="12">
                        <v-text-field
                          v-model="editedItem.username"
                          label="Username"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12">
                        <v-text-field
                          v-model="editedItem.client_id"
                          label="Client ID"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12">
                        <v-text-field
                          v-model="editedItem.password"
                          label="Password Hash"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12">
                        <v-text-field
                          v-model="editedItem.plugin"
                          label="Plugin"
                        ></v-text-field>
                      </v-col>
                    </v-row>
                  </v-container>
                </v-card-text>

                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn color="blue-darken-1" variant="text" @click="close">
                    Cancel
                  </v-btn>
                  <v-btn color="blue-darken-1" variant="text" @click="save">
                    Save
                  </v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
            <v-dialog v-model="dialogDelete" max-width="500px">
              <v-card>
                <v-card-title class="text-h5"
                  >Are you sure you want to delete this device?</v-card-title
                >
                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn color="blue-darken-1" variant="text" @click="closeDelete"
                    >Cancel</v-btn
                  >
                  <v-btn
                    color="blue-darken-1"
                    variant="text"
                    @click="deleteItemConfirm"
                    >OK</v-btn
                  >
                  <v-spacer></v-spacer>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </v-toolbar>
        </template>
        <template v-slot:item.actions="{ item }">
          <v-icon class="me-2" size="small" @click="openInNew(item)">
            mdi-open-in-new
          </v-icon>
          <v-icon class="me-2" size="small" @click="editItem(item)">
            mdi-pencil
          </v-icon>
          <v-icon size="small" @click="deleteItem(item)"> mdi-delete </v-icon>
        </template>
        <template v-slot:no-data>
          <v-btn color="primary" @click="initialize"> Reset </v-btn>
        </template>
      </v-data-table>

      <v-dialog v-model="dialogOpenInNew" max-width="500px">
        <v-card>
          <v-card-title class="text-h5">{{ editedItem.username }}</v-card-title>

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

            <v-spacer></v-spacer>
            <v-btn @click="turnOff">Relay Off</v-btn>
            <v-btn @click="turnOn">Relay On</v-btn>
          </v-card-text>

          <v-card-actions>
            <v-btn color="blue darken-1" variant="text" location="top" @click="closeOpenInNew">Close</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </feed>
  </v-window-item>
</template>


<script>
export default {
  name: 'DevicesTab',
  props: {
    tab: {
      type: Object,
      required: true,
    }
  },

  data: () => ({
    dialog: false,
    dialogDelete: false,
    dialogOpenInNew: false,
    headers: [
      {
        title: 'Username',
        align: 'start',
        sortable: true,
        key: 'username',
      },
      { title: 'Client ID', key: 'client_id', sortable: true },
      { title: 'Password Hash', key: 'password', sortable: true },
      { title: 'Plugin', key: 'plugin', sortable: true },
      { title: 'Actions', key: 'actions', sortable: false },
    ],
    desserts: [],
    editedIndex: -1,
    editedItem: {
      username: '',
      client_id: '',
      password: '',
      plugin: '',
    },
    defaultItem: {
      username: '',
      client_id: '',
      password: '',
      plugin: '',
    },
  }),

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'New Device' : 'Edit Device'
    },
  },

  watch: {
    dialog(val) {
      val || this.close()
    },
    dialogDelete(val) {
      val || this.closeDelete()
    },
  },

  created() {
    this.initialize()
  },

  methods: {
    initialize() {
      this.desserts = [
        {
          username: 'Schedule-Keepr',
          client_id: 'schedule-keepr1',
          password: '$2a$10$XzAWBr0YIniwUSebTUaB7uDi96MEuWIKweNNv93APmhpnKmA.pKvm',
          plugin: 'Smuggr/Schedule-Keepr',
        },
      ]
    },

    openInNew(item) {
      this.editedIndex = this.desserts.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialogOpenInNew = true;
    },

    closeOpenInNew() {
      this.dialogOpenInNew = false;
    },

    editItem(item) {
      this.editedIndex = this.desserts.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
    },

    deleteItem(item) {
      this.editedIndex = this.desserts.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialogDelete = true
    },

    deleteItemConfirm() {
      this.desserts.splice(this.editedIndex, 1)
      this.closeDelete()
    },

    close() {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    closeDelete() {
      this.dialogDelete = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    save() {
      if (this.editedIndex > -1) {
        Object.assign(this.desserts[this.editedIndex], this.editedItem)
      } else {
        this.desserts.push(this.editedItem)
      }
      this.close()
    },
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