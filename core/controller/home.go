// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"context"
	"time"

	"github.com/clivern/walnut/core/component"

	"github.com/gin-gonic/gin"
)

// Home controller
func Home(c *gin.Context, tracing *component.Tracing) {
	profiler := component.NewProfiler(context.Background())

	defer profiler.WithCorrelation(
		c.Request.Header.Get("X-Correlation-ID"),
	).LogDuration(
		time.Now(),
		"homeController",
		component.Info,
	)

	if tracing.IsEnabled() {
		span := tracing.GetTracer().StartSpan("api.homeController")
		span.SetTag("correlation_id", c.Request.Header.Get("X-Correlation-ID"))

		defer span.Finish()
	}

	proxy := component.NewProxy(
		context.Background(),
		c.Request,
		c.Writer,
		"http://127.0.0.1:8000/_health?v=23&fg=34&ok=372he",
		c.Request.Header.Get("X-Correlation-ID"),
	)

	proxy.Redirect()
}
