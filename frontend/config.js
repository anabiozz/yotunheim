const config = {
    baseDomain: 'http://localhost:8888',
    timeInterval: 30000,
    initialState: '',
    defaultChartConfig: {
        chartType: 'LineChart',
        columns: [],
        rows: [],
        options: {
            chartArea: {
                left: 0,
                right: 0,
                top: 8,
                bottom: 5,
                width: '100%'
            },
            backgroundColor: {
                fill: 'transparent'
            },
            subtitle: '',
            legend: { position: 'bottom' },
            hAxis: {
                gridlines: { count: 5 },
                viewWindow: { max: 3000 }
            },
            vAxis: {}
        },
        width: '100%',
        height: '150px'
    },
    defaultTableConfig: {
        header: [
        { name: 'Machine Name', sortable: false },
        { name: 'CPU', sortable: false },
        { name: 'Memory', sortable: false },
        { name: 'Disk', sortable: false }],
        columns: ['inst_location', 'cpu', 'mem', 'disk']
    },
    defaultTableCallsConfig: {
        header: [
            { name: 'Zipkin ID', sortable: true },
            { name: 'conferenceName', sortable: false },
            { name: 'Dial string', sortable: false },
            { name: 'Focus URI', sortable: false },
            { name: 'Call type', sortable: false },
            { name: 'Start time', sortable: false },
            { name: 'Duration', sortable: false },
            { name: 'Endpoint type', sortable: false },
            { name: 'Endpoint ID ', sortable: false },
            { name: 'Tenant ID', sortable: false },
            { name: 'Region', sortable: false },
            { name: 'Packet loss', sortable: false },
            { name: 'Jitter', sortable: false },
            { name: 'Latency', sortable: false }
        ],
        columns: ['zipkinID',
                'conferenceName',
                'dialString',
                'focusURI',
                'callType',
                'startTime',
                'duration',
                'endpointType',
                'endpointID',
                'tenantID',
                'region',
                'packetLoss',
                'jitter',
                'lattency'],
    },
    defaultTableDetailConfig: {
        header: [],
        columns: ['name','value']
    },
}
export default config
