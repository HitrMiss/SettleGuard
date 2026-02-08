import { BrowserProvider } from "ethers";

export function hasMetaMask() {
  return typeof window !== "undefined" && typeof window.ethereum !== "undefined";
}

export async function connectMetaMask() {
  if (!hasMetaMask()) throw new Error("MetaMask is not installed");

  const provider = new BrowserProvider(window.ethereum);
  await provider.send("eth_requestAccounts", []);

  const signer = await provider.getSigner();
  const address = await signer.getAddress();
  const network = await provider.getNetwork();

  return { provider, address, chainId: Number(network.chainId) };
}

/**
 * Force MetaMask to estimate gas by removing any explicit gas settings.
 * This fixes: "transaction gas limit too high (cap: 16777216, tx: 21000000)"
 */
function sanitizeTxRequest(txRequest) {
  const tx = { ...txRequest };

  // Remove explicit gas fields so MetaMask estimates correctly
  delete tx.gas;
  delete tx.gasLimit;

  // Also remove fee overrides unless you intentionally set them
  delete tx.gasPrice;
  delete tx.maxFeePerGas;
  delete tx.maxPriorityFeePerGas;

  // Ensure value is present and hex
  if (tx.value == null) tx.value = "0x0";

  return tx;
}

// Send a transaction using a tx request built by your Go API.
// txRequest fields: { to, from, data, value, gas, maxFeePerGas, maxPriorityFeePerGas, chainId }
export async function sendTxFromApi(txRequest) {
  if (!hasMetaMask()) throw new Error("MetaMask is not installed");
  if (!txRequest?.to) throw new Error("Missing txRequest.to");

  // const clean = sanitizeTxRequest(txRequest);

  // Debug: if gas is still present, something upstream is injecting it
  // if ("gas" in clean || "gasLimit" in clean) {
  //   throw new Error(
  //     `Internal: txRequest still includes gas/gasLimit after sanitize. tx=${JSON.stringify(clean)}`
  //   );
  // }

  // Optional: preflight estimate to catch revert/chain issues early
  // (If you don't want this, you can delete this block.)
  try {
    await window.ethereum.request({
      method: "eth_estimateGas",
      params: [txRequest]
    });
  } catch (e) {
    // If estimate fails due to revert, MetaMask can still show useful error
    // but we keep this info in console for debugging.
    console.warn("eth_estimateGas failed:", e);
  }

  const hash = await window.ethereum.request({
    method: "eth_sendTransaction",
    params: [txRequest]
  });

  return hash; // tx hash
}
