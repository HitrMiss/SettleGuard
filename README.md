# Welcome to 🛡️SettleGuard
⚡*Could have used more caffeine during development*⚡

## Contracts - Sepolia testnet
```json
{
  "network": "sepolia",
  "contracts": {
    "SettleGuard#Governance": "0x2e814CeBa9B63412552AC30D56c9bc326D707a94",
    "SettleGuard#MockUSDC": "0xa9399a841da56fD9718c800e3082827AdE97d6ad",
    "SettleGuard#BondVault": "0x20d2Cd4f77082683f7367Fd46A8BF54097Ac68Fe",
    "SettleGuard#CategoryRegistry": "0x2F316264c65B557BcF17bE446F59Dd658C7894f2",
    "SettleGuard#PaymentVault": "0x357EF41da52817F617aA2D0D34ceBB5DB08fb9Ca",
    "SettleGuard#ProfileRegistry": "0x5cF065dd7B9b8B644e5F303Ee3CA7E6d8570336c",
    "SettleGuard#SettlementEngine": "0x700C50363F845207d560462Bb4aDa20dB0d6763B"
  }
}
```
## Merchant Setup
```shell
cd api
go run ./cmd/test-tool
```
Option 4 will generate a merchant certificate for signing transactions.  
The compressed public key should be store on your ENS record as custom   
You can also generate the encryption key without using the tool. Details to be provided later ;)
```dotenv
Key: settleGuard.p256
Value: 03c69e7e89fe249e6024779993f20a554d26ac45024417344b983b4df492f395e9
```
[Key only shown for test purposes.]  
⚠️ Keep merchant_test.pem private and safe ⚠️  
How? The Wallet that owns your ENS name is not the wallet that needs to receive the funds. In practice is good to swap
wallets to avoid attack vectors. When creating a payment transaction you supply 2 addresses.   

Your public key will be store per transaction in case you need to switch your private key.

```solidity
address _merchantIdentity, // The wallet address that owns your MerchantAccount (Any wallet with the public key)
address _merchantPayout,   // The wallet where the funds will be released to from the buyer. 
```
This is traditional payment gateway flow.

