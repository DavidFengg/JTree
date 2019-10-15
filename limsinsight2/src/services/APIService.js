import axios from "axios";
const API_URL = "http://localhost:8000/Jtree/metadata/0.1.0";

export default {
    getPatients() {
        return axios.post(API_URL + "/query", {
            selected_fields: ["patients.first_name", "patients.last_name", "patients.initials", "patients.gender", "patients.mrn", "patients.dob", "patients.on_hcn", "patients.clinical_history", "patients.patient_type", "patients.se_num", "patients.patient_id", "patients.date_received", "patients.referring_physician", "patients.date_reported", "patients.surgical_date"],
            selected_tables: ["patients"],
            selected_conditions: [[]]
        }).then(res => {
            let filter = res.data;

            // filter the 'patients.' prefix in each key
            filter.forEach( (object) => {
                for (name in object) {
                    if(name.startsWith('patients.')){
                        let replaced_key = name.replace('patients.', '');
                        object[replaced_key] = object[name];
                        delete object[name];
                    }
                }
            })

            return filter;
        }); 
    },
    
    createPatient(data) {
        return axios.post(API_URL + "/patient", {
            "patients.clinical_history": "aiCMRAjWwhTHctcuAxhxKQFDa",
            "patients.date_received": "1963-02-04T00:00:00Z",
            "patients.date_reported": "1970-02-17T00:00:00Z",
            "patients.dob": "1915-05-08T00:00:00Z",
            "patients.first_name": "Jin",
            "patients.gender": "EFfRsWxPLDnJObCsNVlgTeMaPEZQleQYhYzRyWJj",
            "patients.initials": "jzpfRFEgmotaFetHsbZRjxAwnwekrBEmfdzdc",
            "patients.last_name": "Hwang",
            "patients.mrn": "kXBAkjQZLCtTMtTCoaNatyyiNK",
            "patients.on_hcn": "ReKJyiXJrscctNswYNsGRussVmaozFZBsbOJ",
            "patients.patient_type": "FQGZsnwTKSmVoiGLOpbU",
            "patients.referring_physician": "pEdKupdOMeRVjaRzLNTX",
            "patients.se_num": "eUCWKsXbGyRAOmBTvKSJ",
            "patients.surgical_date": "1911-10-25T00:00:00Z"
            // "patients.mrn": "data.mrn",
            // "patients.first_name": data.first_name,
            // "patients.last_name": data.last_name,
            // "patients.initials": data.initials,
            // "patients.gender": data.gender,
            // "patients.dob": data.dob,
            // "patients.on_hcn": data.on_hcn,
            // "patients.clinical_history": data.clinical_history,
            // "patients.patient_type": data.patient_type,
            // "patients.date_recieved": data.dat_recieved,
            // "patients.referring_physician": data.referring_physician,
            // "patients.date_reported": data.date_reported,
            // "patients.surgical_date": data.surgical_date
        }).then().catch(err => console.log(err));
    },

    updatePatient() {

    },

    deletePatient(id) {
        return axios.delete(API_URL + "/patient/" + id).then(res => {
            console.log(res.data);
        });
    }
}