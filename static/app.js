var app = new Vue({
    el: '#app',
    data: {
        headers: ["ID", "Date", "From_account", "To_account", "Sum", "Message"],
        results: [],
        paymentId: '',
    },

    methods: {
        getPayment: function() {
            this.results = [];
            axios
                .get('http://0.0.0.0:8080/getPayment?id=' + this.paymentId)
                .then(response => (
                    this.results = [response.data]), this.paymentId = '');
        },
        getPayments: function() {
            this.results = [];
            axios
                .get('http://0.0.0.0:8080/getPayments?offset=1209600')
                .then(response => (
                    this.results = response.data));
        },
    },
})