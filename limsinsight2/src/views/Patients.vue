<template>
    <div>

    <!-- Patient Table  -->
    <b-table hover :items="patients" :fields="fields">
        <template v-slot:cell(action)="data">
            <b-button size="sm" class="mx-1" v-on:click="showModal(data.item)">Edit</b-button>
            <b-button size="sm" class="mx-1" v-on:click="deletePatient(data.item.patient_id)"> Delete </b-button>
        </template>
    </b-table>

    <!-- Add Patient -->
    <b-table-simple hover> 
        <b-tbody>
            <b-tr>
                <b-td v-for="field in fields" colspan="2">
                    <!-- Input tag doesn't include 'Action' -->
                    <div v-if="showInputTag(field)" class="input-field">
                        <label>{{ field.key }}</label>
                        <input placeholder="" v-model="input[field.key]" type="text">
                    </div>
                </b-td>
                
                <!-- Add button -->
                <b-td>
                    <b-button class="button" size="sm" v-on:click="createPatient()"> Add </b-button>
                </b-td>
            </b-tr>
        </b-tbody>
    </b-table-simple>

    <!-- Edit Modal -->
    <b-modal id="edit" title="Edit Patient Data">
        <b-form-group v-for="field in fields">
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
import APIService from '../services/APIService';

export default {

    data() {
        return {
            fields: [
                {key: "mrn", label: "MRN", sortable: true},
                {key: "se_num", label: "SE Number", sortable: true},
                {key: "first_name", label: "First Name", sortable: true},
                {key: "last_name", label: "Last Name", sortable: true},
                {key: "initials", label: "Initials", sortable: true},
                {key: "gender", label: "Gender", sortable: true},
                {key: "dob", label: "Date of Birth", sortable: true},
                {key: "on_hcn", label: "On HCN", sortable: true},
                {key: "clinical_history", label: "Clinical History", sortable: true},
                {key: "patient_type", label: "Patient Type", sortable: true},
                {key: "patient_id", label: "Patient ID", sortable: true},
                {key: "date_received", label: "Date Received", sortable: true},
                {key: "referring_physician", label: "Referring Physician", sortable: true},
                {key: "date_reported", label: "Date Reported", sortable: true},
                {key: "surgical_date", label: "Surgical Date", sortable: true},
                "Action"
            ],
            patients: [],
            edit: {},
            input: {
                mrn: "",
                se_num: "",
                first_name: "",
                last_name: "",
                initials: "",
                gender: "",
                dob: "",
                on_hcn: "",
                clinical_history: "",
                patient_type: "",
                date_received: "",
                referring_physician: "",
                date_reported: "",
                surgical_date: "", 
            }
        };
    },

    methods: {
        // function initalizes edit information and shows the edit modal
        showModal(patient) {
            // update placeholder information
            this.edit = patient;
            
            this.$bvModal.show('edit');
        },

        getPatients() {
            APIService.getPatients().then(data => {
                this.patients = data;
            }).catch(err => console.error(err));
        },

        createPatient() {
            APIService.createPatient(this.input).then(res => {
                this.getPatients();
            });
        },

        updatePatient() {
            APIService.updatePatient(this.edit).then(res => {
                this.getPatients();
            })
        },

        deletePatient(id) {
            APIService.deletePatient(id).then(res => {
                this.getPatients();
            });
        },
        
        // function returns true if the field is NOT action or the field key is NOT 'patient_id'
        showInputTag(field) {
            return field != "Action" && field.key != "patient_id";
        }
    },

    // Calls getPatients method when component is created
    mounted() {
        this.getPatients();
    }
}
</script>