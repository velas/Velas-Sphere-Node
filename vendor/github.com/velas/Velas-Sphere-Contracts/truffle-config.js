module.exports = {
    networks: {
        local: {
            host: "localhost",
            port: 8545,
            network_id: "*" // Match any network id
        },
        ropsten: {
            network_id: 3,
            host: "localhost",
            port: 8545,
            gas: 2900000
        }
    },
    mocha: {
        reporter: 'eth-gas-reporter',
        reporterOptions: {
            currency: 'USD',
            onlyCalledMethods: true
        }
    },
    coverage: {
        host: "localhost",
        network_id: "*",
        port: 8555,
        gas: 0xfffffffffff,
        gasPrice: 0x01
    }
};
