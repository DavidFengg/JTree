import axios from "axios";

const API_URL = "http://localhost:8000/Jtree/metadata/0.1.0";

export default {
    // Patient http requests
    getPatients() {
        return axios.post(API_URL + "/query", {
            selected_fields: ["patients.first_name", "patients.last_name", "patients.initials", "patients.gender", "patients.mrn", "patients.dob", "patients.on_hcn", "patients.clinical_history", "patients.patient_type", "patients.se_num", "patients.patient_id", "patients.date_received", "patients.referring_physician", "patients.date_reported", "patients.surgical_date"],
            selected_tables: ["patients"],
            selected_conditions: [[]]
        }).then(res => {
            return res.data;
        }); 
    },
    
    createPatient(data) {
        return axios.post(API_URL + "/patient", data)
        .then(res => console.log(res.data))
        .catch(err => console.log(err));
    },

    updatePatient(data) {
        let id = data["patients.patient_id"];

        return axios.put(API_URL + "/patient/" + id, data);
    },

    deletePatient(id) {
        return axios.delete(API_URL + "/patient/" + id).then(res => {
            console.log(res.data);
        });
    },

    // Sample http requests
    getSamples() {
        return axios.post(API_URL + "/query", {
            selected_fields: ["patients.first_name", "patients.last_name", "samples.sample_id", "patients.mrn", "samples.sample_id", "samples.facility", "samples.test_requested", "samples.se_num", "samples.date_collected", "samples.date_received", "samples.sample_type", "samples.material_received", "samples.material_received_num", "samples.material_received_other", "samples.volume_of_blood_marrow", "samples.surgical_num", "samples.tumor_site", "samples.historical_diagnosis", "samples.tumor_percnt_of_total", "samples.tumor_percnt_of_circled", "samples.reviewed_by", "samples.h_e_slide_location", "samples.non_uhn_id", "samples.name_of_requestor", "samples.dna_concentration", "samples.dna_volume", "samples.dna_location", "samples.rna_concentration", "samples.rna_volume", "samples.rna_location", "samples.wbc_location", "samples.plasma_location", "samples.cf_plasma_location", "samples.pb_bm_location", "samples.rna_lysate_location", "samples.sample_size", "samples.study_id", "samples.sample_name", "samples.date_submitted", "samples.container_type", "samples.container_name", "samples.container_id", "samples.container_well", "samples.copath_num", "samples.other_identifier", "samples.has_sample_files", "samples.dna_sample_barcode", "samples.dna_extraction_date", "samples.dna_quality", "samples.ffpe_qc_date", "samples.delta_ct_value", "samples.comments", "samples.rnase_p_date", "samples.dna_quality_by_rnase_p", "samples.rna_quality", "samples.rna_extraction_date", "samples.patient_id"],
            selected_tables: ["patients", "samples"],
            selected_conditions: [[]]
        }).then(res => {
            return res.data;
        }); 
    },

    createSample(data) {
        console.log(data);
        return axios.post(API_URL + "/sample", data)
        .then(res => console.log(res.data))
        .catch(err => console.log(err));
    },

    updateSample(data) {
        console.log(data);
        let id = data["samples.sample_id"];

        return axios.put(API_URL + "/sample/" + id, data);
    },

    deleteSample(id) {
        return axios.delete(API_URL + "/sample/" + id).then(res => {
            console.log(res.data);
        });
    },

    // Experiment http requests
    getExperiments() {
        return axios.post(API_URL + "/query", {
            selected_fields: ["experiments.experiment_id", "experiments.study_id", "experiments.panel_assay_screened", "experiments.test_date", "experiments.chip_cartridge_barcode", "experiments.complete_date", "experiments.pcr", "experiments.sample_id", "experiments.project_name", "experiments.priority", "experiments.opened_date", "experiments.project_id", "experiments.has_project_files", "experiments.procedure_order_datetime"],
            selected_tables: ["experiments"],
            selected_conditions: [[]]
        }).then(res => {
            return res.data;
        });     
    },

    createExperiment(data) {
        console.log(data);
        return axios.post(API_URL + "/experiment", data)
        .then(res => console.log(res.data))
        .catch(err => console.log(err));
    },

    updateExperiment(data) {
        console.log(data);
        let id = data["experiments.experiment_id"];

        return axios.put(API_URL + "/experiment/" + id, data);
    },

    deleteExperiment(id) {
        return axios.delete(API_URL + "/experiment/" + id);
    },
    
    // Result http requests
    getResults() {
        return axios.post(API_URL + "/query", {
            selected_fields: ["results.failed_regions", "results.mean_depth_of_coveage", "results.mlpa_pcr", "results.mutation", "results.overall_hotspots_threshold", "results.overall_quality_threshold", "results.results_id", "results.uid", "results.verification_pcr", "results.experiment_id"],
            selected_tables: ["results"],
            selected_conditions: [[]]            
        }).then(res => {
            return res.data;
        });
    },

    createResult(data) {
        console.log(data);
        return axios.post(API_URL + "/result", data)
        .then(res => console.log(res.data))
        .catch(err => console.log(err));
    },

    updateResult(data) {
        console.log(data);
        let id = data["results.results_id"];

        return axios.put(API_URL + "/result/" + id, data);
    },

    deleteResult(id) {
        return axios.delete(API_URL + "/result/" + id);
    }

}
