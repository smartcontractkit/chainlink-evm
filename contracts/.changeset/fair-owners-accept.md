---
'@chainlink/contracts': patch
---

Compare ownership nominations against the stored current owner in Ownable2Step rather than msg.sender. This allows the constructor to nominate the deployer when it differs from the initial owner, while rejecting nomination of the actual current owner during construction. Normal owner-only transfers and zero-address cancellation remain unchanged.
