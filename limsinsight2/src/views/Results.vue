<template>
    <div>

    <Alert :message="message" @done="clearMessage"/>

    <!-- Results Table -->
    <b-table :items="results" :fields="fields" hover responsive bordered>
        <template v-slot:cell(action)="data">
            <b-button size="sm" class="mx-1" v-on:click="showModal(data.item)">Edit</b-button>
            <b-button size="sm" class="mx-1" v-on:click="deleteResult(data.item)"> Delete </b-button>
        </template>
    </b-table>

    <!-- Dropdown -->
    <b-dropdown text="Experiment ID">
        <b-dropdown-item v-on:click="updateCurrExperiment(experiment)" v-for="(experiment, i) in getExperiments()" v-bind:key="i">
            {{ experiment["experiments.experiment_id"]}}
        </b-dropdown-item>
    </b-dropdown>

     <!-- Add Result -->
    <b-table-simple hover responsive class="mt-3"> 
        <b-tbody>
            <b-tr>
                <!-- Static field based on Experiment ID selected -->
                <b-td colspan="2">
                    <div class="input-field">
                        <label>Experiment ID</label>
                        <input :disabled="true" :placeholder="selected['experiments.experiment_id']" type="text">
                    </div>
                </b-td>

                <!-- Input fields -->
                <b-td v-for="(field,i) in modified()" v-bind:key="i">
                    <label>{{ field.label }}</label>

                    <!-- input type changes based on field property -->
                    <input v-if="field.type == 'number'" placeholder="" v-model="input[field.key]" type="number">
                    <input v-else placeholder="" v-model="input[field.key]" type="text">
                </b-td>

            </b-tr>
        </b-tbody>
    </b-table-simple>

    <!-- Add button -->
    <b-button class="button" size="m" v-on:click="createResult()"> Add </b-button>


    <!-- Edit Modal -->
    <b-modal id="edit" title="Edit Result Data">
        <b-form-group>
            <div v-for="(field,i) in modified()" v-bind:key="i">
                <label> {{ field.label }}</label>

                <!-- input type changes based on field property -->
                <b-form-input v-if="field.type == 'number'" placeholder="" v-model="edit[field.key]" type="number"></b-form-input>
                <b-form-input v-else placeholder="" v-model="edit[field.key]" type="text"></b-form-input>
            </div>
        </b-form-group>

        <!-- Confirmation button -->
        <b-button class="btn btn-primary" v-on:click="updateResult()">Confirm</b-button>
    </b-modal>

    </div>
</template>

<script>
import Alert from "../components/Alert";
import Shared from "../shared";
import APIService from '../services/APIService';

export default {
    components: {
        Alert
    },

    data() {
        return {
            fields: [
                {key: "results.experiment_id", label: "Experiment ID", sortable: true},
                {key: "results.failed_regions", label: "Failed Regions", sortable: true},
                {key: "results.mean_depth_of_coveage", label: "Mean Depth", type: "number", sortable: true},
                {key: "results.mlpa_pcr", label: "MLPA PCR", sortable: true},
                {key: "results.mutation", label: "Mutation", sortable: true},
                {key: "results.overall_hotspots_threshold", label: "Overall Hotspots Threshold", type: "number", sortable: true},
                {key: "results.overall_quality_threshold", label: "Overall Quality Threshold", type: "number", sortable: true},
                {key: "results.results_id", label: "Results ID", sortable: true},
                {key: "results.uid", label: "UID", sortable: true},
                {key: "results.verification_pcr", label: "Verification PCR", sortable: true},
                "Action"
            ],
            results: [],
            edit: {},
            input: {
                "results.experiment_id": "",
                "results.failed_regions": "",
                "results.mean_depth_of_coveage": "",
                "results.mlpa_pcr": "",
                "results.mutation": "",
                "results.overall_hotspots_threshold": "",
                "results.overall_quality_threshold": "",
                "results.uid": "",
                "results.verification_pcr": "",
            },
            selected: {},
            // Error handling
            message: ""
        }
    },

    methods: {

        // function initalizes edit information and shows the edit modal
        showModal(experiment) {
            // update placeholder information by copying values
            this.edit = Object.assign({}, experiment);
            
            this.$bvModal.show('edit');
        },

        getResults() {
            APIService.getResults().then(data => {
                this.results = data;
                this.$store.dispatch('addUniqueResults', this.results);
            });
        },

        createResult() {
            // check if all fields have been filled
            if (Shared.emptyFields(this.input)) {
                this.updateMessage("Please fill in all fields");
                return;
            }

            // creates a new object with corrected data types 
            let modify = Shared.convert(this.input, this.fields);

            APIService.createResult(modify).then(res => {
                this.getResults();
            });
        },

        updateResult() {
            // creates a new object with corrected data types 
            let modify = Shared.convert(this.edit, this.fields);

            APIService.updateResult(modify).then(res => {
                this.getResults();
            });
        },

        deleteResult(data) {
            let id = data["results.results_id"];

            APIService.deleteResult(id).then(res => {
                if (res == 405) {
                    this.updateMessage("Result with ID: " + id + " cannot be deleted");
                }
                else {
                    this.getSamples();
                }
            });
        },

        // returns a modified version of the fields object with only attributes needed for creating results
        modified() {
            let modified = [];
            for (let i = 0; i < this.fields.length - 1; i++) {
                let key = this.fields[i].key;

                // filter for keys needed in adding results
                if (!(key == "results.results_id" || key == "results.experiment_id" || key == "Action")) {
                    modified.push(this.fields[i]);
                }
            }

            return modified;
        },

        // get the experiments saved from store
        getExperiments() {
            return this.$store.getters.experiments;
        },

        // update the current experiment that the user selected
        updateCurrExperiment(experiment) {
            this.selected = experiment;

            // update the input's experiment id
            this.input["results.experiment_id"] = experiment["experiments.experiment_id"];
        },

        // Updates the message to be sent to the alert component
        updateMessage(message) {
            this.message = message;
        },

        // Clears the message once alert has finished
        clearMessage() {
            this.message = "";
        }
    },

    mounted() {
        this.getResults();
    }
}
</script>