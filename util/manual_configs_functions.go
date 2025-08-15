package util

import (
	"github.com/omec-project/amf/factory"
	"github.com/omec-project/amf/logger"
	"github.com/omec-project/openapi/Nnrf_NFDiscovery"
	"github.com/omec-project/openapi/models"
)

func SearchNFInstancesWithManualConfig(manualConfig *factory.ManualConfig, targetNfType, requestNfType models.NfType, param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (models.SearchResult, error) {
	logger.AppLog.Infof("Using manual configuration for NF discovery")

	if manualConfig == nil {
		return models.SearchResult{}, nil
	}

	// Create a SearchResult based on the manual configuration
	result := models.SearchResult{
		NfInstances: make([]models.NfProfile, 0),
	}

	result.NfInstances = append(result.NfInstances, manualConfig.NFs[targetNfType]...)

	return result, nil
}
