<template>
    <div>

    <!-- ResultDetails Table -->
    <b-table :items="resultdetails" :fields="fields" hover responsive bordered>
        <template v-slot:cell(action)="data">
            <b-button size="sm" class="mx-1" v-on:click="showModal(data.item)">Edit</b-button>
            <b-button size="sm" class="mx-1" v-on:click="deleteResultDetail(data.item)"> Delete </b-button>
        </template>
    </b-table>

    <!-- Dropdown -->
    <b-dropdown text="Result ID">
        <b-dropdown-item v-on:click="updateCurrResult(result)" v-for="(result, i) in getResults()" v-bind:key="i">
            {{ result["results.results_id"]}}
        </b-dropdown-item>
    </b-dropdown>

     <!-- Add ResultDetail -->
    <b-table-simple hover responsive class="mt-3"> 
        <b-tbody>
            <b-tr>
                <!-- Static field based on Result ID selected -->
                <b-td colspan="2">
                    <div class="input-field">
                        <label>Result ID</label>
                        <input :disabled="true" :placeholder="selected['results.results_id']" type="text">
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
    <b-button class="button" size="m" v-on:click="createResultDetail()"> Add </b-button>


    <!-- Edit Modal -->
    <b-modal id="edit" title="Edit Result Data">
        <b-form-group>
            <div v-for="(field,i) in modified()" v-bind:key="i">
                <label> {{ field.label }}</label>

                <!-- input type changes based on field property -->
                <b-form-checkbox v-if="field.type == 'checkbox'" placeholder="" v-model="edit[field.key]" type="checkbox"></b-form-checkbox>
                <b-form-input v-else placeholder="" v-model="edit[field.key]" type="text"></b-form-input>
            </div>
        </b-form-group>

        <!-- Confirmation button -->
        <b-button class="btn btn-primary" v-on:click="updateResultDetail()">Confirm</b-button>
    </b-modal>

    </div>
</template>

<script>
import APIService from '../services/APIService';

export default {
    data() {
        return {
            fields: [
                {key: "resultdetails.results_id" , label: "Results ID", sortable: true},
                {key: "resultdetails.VAF", label: "VAF", type: "number", sortable: true},
                {key: "resultdetails.c_nomenclature", label: "C.Nomenclature", sortable: true},
                {key: "resultdetails.coverage", label: "Coverage", type: "number", sortable: true},
                {key: "resultdetails.exon", label: "Exon", type: "number", sortable: true},
                {key: "resultdetails.gene", label: "Gene", sortable: true},
                {key: "resultdetails.p_nomenclature", label: "P.Nomenclature", sortable: true},
                {key: "resultdetails.pcr", label: "PCR", sortable: true},
                {key: "resultdetails.quality_score", label: "Quality Score", type: "number", sortable: true},
                {key: "resultdetails.result", label: "Result", sortable: true},
                {key: "resultdetails.results_details_id", label: "Results Details ID", sortable: true},
                {key: "resultdetails.risk_score" , label: "Risk Score", type: "number", sortable: true},
                {key: "resultdetails.uid", label: "UID", sortable: true},
                "Action"
            ],
            resultdetails: [],
            edit: {},
            input: {
                "resultdetails.VAF": "",
                "resultdetails.c_nomenclature": "",
                "resultdetails.coverage": "",
                "resultdetails.exon": "",
                "resultdetails.gene": "",
                "resultdetails.p_nomenclature": "",
                "resultdetails.pcr": "",
                "resultdetails.quality_score": "",
                "resultdetails.result": "",
                "resultdetails.results_id" : "",
                "resultdetails.risk_score" : "",
                "resultdetails.uid": "",
            },
            selected: {}
        }
    },

    methods: {

        // function initalizes edit information and shows the edit modal
        showModal(result) {
            // update placeholder information by copying values
            this.edit = Object.assign({}, result);
            
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

        getResultDetails() {
            APIService.getResultDetails().then(data => {
                this.resultdetails = data;
            });
        },

        createResultDetail() {
            let modify = this.convert(this.input);

            APIService.createResultDetail(modify).then(res => {
                this.getResultDetails();
            });
        },

        updateResultDetail() {
            let modify = this.convert(this.edit);

            APIService.updateResultDetail(modify).then(res => {
                this.getResultDetails();
            });
        },

        deleteResultDetail(data) {
            let id = data["resultdetails.results_details_id"];

            APIService.deleteResultDetail(id).then(res => {
                this.getResultDetails();
            });
        },

        // returns a modified version of the fields array with only objects needed for creating result details
        modified() {
            let modified = [];
            for (let i = 0; i < this.fields.length - 1; i++) {
                let key = this.fields[i].key;

                // adds only the objects needed for adding result details
                if (!(key == "resultdetails.results_details_id" || key == "resultdetails.results_id" || key == "Action")) {
                    modified.push(this.fields[i]);
                }
            }

            return modified;
        },

        // get the results saved from store
        getResults() {
            return this.$store.getters.results;
        },

        // update the current result that the user selected
        updateCurrResult(result) {
            this.selected = result;

            // update the input's result id
            this.input["resultdetails.results_id"] = result["results.results_id"];
        }
    },

    mounted() {
        this.getResultDetails();
    }
}
</script>