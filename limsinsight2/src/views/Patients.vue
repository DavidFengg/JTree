<template>
    <div>

    <Alert :message="message" @done="clearMessage"/>

    <!-- Patient Table  -->
    <b-table hover bordered responsive :items="patients" :fields="fields">
        <template v-slot:cell(action)="data">
            <b-button size="sm" class="mx-1" v-on:click="showModal(data.item)">Edit</b-button>
            <b-button size="sm" class="mx-1" v-on:click="deletePatient(data.item)"> Delete </b-button>
        </template>
    </b-table>

    <!-- Add Patient -->
    <b-table-simple hover responsive> 
        <b-tbody>
            <b-tr>
                <!-- Loops through fields objects and filters for keys needed for input-->
                <b-td v-for="(field,i) in fields" v-bind:key="i" colspan="2">
                    <div v-if="showInputTag(field)" class="input-field">
                        <label>{{ field.label }}</label>

                        <input v-if="field.type == 'date'" placeholder="" v-model="input[field.key]" type="datetime-local">
                        <input v-else placeholder="" v-model="input[field.key]" type="text">
                    </div>
                </b-td>
            </b-tr>
        </b-tbody>
    </b-table-simple>

    <!-- Add button -->
    <b-button class="button" size="m" v-on:click="createPatient()"> Add </b-button>


    <!-- Edit Modal -->
    <b-modal id="edit" title="Edit Patient Data">
        <b-form-group v-for="(field,i) in fields" v-bind:key="i">
            <!-- Input tag doesn't include 'Action' -->
            <div v-if="showInputTag(field)">
                <label> {{ field.label }}</label>
                <b-form-input
                    v-model="edit[field.key]"
                    placeholder= "">
                </b-form-input>
            </div>
        </b-form-group>

        <!-- Confirmation button -->
        <b-button class="btn btn-primary" v-on:click="updatePatient()">Confirm</b-button>
    </b-modal>

    </div>
</template>

<script>
import Alert from "../components/Alert";
import shared from "../shared";
import APIService from '../services/APIService';

export default {
    components: {
        Alert
    },

    data() {
        return {
            fields: [
                {key: "patients.mrn", label: "MRN", sortable: true},
                {key: "patients.se_num", label: "SE Number", sortable: true},
                {key: "patients.first_name", label: "First Name", sortable: true},
                {key: "patients.last_name", label: "Last Name", sortable: true},
                {key: "patients.initials", label: "Initials", sortable: true},
                {key: "patients.gender", label: "Gender", sortable: true},
                {key: "patients.dob", label: "Date of Birth", type: "date", sortable: true},
                {key: "patients.on_hcn", label: "On HCN", sortable: true},
                {key: "patients.clinical_history", label: "Clinical History", sortable: true},
                {key: "patients.patient_type", label: "Patient Type", sortable: true},
                {key: "patients.patient_id", label: "Patient ID", sortable: true},
                {key: "patients.date_received", label: "Date Received", type: "date", sortable: true},
                {key: "patients.referring_physician", label: "Referring Physician", sortable: true},
                {key: "patients.date_reported", label: "Date Reported", type: "date", sortable: true},
                {key: "patients.surgical_date", label: "Surgical Date", type: "date", sortable: true},
                "Action"
            ],
            patients: [],
            edit: {},
            input: {
                "patients.mrn": "",
                "patients.se_num": "",
                "patients.first_name": "",
                "patients.last_name": "",
                "patients.initials": "",
                "patients.gender": "",
                "patients.dob": "",
                "patients.on_hcn": "",
                "patients.clinical_history": "",
                "patients.patient_type": "",
                "patients.date_received": "",
                "patients.referring_physician": "",
                "patients.date_reported": "",
                "patients.surgical_date": "", 
            },
            // Error handling
            message: "",
        };
    },

    methods: {

        // function initalizes edit information and shows the edit modal
        showModal(patient) {
            // update placeholder information by copying values
            this.edit = Object.assign({}, patient);
            
            this.$bvModal.show('edit');
        },

        getPatients() {
            APIService.getPatients().then(data => {
                this.patients = data;
                this.$store.dispatch("addUniquePatients", this.patients);
            }).catch(err => console.error(err));
        },

        createPatient() {
            // creates a new object with corrected data types 
            let modify = Shared.convert(this.input, this.fields);

            APIService.createPatient(modify).then(res => {
                this.getPatients();
            });
        },

        updatePatient() {
            APIService.updatePatient(this.edit).then(res => {
                this.getPatients();
            })
        },

        deletePatient(data) {
            let id = data["patients.patient_id"];

            APIService.deletePatient(id).then(res => {
                // call updateMessage func if status code is 405
                if (res == 405) {
                    this.updateMessage(id);
                }
                else {
                    this.getPatients();
                }
            });
        },
        
        // function returns true if the field is NOT action or the field key is NOT 'patient.patient_id'
        showInputTag(field) {
            return field != "Action" && field.key != "patients.patient_id";
        },

        // Updates the message to be sent to the alert component
        updateMessage(id) {
            this.message = "Patient with ID: " + id + " cannot be deleted";
        },

        // Clears the message once alert has finished
        clearMessage() {
            this.message = "";
        }
    },

    // Calls getPatients method when component is created
    mounted() {
        this.getPatients();
    }
}
</script>