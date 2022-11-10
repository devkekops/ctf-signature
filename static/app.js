var app = new Vue({
    el: '#app',
    data: {
        headers: ["ID", "Date", "From_account", "To_account", "Sum", "Message"],
        results: [],
    },

    methods: {
        getPayment: function() {
            this.results = [];
            axios
                .get('http://0.0.0.0:8080/getPayment')
                .then(response => (
                    this.results = response.data));
        },
        getPayments: function() {
            this.results = [];
            axios
                .get('http://0.0.0.0:8080/getPayments')
                .then(response => (
                    this.results = response.data));
        },
    },
})