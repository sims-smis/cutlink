import { useEffect, useState } from "react";
import "./App.css";
import type { CreateResponse, LinksResponse } from "./types";

const API = "http://localhost:3000";

function App() {
  const [url, setUrl] = useState("");
  const [customShort, setCustomShort] = useState("");
  const [expiry, setExpiry] = useState(24);

  const [createdLink, setCreatedLink] = useState("");
  const [links, setLinks] = useState<string[]>([]);
  const [count, setCount] = useState(0);

  async function loadLinks() {
    try {
      const res = await fetch(`${API}/api/v1/links`);
      const data: LinksResponse = await res.json();

      setLinks(data.links);
      setCount(data.count);
    } catch (err) {
      console.error(err);
    }
  }

  useEffect(() => {
    loadLinks();
  }, []);

  async function shortenUrl(e: React.FormEvent) {
    e.preventDefault();

    try {
      const res = await fetch(`${API}/api/v1`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          url,
          short: customShort,
          expiry,
        }),
      });

      const data: CreateResponse = await res.json();

      setCreatedLink(data.short_url);

      setUrl("");
      setCustomShort("");
      setExpiry(24);

      loadLinks();
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div className="container">
      <h1>CutLink</h1>

      <form onSubmit={shortenUrl} className="card">
        <input
          placeholder="Long URL"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          required
        />

        <input
          placeholder="Custom alias (optional)"
          value={customShort}
          onChange={(e) => setCustomShort(e.target.value)}
        />

        <input
          type="number"
          min={1}
          value={expiry}
          onChange={(e) => setExpiry(Number(e.target.value))}
        />

        <button>Shorten URL</button>
      </form>

      {createdLink && (
        <div className="result">
          <h3>Your Short URL</h3>

          <a
            href={createdLink}
            target="_blank"
            rel="noreferrer"
          >
            {createdLink}
          </a>

          <button
            onClick={() =>
              navigator.clipboard.writeText(createdLink)
            }
          >
            Copy
          </button>
        </div>
      )}

      <div className="stats">
        <h2>Total Links: {count}</h2>
      </div>

      <div className="list">
        <h2>Existing Links</h2>

        {links.map((link) => (
          <div key={link} className="link-row">
            <span>{link}</span>

            <button
              onClick={() =>
                navigator.clipboard.writeText(
                  `${API}/${link}`
                )
              }
            >
              Copy
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;