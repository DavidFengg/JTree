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
            });

            return filter;
        }); 
    },
    
    createPatient(data) {
        return axios.post(API_URL + "/patient", {
            "patients.mrn": data.mrn,
            "patients.se_num": data.se_num,
            "patients.first_name": data.first_name,
            "patients.last_name": data.last_name,
            "patients.initials": data.initials,
            "patients.gender": data.gender,
            "patients.dob": data.dob,
            "patients.on_hcn": data.on_hcn,
            "patients.clinical_history": data.clinical_history,
            "patients.patient_type": data.patient_type,
            "patients.date_received": data.date_received,
            "patients.referring_physician": data.referring_physician,
            "patients.date_reported": data.date_reported,
            "patients.surgical_date": data.surgical_date
        }).then(res => console.log(res.data)).catch(err => console.log(err));
    },

    updatePatient(data) {
        return axios.put(API_URL + "/patient/" + data.patient_id, {
            "patients.mrn": data.mrn,
            "patients.se_num": data.se_num,
            "patients.first_name": data.first_name,
            "patients.last_name": data.last_name,
            "patients.initials": data.initials,
            "patients.gender": data.gender,
            "patients.dob": data.dob,
            "patients.on_hcn": data.on_hcn,
            "patients.clinical_history": data.clinical_history,
            "patients.patient_type": data.patient_type,
            "patients.date_received": data.date_received,
            "patients.referring_physician": data.referring_physician,
            "patients.date_reported": data.date_reported,
            "patients.surgical_date": data.surgical_date
        }).then().catch(err => console.log(err));
    },

    deletePatient(id) {
        return axios.delete(API_URL + "/patient/" + id).then(res => {
            console.log(res.data);
        });
    }
}