import axios from "axios";
import {addPrefix, removePrefix} from "./helper";

const API_URL = "http://localhost:8000/Jtree/metadata/0.1.0";

export default {
    getPatients() {
        return axios.post(API_URL + "/query", {
            selected_fields: ["patients.first_name", "patients.last_name", "patients.initials", "patients.gender", "patients.mrn", "patients.dob", "patients.on_hcn", "patients.clinical_history", "patients.patient_type", "patients.se_num", "patients.patient_id", "patients.date_received", "patients.referring_physician", "patients.date_reported", "patients.surgical_date"],
            selected_tables: ["patients"],
            selected_conditions: [[]]
        }).then(res => {
            let filter = res.data;

            // remove the 'patients.' prefix for each object
            filter.forEach( (object) => {
                object = removePrefix(object, "patients");
            });

            return filter;
        }); 
    },
    
    createPatient(data) {
        data = addPrefix(data);

        return axios.post(API_URL + "/patient", data)
        .then(res => console.log(res.data))
        .catch(err => console.log(err));
    },

    updatePatient(data) {
        let id = data.patient_id;

        data = addPrefix(data, "patients");

        return axios.put(API_URL + "/patient/" + id, data);
    },

    deletePatient(id) {
        return axios.delete(API_URL + "/patient/" + id).then(res => {
            console.log(res.data);
        });
    },

    getSamples() {
        return axios.post(API_URL + "/query", {
            selected_fields: ["patients.first_name", "patients.last_name", "samples.sample_id", "patients.mrn", "samples.sample_id", "samples.facility", "samples.test_requested", "samples.se_num", "samples.date_collected", "samples.date_received", "samples.sample_type", "samples.material_received", "samples.material_received_num", "samples.material_received_other", "samples.volume_of_blood_marrow", "samples.surgical_num", "samples.tumor_site", "samples.historical_diagnosis", "samples.tumor_percnt_of_total", "samples.tumor_percnt_of_circled", "samples.reviewed_by", "samples.h_e_slide_location", "samples.non_uhn_id", "samples.name_of_requestor", "samples.dna_concentration", "samples.dna_volume", "samples.dna_location", "samples.rna_concentration", "samples.rna_volume", "samples.rna_location", "samples.wbc_location", "samples.plasma_location", "samples.cf_plasma_location", "samples.pb_bm_location", "samples.rna_lysate_location", "samples.sample_size", "samples.study_id", "samples.sample_name", "samples.date_submitted", "samples.container_type", "samples.container_name", "samples.container_id", "samples.container_well", "samples.copath_num", "samples.other_identifier", "samples.has_sample_files", "samples.dna_sample_barcode", "samples.dna_extraction_date", "samples.dna_quality", "samples.ffpe_qc_date", "samples.delta_ct_value", "samples.comments", "samples.rnase_p_date", "samples.dna_quality_by_rnase_p", "samples.rna_quality", "samples.rna_extraction_date", "samples.patient_id"],
            selected_tables: ["patients", "samples"],
            selected_conditions: [[]]
        }).then(res => {
            let filter = res.data;

            // remove the 'samples.' prefix for each object
            filter.forEach( (object) => {
                object = removePrefix(object, "samples");
            });

            return filter;
        }); 
    
    }
}