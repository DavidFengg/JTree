<template>
    <div>
        <b-alert fade dismissed variant="danger"
            :show="alertTimer"
            @dismiss-count-down="changeTimer">
            {{ message }}
        </b-alert>
    </div>
</template>

<script>
export default {
    props: {
      message: String  
    },

    data() {
        return {
            alertTimer: 0
        }
    },

    methods: {
        changeTimer(alertTimer) {
            this.alertTimer = alertTimer;
        }
    },

    watch: {
        // start alert message if message is not empty
        "message": function() {
            if (this.message != "" ) {
                this.alertTimer = 10;
            }
        },

        // send emit back to parent component to notify that alert is done
        "alertTimer": function() {
            if (this.alertTimer == 0) {
                this.$emit("done");
            }
        }
    }
}
</script>