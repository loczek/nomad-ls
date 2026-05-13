package validation

import (
	"context"
	"fmt"
	"log"
	"slices"

	"github.com/hashicorp/hcl-lang/decoder"
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/reference"
	"github.com/hashicorp/hcl/v2"
)

func UnreferencedOrigins(ctx context.Context, pathCtx *decoder.PathContext) lang.DiagnosticsMap {
	diagsMap := make(lang.DiagnosticsMap)

	for _, origin := range pathCtx.ReferenceOrigins {
		localOrigin, ok := origin.(reference.LocalOrigin)

		if !ok {
			continue
		}

		address := localOrigin.Address()

		supported := []string{"var", "local"}
		firstStep := address[0].String()
		if !slices.Contains(supported, firstStep) {
			continue
		}

		_, ok = pathCtx.ReferenceTargets.Match(localOrigin)
		if !ok {
			// target not found
			fileName := origin.OriginRange().Filename
			d := &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  fmt.Sprintf("No declaration found for %q", address),
				Subject:  origin.OriginRange().Ptr(),
			}
			diagsMap[fileName] = diagsMap[fileName].Append(d)

			continue
		}
	}

	return diagsMap
}

func DisplayUnreferencedOrigins(ctx context.Context, pathCtx *decoder.PathContext) lang.DiagnosticsMap {
	diagsMap := make(lang.DiagnosticsMap)

	for _, origin := range pathCtx.ReferenceOrigins {
		localOrigin, ok := origin.(reference.LocalOrigin)

		if !ok {
			continue
		}

		address := localOrigin.Address()

		_, ok = pathCtx.ReferenceTargets.Match(localOrigin)
		fileName := origin.OriginRange().Filename
		d := &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  fmt.Sprintf("No declaration found for %q", address),
			Subject:  origin.OriginRange().Ptr(),
		}
		diagsMap[fileName] = diagsMap[fileName].Append(d)
	}

	log.Printf("diagsMap: %+v", diagsMap)

	return diagsMap
}
