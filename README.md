<h1 align="center"><strong>Tracing for Mojito</strong></h1>
<p align="center">
    <a href="https://goreportcard.com/report/github.com/dragse/mojito-extension-tracing" alt="Go Report Card">
        <img src="https://goreportcard.com/badge/github.com/dragse/mojito-extension-tracing" /></a>
	<a href="https://github.com/dragse/mojito-extension-tracing" alt="Go Version">
        <img src="https://img.shields.io/github/go-mod/go-version/go-mojito/extension-plausible.svg" /></a>
	<a href="https://godoc.org/github.com/dragse/mojito-extension-tracing" alt="GoDoc reference">
        <img src="https://img.shields.io/badge/godoc-reference-blue.svg"/></a>
	<a href="https://github.com/dragse/mojito-extension-tracing/blob/main/LICENSE" alt="Licence">
        <img src="https://img.shields.io/github/license/Ileriayo/markdown-badges?style=flat-square" /></a>
	<a href="https://makeapullrequest.com">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="PRs Welcome"></a>
</p>
<p align="center">
    <a href="https://go.dev/" alt="Made with Go">
        <img src="https://ForTheBadge.com/images/badges/made-with-go.svg" /></a>

</p>
<p align="center">
Tracing for Mojito provides easy tracking and configuration of OpenTracing Provider.</p>

<p align="center"><strong>SonarCloud Report</strong></p>
<p align="center">
    <a href="https://sonarcloud.io/summary/overall?id=dragse_mojito-extension-tracing" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_extension-plausible&metric=alert_status" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=dragse_mojito-extension-tracing" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_extension-plausible&metric=sqale_rating" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=dragse_mojito-extension-tracing" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_extension-plausible&metric=reliability_rating" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=dragse_mojito-extension-tracing" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_extension-plausible&metric=security_rating" /></a>
	<br>
    <a href="https://sonarcloud.io/summary/overall?id=dragse_mojito-extension-tracing" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_extension-plausible&metric=vulnerabilities" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=dragse_mojito-extension-tracing" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_extension-plausible&metric=code_smells" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=dragse_mojito-extension-tracing" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_extension-plausible&metric=bugs" /></a>
</p>
<p align="center"><strong>Supported Exporter</strong></p>
<ul style="list-style: none;">
    <li><input type="checkbox" checked><a href="https://www.jaegertracing.io/">Jaeger</a></li>
</ul>

<p align="center"><strong>Documentation</strong></p>
<h2>Enabling Tracing</h2>
<p>
    Enabling Open-Telemetry Tracing is as simple as registering a middleware on your router.
</p>
<pre>
<code>
import (
    tracing_extension "github.com/dragse/mojito-extension-tracing"
    "github.com/go-mojito/mojito"
)</code>

<code>
func init() {
    tracing_extension.Configure(
        tracing_extension.JAEGER, // Use JAEGER as Provider
        "mojito-service", map[string]any{
            "environment": "productive", // Custom Tags for all Traces
        },
        mojito_extension_tracing.ExporterConfig{
            ProviderURL: "http://localhost:14268/api/traces", // Jaeger Endpoint
    })
}</code>

<code>
func main() {
    mojito.WithMiddleware(tracing_extension.Middleware)
}</code>
</pre>

<h2>Custom more detailed Tracing</h2>
<p>
    For more detailed Information  you can trace every single function which is dynamically connected to the method Tracing
</p>
<pre>
<code>
func HomeHandler(ctx mojito.RendererContext, cache mojito.Cache) {
	span := mojito_extension_tracing.StartTracing(ctx, "Home Handler")
	defer span.End()

	span.AddEvent("Load Cache")
	var lastVisit time.Time
	cache.GetOrDefault("lastVisit", &lastVisit, time.Now())
	span.SetAttributes(attribute.String("lastVisit", lastVisit.String()))
	span.AddEvent("Set new lastVisit-Variable")
	cache.Set("lastVisit", time.Now())

	span.AddEvent("Set Render-Information")
	ctx.ViewBag().Set("lastVisit", lastVisit)
	ctx.MustView("home")
}
</code>
</pre>
