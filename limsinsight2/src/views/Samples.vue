<template>
    <div>

    <!-- Sample Table  -->
    <b-table hover :items="samples" :fields="fields">
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
                {key: "patients.first_name", label: "First Name", sortable: true},
                {key: "patients.last_name", label: "Last Name", sortable: true},
                {key: "patients.mrn", label: "MRN", sortable: true},
                {key: "samples.sample_id", label: "Sample ID", sortable: true},
                {key: "samples.facility", label: "Facility", sortable: true},
                {key: "samples.test_requested", label: "Test Requested", sortable: true},
                {key: "samples.se_num", label: "SE Number", sortable: true},
                {key: "samples.date_collected", label: "Date Collected", sortable: true},
                {key: "samples.date_received", label: "Date Received", sortable: true},
                {key: "samples.sample_type", label: "Sample Type", sortable: true},
                {key: "samples.material_received", label: "Material Received", sortable: true},
                {key: "samples.material_received_num", label: "Material Received Number", sortable: true},
                {key: "samples.material_received_other", label: "Material Received Other", sortable: true},
                {key: "samples.volume_of_blood_marrow", label: "Volume of Blood Marrow", sortable: true},
                {key: "samples.surgical_num", label: "Surgical Number", sortable: true},
                {key: "samples.tumor_site", label: "Tumor Site", sortable: true},
                {key: "samples.historical_diagnosis", label: "Historical Diagnosis", sortable: true},
                {key: "samples.tumor_percnt_of_total", label: "Tumor Percent of Total", sortable: true},
                {key: "samples.tumor_percnt_of_circled", label: "Tumor Percent of Circled", sortable: true},
                {key: "samples.reviewed_by", label: "Reviewed By", sortable: true},
                {key: "samples.h_e_slide_location", label: "H E Slide Location", sortable: true},
                {key: "samples.non_uhn_id", label: "NonUHN ID", sortable: true},
                {key: "samples.name_of_requestor", label: "Name of Requestor", sortable: true},
                {key: "samples.dna_concentration", label: "DNA Concentration", sortable: true},
                {key: "samples.dna_volume", label: "DNA Volume", sortable: true},
                {key: "samples.dna_location", label: "DNA Location", sortable: true},
                {key: "samples.rna_concentration", label: "RNA Concentration", sortable: true},
                {key: "samples.rna_volume", label: "RNA Volume", sortable: true},
                {key: "samples.rna_location", label: "RNA Location", sortable: true},
                {key: "samples.wbc_location", label: "WBC Location", sortable: true},
                {key: "samples.plasma_location", label: "Plasma Location", sortable: true},
                {key: "samples.cf_plasma_location", label: "CF Plasma Location", sortable: true},
                {key: "samples.pb_bm_location", label: "PB BM Location", sortable: true},
                {key: "samples.rna_lysate_location", label: "RNA Lysate Location", sortable: true},
                {key: "samples.sample_size", label: "Sample Size", sortable: true},
                {key: "samples.study_id", label: "Study ID", sortable: true},
                {key: "samples.sample_name", label: "Sample Name", sortable: true},
                {key: "samples.date_submitted", label: "Date Submitted", sortable: true},
                {key: "samples.container_type", label: "Container Type", sortable: true},
                {key: "samples.container_name", label: "Container Name", sortable: true},
                {key: "samples.container_id", label: "Container ID", sortable: true},
                {key: "samples.container_well", label: "Container Well", sortable: true},
                {key: "samples.copath_num", label: "CoPath Number", sortable: true},
                {key: "samples.other_identifier", label: "Other Identifier", sortable: true},
                {key: "samples.has_sample_files", label: "Has Sample Files", sortable: true},
                {key: "samples.dna_sample_barcode", label: "DNA Sample Barcode", sortable: true},
                {key: "samples.dna_extraction_date", label: "DNA Extraction Date", sortable: true},
                {key: "samples.dna_quality", label: "DNA Quality", sortable: true},
                {key: "samples.ffpe_qc_date", label: "FFPE QC Date", sortable: true},
                {key: "samples.delta_ct_value", label: "Delta CT Value", sortable: true},
                {key: "samples.comments", label: "Comments", sortable: true},
                {key: "samples.rnase_p_date", label: "RNASE P Date", sortable: true},
                {key: "samples.dna_quality_by_rnase_p", label: "DNA Quality by RNASE P", sortable: true},
                {key: "samples.rna_quality", label: "RNA Quality", sortable: true},
                {key: "samples.rna_extraction_date", label: "RNA Extraction Date", sortable: true},
                {key: "samples.patient_id", label: "Patient ID", sortable: true},
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

    // Calls getSamples method when component is created
    mounted() {
        this.getSamples();
    }
}
</script>