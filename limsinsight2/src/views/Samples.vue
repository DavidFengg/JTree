<template>
    <div>

    <!-- Samples Table  -->
    <b-table :items="samples" :fields="fields" hover responsive bordered>
        <template v-slot:cell(action)="data">
            <b-button size="sm" class="mx-1" v-on:click="showModal(data.item)">Edit</b-button>
            <b-button size="sm" class="mx-1" v-on:click="deleteSample(data.item)"> Delete </b-button>
        </template>
    </b-table>

    <!-- Patient ID dropdown -->
    <b-dropdown text="Patient ID">
        <b-dropdown-item v-on:click="updateCurrPatient(patient)" v-for="patient in getPatients()">
            {{ patient["patients.patient_id"]}}
        </b-dropdown-item>
    </b-dropdown>

    <!-- Selected ID text -->
    <div class="mt-3">Selected: <strong>{{ selected["patients.patient_id"] }}</strong></div>

    <!-- Add Sample -->
    <b-table-simple hover responsive class="mt-3"> 
        <b-tbody>
            <b-tr>
                <!-- Static text fields based on Patient ID selected -->
                <b-td colspan="2">
                    <div class="input-field">
                        <label>First Name</label>
                        <input :disabled="true" :placeholder="selected['patients.first_name']" type="text">
                    </div>
                </b-td>
                <b-td colspan="2">
                    <div class="input-field">
                        <label>Last Name</label>
                        <input :disabled="true" :placeholder="selected['patients.last_name']" type="text">
                    </div>
                </b-td>
                <b-td colspan="2">
                    <div class="input-field">
                        <label>MRN</label>
                        <input :disabled="true" :placeholder="selected['patients.mrn']" type="text">
                    </div>
                </b-td>
                
                <!-- Input fields -->
                <b-td v-for="(field,i) in modified()" v-bind:key="i">
                        <label>{{ field.label }}</label>

                        <!-- input type changes based on field property -->
                        <input v-if="field.type == 'date'" placeholder="" v-model="input[field.key]" type="datetime-local">
                        <input v-else-if="field.type == 'number'" placeholder="" v-model="input[field.key]" type="number">
                        <input v-else-if="field.label.includes('Has Sample Files')" placeholder="" v-model="input[field.key]" type="checkbox">
                        <input v-else placeholder="" v-model="input[field.key]" type="text">
                </b-td>

            </b-tr>
        </b-tbody>
    </b-table-simple>

    <!-- Add Button -->
    <b-button class="button" size="m" v-on:click="createSample()"> Add </b-button>

    
    <!-- Edit Modal -->
    <b-modal id="edit" title="Edit Sample Data">
        <b-form-group>
            <div v-for="(field,i) in modified()" v-bind:key="i">
                <label> {{ field.label }}</label>

                <!-- input type changes based on field property -->
                <b-form-input v-if="field.type == 'number'" v-model="edit[field.key]" placeholder= "" type="number"></b-form-input>
                <b-form-checkbox v-else-if="field.label.includes('Has Sample Files')" v-model="edit[field.key]" placeholder= ""></b-form-checkbox>
                <b-form-input v-else v-model="edit[field.key]" placeholder= "" type="text"></b-form-input>
            </div>
        </b-form-group>

        <!-- Confirmation button -->
        <b-button class="btn btn-primary" v-on:click="updateSample()">Confirm</b-button>
    </b-modal>
    
    </div>
</template>

<style scoped>
    .input-field {
        white-space: pre-wrap;
    }
</style>

<script>
import APIService from '../services/APIService';

