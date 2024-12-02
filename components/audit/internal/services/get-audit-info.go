package services

import (
	"context"
	"github.com/LerianStudio/midaz/components/audit/internal/adapters/mongodb/audit"
	"github.com/LerianStudio/midaz/pkg"
	"github.com/LerianStudio/midaz/pkg/mopentelemetry"
	"github.com/google/uuid"
)

func (uc *UseCase) GetAuditInfo(ctx context.Context, organizationID uuid.UUID, ledgerID uuid.UUID, transactionID uuid.UUID) (*audit.Audit, error) {
	logger := pkg.NewLoggerFromContext(ctx)
	tracer := pkg.NewTracerFromContext(ctx)

	ctx, span := tracer.Start(ctx, "query.get_audit_info")
	defer span.End()

	logger.Infof("Retrieving audit info")

	auditID := audit.AuditID{
		OrganizationID: organizationID.String(),
		LedgerID:       ledgerID.String(),
		TransactionID:  transactionID.String(),
	}

	auditInfo, err := uc.AuditRepo.FindByID(ctx, audit.TreeCollection, auditID)
	if err != nil {
		mopentelemetry.HandleSpanError(&span, "Failed to get audit info", err)

		logger.Errorf("Error getting audit info: %v", err)

		return nil, err
	}

	return auditInfo, nil
}
