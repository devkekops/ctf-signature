var app = new Vue({
    el: '#app',
    data: {
        headers: ["ID", "Date", "From_account", "To_account", "Sum", "Message"],
        results: [],
        paymentId: '',
        secretKey: 'top_secret',
        offset: '1209600'
    },
    methods: {
        getPayment: function() {
            this.results = [];
            const headers = {'X-SIG-TOKEN': CryptoJS.MD5('id=' + this.paymentId + 'secret_key=' + this.secretKey)}
            axios
                .get('http://0.0.0.0:8080/api/getPayment?id=' + this.paymentId, {headers})
                .then(response => (
                    this.results = [response.data]), this.paymentId = '');
        },
        getPayments: function() {
            this.results = [];
            const headers = {'X-SIG-TOKEN': CryptoJS.MD5('offset=' + this.offset + 'secret_key=' + this.secretKey)}
            axios
                .get('http://0.0.0.0:8080/api/getPayments?offset=' + this.offset, {headers})
                .then(response => (
                    this.results = response.data));
        },
    },
})