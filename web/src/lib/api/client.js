const DEFAULT_API = 'http://localhost:8080';

export function apiBase() {
  // You can override with Vite env: VITE_API_BASE
  return import.meta.env.VITE_API_BASE || DEFAULT_API;
}

export async function createOrder(payload) {
  const res = await fetch(`${apiBase()}/orders`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  });
  if (!res.ok) throw new Error((await res.json()).error || 'Create order failed');
  return res.json();
}

export async function getOrder(id) {
  const res = await fetch(`${apiBase()}/orders/${id}`);
  if (!res.ok) throw new Error((await res.json()).error || 'Get order failed');
  return res.json();
}

export async function approveOrder(id) {
  const res = await fetch(`${apiBase()}/orders/${id}/approve`, { method: 'POST' });
  if (!res.ok) throw new Error((await res.json()).error || 'Approve failed');
  return res.json();
}

export async function disputeOrder(id, payload) {
  const res = await fetch(`${apiBase()}/orders/${id}/dispute`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  });
  if (!res.ok) throw new Error((await res.json()).error || 'Dispute failed');
  return res.json();
}

export async function reverseOrder(id) {
  const res = await fetch(`${apiBase()}/orders/${id}/reverse`, { method: 'POST' });
  if (!res.ok) throw new Error((await res.json()).error || 'Reverse failed');
  return res.json();
}
