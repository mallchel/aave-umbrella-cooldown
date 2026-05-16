# Umbrella Step 3: Address, ABI, and Event Decoding

Date: 2026-05-13

## Confirmed contract mapping

Provided contract (implementation):
- 0x75e8aC0c063B6966E2A9954adEdf39BdE9370197

Transaction target (proxy):
- 0xa484ab92fe32b143aee7019fc1502b1daa522d31

Verification:
- EIP-1967 implementation slot for proxy
  - slot: 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc
  - value: 0x00000000000000000000000075e8ac0c063b6966e2a9954adedf39bde9370197
- Conclusion: proxy points to the provided implementation address.

## Start transaction decoding

Reference tx:
- 0xdbb0140fdb07a1a2c0f778e9674cb3748d70c61aa40e4b6b3eb9332803a2f661

Decoded function call:
- to: 0xa484ab92fe32b143aee7019fc1502b1daa522d31
- input: 0x787a08a6
- method: cooldown()

Emitted event in receipt:
- contract address: 0xa484ab92fe32b143aee7019fc1502b1daa522d31
- topic0: 0xddc8760931d97309f92a4266c6046f83387e6407bcd727e7dd2130bfc430c419

Event payload mapping:
- topic1 (indexed): user = 0xec581f2247ccc58e15f2d3b3558a3af0f8703ed1
- data[0]: amount
- data[1]: endOfCooldown
- data[2]: unstakeWindow

## Event signatures for indexer

Primary event to open withdrawal queue entry:
- topic0: 0xddc8760931d97309f92a4266c6046f83387e6407bcd727e7dd2130bfc430c419

Primary event to mark exit/withdraw execution:
- Withdraw(address,address,address,uint256,uint256)
- topic0: 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db

## Mapping to DB columns

For reconciliation on Withdraw:
- identify user + amount + block window
- set status <- withdrawn
- set withdraw_tx_hash <- tx hash of withdraw

## Notes

- Use proxy address in log filters for runtime indexing on mainnet.
- Keep implementation ABI for decoding (as provided).
- If proxy implementation upgrades, ABI/topic compatibility must be revalidated.
