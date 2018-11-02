new Vue({
    el: '#app',
    data() {
        return {
            bucket: null
        }

    },
    computed: {
        isBucketEmpty() {
            return !this.bucket
        }
    },
    created() {
        axios.get('/api/buckets/1')
            .then(({data}) => this.bucket = data)
            .catch(err =>console.log(err))
    },
    methods: {

    }
})