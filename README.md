# Wallet Emulator

Wallet emulator is a small web app that you can use to test and debug the [wallet-lib](https://github.com/clovrlabs/wallet-lib) on its own, without having to package it for mobile.

It consists of a small backend that reproduces how the Mobile app would work, and a frontend that is just the presentational part of the wallet. So instead of a CLI you can click on the buttons just as you would do in mobile.

# Developing
- Make sure you have go 1.19 installed
- Initialize the submodules
    - `git submodule update --init --recursive`
    - `git submodule update --recursive`
- Run `go get`
- Add some breakpoints!
- If you have VSCode, all you have to do is go to the `Run and Debug` tab, and select in which environment you want to run the wallet.
    - Emulator Develop: You need to have all services running locally in order for it to connect to regtest network.
    - Emulator Staging: This will connect to the regtest network that is deployed in staging, so it should work out of the box.
    - Emulator Production: This will connect to the mainnet network. Warning: If you have funds in your wallet and you delete the `data` directory, you will lose your funds.