export default {

    data() {
        return {
            fields: [
                {key: "samples.patient_id", label: "Patient ID", sortable: true},
                {key: "patients.first_name", label: "First Name", sortable: true},
                {key: "patients.last_name", label: "Last Name", sortable: true},
                {key: "patients.mrn", label: "MRN", sortable: true},
                {key: "samples.sample_id", label: "Sample ID", sortable: true},
                {key: "samples.facility", label: "Facility", sortable: true},
                {key: "samples.test_requested", label: "Test Requested", sortable: true},
                {key: "samples.se_num", label: "SE Number", sortable: true},
                {key: "samples.date_collected", label: "Date Collected", type: 'date', sortable: true},
                {key: "samples.date_received", label: "Date Received", type: 'date', sortable: true},
                {key: "samples.sample_type", label: "Sample Type", sortable: true},
                {key: "samples.material_received", label: "Material Received", sortable: true},
                {key: "samples.material_received_num", label: "Material Received Number", sortable: true},
                {key: "samples.material_received_other", label: "Material Received Other", sortable: true},
                {key: "samples.volume_of_blood_marrow", label: "Volume of Blood Marrow", type:"number", sortable: true},
                {key: "samples.surgical_num", label: "Surgical Number", sortable: true},
                {key: "samples.tumor_site", label: "Tumor Site", sortable: true},
                {key: "samples.historical_diagnosis", label: "Historical Diagnosis", sortable: true},
                {key: "samples.tumor_percnt_of_total", label: "Tumor Percent of Total", type:"number", sortable: true},
                {key: "samples.tumor_percnt_of_circled", label: "Tumor Percent of Circled", type:"number", sortable: true},
                {key: "samples.reviewed_by", label: "Reviewed By", sortable: true},
                {key: "samples.h_e_slide_location", label: "H E Slide Location", sortable: true},
                {key: "samples.non_uhn_id", label: "NonUHN ID", sortable: true},
                {key: "samples.name_of_requestor", label: "Name of Requestor", sortable: true},
                {key: "samples.dna_concentration", label: "DNA Concentration", type:"number", sortable: true},
                {key: "samples.dna_volume", label: "DNA Volume", type:"number", sortable: true},
                {key: "samples.dna_location", label: "DNA Location", sortable: true},
                {key: "samples.rna_concentration", label: "RNA Concentration", type:"number", sortable: true},
                {key: "samples.rna_volume", label: "RNA Volume", type:"number", sortable: true},
                {key: "samples.rna_location", label: "RNA Location", sortable: true},
                {key: "samples.wbc_location", label: "WBC Location", sortable: true},
                {key: "samples.plasma_location", label: "Plasma Location", sortable: true},
                {key: "samples.cf_plasma_location", label: "CF Plasma Location", sortable: true},
                {key: "samples.pb_bm_location", label: "PB BM Location", sortable: true},
                {key: "samples.rna_lysate_location", label: "RNA Lysate Location", sortable: true},
                {key: "samples.sample_size", label: "Sample Size", sortable: true},
                {key: "samples.study_id", label: "Study ID", sortable: true},
                {key: "samples.sample_name", label: "Sample Name", sortable: true},
                {key: "samples.date_submitted", label: "Date Submitted", type: 'date', sortable: true},
                {key: "samples.container_type", label: "Container Type", sortable: true},
                {key: "samples.container_name", label: "Container Name", sortable: true},
                {key: "samples.container_id", label: "Container ID", sortable: true},
                {key: "samples.container_well", label: "Container Well", sortable: true},
                {key: "samples.copath_num", label: "CoPath Number", sortable: true},
                {key: "samples.other_identifier", label: "Other Identifier", sortable: true},
                {key: "samples.has_sample_files", label: "Has Sample Files", type:"checkbox", sortable: true},
                {key: "samples.dna_sample_barcode", label: "DNA Sample Barcode", sortable: true},
                {key: "samples.dna_extraction_date", label: "DNA Extraction Date", type: 'date', sortable: true},
                {key: "samples.dna_quality", label: "DNA Quality", sortable: true},
                {key: "samples.ffpe_qc_date", label: "FFPE QC Date", type: 'date', sortable: true},
                {key: "samples.delta_ct_value", label: "Delta CT Value", type:"number", sortable: true},
                {key: "samples.comments", label: "Comments", sortable: true},
                {key: "samples.rnase_p_date", label: "RNASE P Date", type: 'date', sortable: true},
                {key: "samples.dna_quality_by_rnase_p", type:"number", label: "DNA Quality by RNASE P", sortable: true},
                {key: "samples.rna_quality", label: "RNA Quality", type:"number", sortable: true},
                {key: "samples.rna_extraction_date", label: "RNA Extraction Date", type: 'date', sortable: true},
                "Action"
            ],
            samples: [],
            edit: {},
            input: {
                "samples.facility": "",
                "samples.test_requested": "",
                "samples.se_num": "",
                "samples.date_collected": "",
                "samples.date_received": "",
                "samples.sample_type": "",
                "samples.material_received": "",
                "samples.material_received_num": "",
                "samples.material_received_other": "",
                "samples.volume_of_blood_marrow": 0,
                "samples.surgical_num": "",
                "samples.tumor_site": "",
                "samples.historical_diagnosis": "",
                "samples.tumor_percnt_of_total": 0,
                "samples.tumor_percnt_of_circled": 0,
                "samples.reviewed_by": "",
                "samples.h_e_slide_location": "",
                "samples.non_uhn_id": "",
                "samples.name_of_requestor": "",
                "samples.dna_concentration": 0,
                "samples.dna_volume": 0,
                "samples.dna_location": "",
                "samples.rna_concentration": 0,
                "samples.rna_volume": 0,
                "samples.rna_location": "",
                "samples.wbc_location": "",
                "samples.plasma_location": "",
                "samples.cf_plasma_location": "",
                "samples.pb_bm_location": "",
                "samples.rna_lysate_location": "",
                "samples.sample_size": "",
                "samples.study_id": "",
                "samples.sample_name": "",
                "samples.date_submitted": "",
                "samples.container_type": "",
                "samples.container_name": "",
                "samples.container_id": "",
                "samples.container_well": "",
                "samples.copath_num": "",
                "samples.other_identifier": "",
                "samples.has_sample_files": false,
                "samples.dna_sample_barcode": "",
                "samples.dna_extraction_date": "",
                "samples.dna_quality": "",
                "samples.ffpe_qc_date": "",
                "samples.delta_ct_value": 0,
                "samples.comments": "",
                "samples.rnase_p_date": "",
                "samples.dna_quality_by_rnase_p": 0,
                "samples.rna_quality": 0,
                "samples.rna_extraction_date": "",
                "samples.patient_id": ""
            },
            // chosen patient from dropdown
            selected: {}
        };
    },

    methods: {

        // function initalizes edit information and shows the edit modal
        showModal(sample) {
            // update placeholder information by copying values
            this.edit = Object.assign({}, sample);
            
            this.$bvModal.show('edit');
        },

        // function converts fields from the input/edit objects to the correct data types
        convert(object) {
            let modify = Object.assign({}, object);

            // find which fields require conversions
            for (let key in modify) {
                for (let i = 0; i < this.fields.length; i++) {
                    if (this.fields[i].key == key) {

                        // convert string to number
                        if (this.fields[i].type == "number") {
                            console.log(this.fields[i].key);
                            modify[key] = Number(modify[key]);
                        }
                        // convert dates to iso string
                        else if (this.fields[i].type == "date") {
                            let date = new Date(modify[key]);

                            modify[key] = date.toISOString();
                        }   
                    }
                }
            }
            return modify;
        },

        getSamples() {
            APIService.getSamples().then(data => {
                this.samples = data;
            }).catch(err => console.error(err));
        },

        createSample() {
            let modify = this.convert(this.input);

            APIService.createSample(modify).then(res => {
                this.getSamples();
            });
        },

        updateSample() {
            let modify = this.convert(this.edit);

            APIService.updateSample(modify).then(res => {
                this.getSamples();
            })
        },

        deleteSample(data) {
            let id = data["samples.sample_id"];

            APIService.deleteSample(id).then(res => {
                this.getSamples();
            });
        },
        
        // 
        showInputTag(field) {
            return field != "Action" && field.key != "samples.sample_id" && !field.key.startsWith("patients");
        },

        // returns a modified version of the fields array with only objects needed for creating samples
        modified() {
            let modified = [];
            for (let i = 0; i < this.fields.length - 1; i++) {
                let key = this.fields[i].key;

                // adds only the objects needed for adding samples
                if (!(key.startsWith("patients") || key == "samples.sample_id" || key == "Action")) {
                    modified.push(this.fields[i]);
                }
            }

            return modified;
        },

        // get the patients saved from store
        getPatients() {
            return this.$store.getters.patients;
        },

        // update the current patient that the user selected
        updateCurrPatient(patient) {
            this.selected = patient;

            // update the input's patient id
            this.input["samples.patient_id"] = patient["patients.patient_id"];
        }
    },

    // Calls getSamples method when component is created
    mounted() {
        this.getSamples();
    }
}
</script>