<template>
    <div>
    <!-- Patient Table  -->
    <b-table hover :items="patients" :fields="fields">
        <template v-slot:cell(action)="data">
            <b-button size="sm" class="mx-1" v-b-modal="'edit'" v-on:click="updatePlaceholder(data.item)">Edit</b-button>
            <b-button size="sm" class="mx-1" v-on:click="deletePatient(data.item.patient_id)"> Delete </b-button>
        </template>
    </b-table>

    <!-- Add Patient -->
    <b-table-simple hover> 
        <b-tbody>
            <b-tr>
                <b-td v-for="field in fields" colspan="2">
                    
                    <!-- Input tag doesn't include 'Action' -->
                    <div v-if="showInput(field)" class="input-field">
                        <label>{{ field.key }}</label>
                        <input placeholder="" v-model="input[field.key]" type="text">
                    </div>
                </b-td>
                
                <b-td>
                    <b-button class="button" size="sm" v-on:click="createPatient()"> Add </b-button>
                </b-td>
            </b-tr>
        </b-tbody>
    </b-table-simple>

    </div>
</template>

<script>
import APIService from '../services/APIService';

export default {

    data() {
        return {
            fields: [
                {key: "mrn", label: "MRN", sortable: true},
                {key: "first_name", sortable: true},
                {key: "last_name", sortable: true},
                {key: "initials", sortable: true},
                {key: "gender", sortable: true},
                {key: "dob", label: "Date of Birth", sortable: true},
                {key: "on_hcn", sortable: true},
                {key: "clinical_history", sortable: true},
                {key: "patient_type", sortable: true},
                {key: "patient_id", sortable: true},
                {key: "date_recieved", sortable: true},
                {key: "referring_physician", sortable: true},
                {key: "date_reported", sortable: true},
                {key: "surgical_date", sortable: true},
                "Action"
            ],
            patients: [],
            input: {
                mrn: "",
                first_name: "",
                last_name: "",
                initials: "",
                gender: "",
                dob: "",
                on_hcn: "",
                clinical_history: "",
                patient_type: "",
                date_recieved: "",
                referring_physician: "",
                date_reported: "",
                surgical_date: "", 
            }
        };
    },

    methods: {
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

        updatePlaceholder(patient) {
            // copy patient object data
            let copyData = JSON.parse(JSON.stringify(patient));
            this.edit = copyData;
        },

        deletePatient(id) {
            APIService.deletePatient(id).then(res => {
                this.getPatients();
            });
        },
        
        // function returns true if the field is not action or the field key is not 'patient_id'
        showInput(field) {
            return field != "Action" && field.key != "patient_id";
        }
    },


    mounted() {
        this.getPatients();
    }
}
</script>