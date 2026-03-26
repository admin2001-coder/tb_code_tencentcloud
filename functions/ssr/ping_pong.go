package lib

import (
	"github.com/taubyte/go-sdk/event"
)

//export ping
func ping(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	html := `<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<title>Taubyte SSR</title>
	<meta name="description" content="Dependency-free SSR website on Taubyte" />
	<style>
		:root {
			--bg: #0b1020;
			--panel: #121933;
			--panel-2: #1a2447;
			--text: #eef2ff;
			--muted: #aab6e8;
			--accent: #7c9cff;
			--accent-2: #4f6fe8;
			--border: rgba(255,255,255,0.08);
			--shadow: 0 10px 30px rgba(0,0,0,0.35);
		}

		* { box-sizing: border-box; }

		html, body {
			margin: 0;
			padding: 0;
			font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
			background:
				radial-gradient(circle at top left, rgba(124,156,255,0.18), transparent 30%),
				radial-gradient(circle at top right, rgba(79,111,232,0.15), transparent 28%),
				linear-gradient(180deg, #08101f 0%, var(--bg) 100%);
			color: var(--text);
			min-height: 100%;
		}

		a {
			color: inherit;
			text-decoration: none;
		}

		.container {
			max-width: 1120px;
			margin: 0 auto;
			padding: 32px 20px 64px;
		}

		.nav {
			display: flex;
			align-items: center;
			justify-content: space-between;
			gap: 16px;
			padding: 14px 18px;
			border: 1px solid var(--border);
			background: rgba(18,25,51,0.7);
			backdrop-filter: blur(10px);
			border-radius: 18px;
			box-shadow: var(--shadow);
			position: sticky;
			top: 20px;
		}

		.brand {
			display: flex;
			align-items: center;
			gap: 12px;
			font-weight: 700;
			letter-spacing: 0.02em;
		}

		.logo {
			width: 36px;
			height: 36px;
			border-radius: 12px;
			background: linear-gradient(135deg, var(--accent), var(--accent-2));
			display: inline-flex;
			align-items: center;
			justify-content: center;
			font-weight: 800;
			box-shadow: 0 8px 20px rgba(79,111,232,0.45);
		}

		.nav-links {
			display: flex;
			gap: 18px;
			flex-wrap: wrap;
			color: var(--muted);
			font-size: 14px;
		}

		.hero {
			display: grid;
			grid-template-columns: 1.2fr 0.8fr;
			gap: 24px;
			align-items: stretch;
			padding-top: 28px;
		}

		.card {
			border: 1px solid var(--border);
			background: linear-gradient(180deg, rgba(18,25,51,0.88), rgba(14,20,40,0.95));
			border-radius: 24px;
			box-shadow: var(--shadow);
		}

		.hero-copy {
			padding: 34px;
		}

		.badge {
			display: inline-flex;
			align-items: center;
			gap: 8px;
			padding: 8px 12px;
			border-radius: 999px;
			background: rgba(124,156,255,0.12);
			border: 1px solid rgba(124,156,255,0.22);
			color: #dce5ff;
			font-size: 13px;
			font-weight: 600;
		}

		h1 {
			font-size: clamp(34px, 6vw, 62px);
			line-height: 1.02;
			letter-spacing: -0.04em;
			margin: 18px 0 16px;
		}

		.lead {
			font-size: 18px;
			line-height: 1.7;
			color: var(--muted);
			max-width: 62ch;
			margin: 0 0 28px;
		}

		.actions {
			display: flex;
			gap: 14px;
			flex-wrap: wrap;
		}

		.btn {
			display: inline-flex;
			align-items: center;
			justify-content: center;
			padding: 13px 18px;
			border-radius: 14px;
			font-weight: 700;
			border: 1px solid transparent;
			transition: transform 0.15s ease, opacity 0.15s ease;
		}

		.btn:hover {
			transform: translateY(-1px);
		}

		.btn-primary {
			background: linear-gradient(135deg, var(--accent), var(--accent-2));
			color: white;
			box-shadow: 0 10px 22px rgba(79,111,232,0.36);
		}

		.btn-secondary {
			background: rgba(255,255,255,0.03);
			border-color: var(--border);
			color: var(--text);
		}

		.hero-panel {
			padding: 24px;
			display: flex;
			flex-direction: column;
			justify-content: center;
		}

		.stat-grid {
			display: grid;
			grid-template-columns: 1fr 1fr;
			gap: 14px;
		}

		.stat {
			padding: 18px;
			border-radius: 18px;
			background: rgba(255,255,255,0.03);
			border: 1px solid var(--border);
		}

		.stat h3 {
			margin: 0 0 8px;
			font-size: 13px;
			color: var(--muted);
			font-weight: 600;
		}

		.stat p {
			margin: 0;
			font-size: 28px;
			font-weight: 800;
			letter-spacing: -0.03em;
		}

		.section {
			margin-top: 24px;
			padding: 26px;
		}

		.section h2 {
			margin: 0 0 12px;
			font-size: 24px;
			letter-spacing: -0.03em;
		}

		.section p {
			margin: 0 0 18px;
			color: var(--muted);
			line-height: 1.7;
		}

		.features {
			display: grid;
			grid-template-columns: repeat(3, 1fr);
			gap: 16px;
		}

		.feature {
			padding: 18px;
			border-radius: 18px;
			background: rgba(255,255,255,0.03);
			border: 1px solid var(--border);
		}

		.feature strong {
			display: block;
			font-size: 16px;
			margin-bottom: 8px;
		}

		.feature span {
			color: var(--muted);
			line-height: 1.6;
			font-size: 14px;
		}

		.footer {
			margin-top: 26px;
			padding: 18px 6px 0;
			color: var(--muted);
			font-size: 14px;
			text-align: center;
		}

		.code {
			display: inline-block;
			padding: 3px 8px;
			border-radius: 10px;
			background: rgba(255,255,255,0.06);
			border: 1px solid var(--border);
			font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
			font-size: 13px;
			color: #dfe7ff;
		}

		@media (max-width: 900px) {
			.hero,
			.features {
				grid-template-columns: 1fr;
			}

			.nav {
				flex-direction: column;
				align-items: flex-start;
			}
		}
	</style>
</head>
<body>
	<div class="container">
		<header class="nav">
			<div class="brand">
				<div class="logo">T</div>
				<div>Taubyte SSR Site</div>
			</div>
			<nav class="nav-links">
				<a href="#features">Features</a>
				<a href="#about">About</a>
				<a href="#deploy">Deploy</a>
			</nav>
		</header>

		<section class="hero">
			<div class="card hero-copy">
				<div class="badge">Rendered fully on the server</div>
				<h1>Dependency-free SSR website using one Taubyte function.</h1>
				<p class="lead">
					This page is generated directly from your exported Go function with no frontend framework,
					no client rendering requirement, and no extra runtime dependencies beyond the Taubyte SDK.
				</p>
				<div class="actions">
					<a class="btn btn-primary" href="#features">Explore Features</a>
					<a class="btn btn-secondary" href="#deploy">Deployment Notes</a>
				</div>
			</div>

			<div class="card hero-panel">
				<div class="stat-grid">
					<div class="stat">
						<h3>Rendering</h3>
						<p>SSR</p>
					</div>
					<div class="stat">
						<h3>Frontend deps</h3>
						<p>0</p>
					</div>
					<div class="stat">
						<h3>Runtime style</h3>
						<p>Pure Go</p>
					</div>
					<div class="stat">
						<h3>Response</h3>
						<p>HTML</p>
					</div>
				</div>
			</div>
		</section>

		<section class="card section" id="features">
			<h2>What this gives you</h2>
			<p>
				Because the HTML is emitted directly from the function, the website is naturally server-side rendered.
				That means fast first paint, no hydration errors, and easy deployment as a single backend function.
			</p>
			<div class="features">
				<div class="feature">
					<strong>No JS framework required</strong>
					<span>Everything is returned as a normal HTML document from the backend function.</span>
				</div>
				<div class="feature">
					<strong>Easy to extend</strong>
					<span>Add routes, inject dynamic values, or render templates with plain Go strings.</span>
				</div>
				<div class="feature">
					<strong>Portable</strong>
					<span>Works cleanly as a simple HTTP response without bundlers or client-side tooling.</span>
				</div>
			</div>
		</section>

		<section class="card section" id="about">
			<h2>How it works</h2>
			<p>
				The function calls <span class="code">e.HTTP()</span>, gets the HTTP writer, and writes a complete HTML document.
				That is enough for a real SSR website. The browser receives already-rendered markup and displays it immediately.
			</p>
		</section>

		<section class="card section" id="deploy">
			<h2>Deployment notes</h2>
			<p>
				Keep this function attached to an HTTP route in your Taubyte service. When that route is hit,
				the function returns the full page. You can split larger pages into helper functions later,
				but this single-file version is the simplest starting point.
			</p>
		</section>

		<div class="footer">
			Built with a single exported Go function on Taubyte.
		</div>
	</div>
</body>
</html>`

	h.Write([]byte(html))
	return 0
}