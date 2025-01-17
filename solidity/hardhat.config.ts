import {HardhatUserConfig} from "hardhat/config"
import "hardhat-dependency-compiler"
import "@nomiclabs/hardhat-ethers"

import './tasks/task'

const config: HardhatUserConfig = {
    defaultNetwork: "localhost",
    networks: {
        hardhat: {
            mining: {
                interval: 1000
            }
        },
        localhost: {
            url: `${process.env.LOCAL_URL || "http://localhost:8545"}`,
        },
    },
    solidity: {
        compilers: [
            {
                version: "0.8.0",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200
                    }
                }
            },
            {
                version: "0.8.1",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200
                    }
                }
            },
            {
                version: "0.8.2",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200
                    }
                }
            },
        ]
    },
    dependencyCompiler: {
        paths: [
            "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol",
            "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol",
        ],
    },
};

export default config;