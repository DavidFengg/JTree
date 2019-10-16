<template>
    <div>

    <!-- Sample Table  -->
    <b-table hover :items="patients" :fields="fields">
        <template v-slot:cell(action)="data">
            <b-button size="sm" class="mx-1" v-on:click="showModal(data.item)">Edit</b-button>
            <b-button size="sm" class="mx-1" v-on:click="deleteSample(data.item.sample_id)"> Delete </b-button>
        </template>
    </b-table>

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
            samples: [],
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

        getSamples() {
            APIService.getSamples().then(data => {
                this.samples = data;
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
        this.getSamples();
    }
}
</script>