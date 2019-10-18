<template>
    <div>

    <!-- Experiments Table -->
    <b-table :items="experiments" :fields="fields" hover responsive bordered>
        <template v-slot:cell(action)="data">
            <b-button size="sm" class="mx-1" v-on:click="showModal(data.item)">Edit</b-button>
            <b-button size="sm" class="mx-1" v-on:click="deleteExperiment(data.item)"> Delete </b-button>
        </template>
    </b-table>

    <!-- Dropdown -->
    <b-dropdown text="Sample ID">
        <b-dropdown-item v-on:click="updateCurrSample(sample)" v-for="sample in getSamples()">
            {{ sample["samples.sample_id"]}}
        </b-dropdown-item>
    </b-dropdown>

     <!-- Add Sample -->
    <b-table-simple hover responsive class="mt-3"> 
        <b-tbody>
            <b-tr>
                <!-- Static field based on Sample ID selected -->
                <b-td colspan="2">
                    <div class="input-field">
                        <label>Sample ID</label>
                        <input :disabled="true" :placeholder="selected['samples.sample_id']" type="text">
                    </div>
                </b-td>

                <!-- Input fields -->
                <b-td v-for="(field,i) in modified()" v-bind:key="i">
                    <label>{{ field.label }}</label>

                    <!-- input type changes based on field property -->
                    <input v-if="field.type == 'date'" placeholder="" v-model="input[field.key]" type="datetime-local">
                    <input v-else-if="field.type == 'checkbox'" placeholder="" v-model="input[field.key]" type="checkbox">
                    <input v-else placeholder="" v-model="input[field.key]" type="text">
                </b-td>

            </b-tr>
        </b-tbody>
    </b-table-simple>

    <!-- Add button -->
    <b-button class="button" size="m" v-on:click="createExperiment()"> Add </b-button>


    <!-- Edit Modal -->
    <b-modal id="edit" title="Edit Experiment Data">
        <b-form-group>
            <div v-for="(field,i) in modified()" v-bind:key="i">
                <label> {{ field.label }}</label>

                <!-- input type changes based on field property -->
                <b-form-checkbox v-if="field.type == 'checkbox'" placeholder="" v-model="edit[field.key]" type="checkbox"></b-form-checkbox>
                <b-form-input v-else placeholder="" v-model="edit[field.key]" type="text"></b-form-input>
            </div>
        </b-form-group>

        <!-- Confirmation button -->
        <b-button class="btn btn-primary" v-on:click="updateExperiment()">Confirm</b-button>
    </b-modal>

    </div>
</template>

<script>
import APIService from '../services/APIService';

export default {
    data() {
        return {
            fields: [
                {key: "experiments.sample_id", label: "Sample ID", sortable: true},
                {key: "experiments.experiment_id", label: "Experiment ID", sortable: true},
                {key: "experiments.study_id", label: "Study ID", sortable: true},
                {key: "experiments.panel_assay_screened", label: "Panel/Assay", sortable: true},
                {key: "experiments.test_date", label: "Test Date", type: "date", sortable: true},
                {key: "experiments.chip_cartridge_barcode", label: "Chip Cartridge Barcode", sortable: true},
                {key: "experiments.complete_date", label: "Complete Date", type: "date", sortable: true},
                {key: "experiments.pcr", label: "PCR", sortable: true},
                {key: "experiments.project_name", label: "Project Name", sortable: true},
                {key: "experiments.priority", label: "Priority", sortable: true},
                {key: "experiments.opened_date", label: "Opened Date", type: "date", sortable: true},  
                {key: "experiments.project_id", label: "Project ID", sortable: true},
                {key: "experiments.has_project_files", label: "Has Project Files", type: "checkbox", sortable: true},
                {key: "experiments.procedure_order_datetime", label: "Procedure Order Date Time", type: "date", sortable: true},
                "Action"
            ],
            experiments: [],
            edit: {},
            input: {
                "experiments.sample_id": "",
                "experiments.study_id": "",
                "experiments.panel_assay_screened": "",
                "experiments.test_date": "",
                "experiments.chip_cartridge_barcode": "",
                "experiments.complete_date": "",
                "experiments.pcr": "",
                "experiments.project_name": "",
                "experiments.priority": "",
                "experiments.opened_date": "",
                "experiments.project_id": "",
                "experiments.has_project_files": "",
                "experiments.procedure_order__datetime": ""
            },
            selected: {}
        }
    },

    methods: {

        // function initalizes edit information and shows the edit modal
        showModal(experiment) {
            // update placeholder information by copying values
            this.edit = Object.assign({}, experiment);
            
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

        getExperiments() {
            APIService.getExperiments().then(data => {
                this.experiments = data;
            });
        },

        createExperiment() {
            let modify = this.convert(this.input);

            APIService.createExperiment(modify).then(res => {
                this.getExperiments();
            });
        },

        updateExperiment() {
            let modify = this.convert(this.edit);

            APIService.updateExperiment(modify).then(res => {
                this.getExperiments();
            });
        },

        deleteExperiment(data) {
            let id = data["experiments.experiment_id"];

            APIService.deleteExperiment(id).then(res => {
                this.getExperiments();
            });
        },

        // returns a modified version of the fields array with only objects needed for creating experiments
        modified() {
            let modified = [];
            for (let i = 0; i < this.fields.length - 1; i++) {
                let key = this.fields[i].key;

                // adds only the objects needed for adding experiemtns
                if (!(key == "experiments.experiment_id" || key == "experiments.sample_id" || key == "Action")) {
                    modified.push(this.fields[i]);
                }
            }

            return modified;
        },

        // get the patients saved from store
        getSamples() {
            return this.$store.getters.samples;
        },

        // update the current sample that the user selected
        updateCurrSample(sample) {
            this.selected = sample;

            // update the input's patient id
            this.input["experiments.sample_id"] = sample["samples.sample_id"];
        }


    },

    mounted() {
        this.getExperiments();
    }
}
</script>