<template>
    <b-table hover :items="patients" :fields="fields">
        <template v-slot:cell(action)="data">
            <b-button size="sm" class="mx-1" v-b-modal="'edit'" v-on:click="updatePlaceholder(data.item)">Edit</b-button>
            <b-button size="sm" class="mx-1" v-on:click="deletePatient(data.item.id)"> Delete </b-button>
        </template>
    </b-table>
</template>

<script>
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
                {key: "date_received", sortable: true},
                {key: "referring_physician", sortable: true},
                {key: "date_reported", sortable: true},
                {key: "surgical_date", sortable: true}
            ],

        };
    }
}
</script>