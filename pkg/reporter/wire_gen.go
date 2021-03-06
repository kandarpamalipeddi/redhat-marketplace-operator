// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package reporter

import (
	"context"
	"github.com/redhat-marketplace/redhat-marketplace-operator/pkg/controller"
	"github.com/redhat-marketplace/redhat-marketplace-operator/pkg/managers"
	"github.com/redhat-marketplace/redhat-marketplace-operator/pkg/utils/reconcileutils"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// Injectors from wire.go:

func NewTask(ctx context.Context, reportName ReportName, config2 *Config) (*Task, error) {
	restConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	restMapper, err := managers.NewDynamicRESTMapper(restConfig)
	if err != nil {
		return nil, err
	}
	opsSrcSchemeDefinition := controller.ProvideOpsSrcScheme()
	monitoringSchemeDefinition := controller.ProvideMonitoringScheme()
	olmV1SchemeDefinition := controller.ProvideOLMV1Scheme()
	olmV1Alpha1SchemeDefinition := controller.ProvideOLMV1Alpha1Scheme()
	openshiftConfigV1SchemeDefinition := controller.ProvideOpenshiftConfigV1Scheme()
	localSchemes := controller.ProvideLocalSchemes(opsSrcSchemeDefinition, monitoringSchemeDefinition, olmV1SchemeDefinition, olmV1Alpha1SchemeDefinition, openshiftConfigV1SchemeDefinition)
	scheme, err := managers.ProvideScheme(restConfig, localSchemes)
	if err != nil {
		return nil, err
	}
	clientOptions := getClientOptions()
	cache, err := managers.ProvideNewCache(restConfig, restMapper, scheme, clientOptions)
	if err != nil {
		return nil, err
	}
	client, err := managers.ProvideClient(restConfig, restMapper, scheme, cache, clientOptions)
	if err != nil {
		return nil, err
	}
	logrLogger := _wireLoggerValue
	clientCommandRunner := reconcileutils.NewClientCommand(client, scheme, logrLogger)
	cacheIsIndexed := managers.CacheIsIndexed{}
	cacheIsStarted := managers.StartCache(ctx, cache, logrLogger, cacheIsIndexed)
	uploaderTarget := config2.UploaderTarget
	uploader, err := ProvideUploader(ctx, clientCommandRunner, logrLogger, cacheIsStarted, uploaderTarget)
	if err != nil {
		return nil, err
	}
	task := &Task{
		ReportName: reportName,
		CC:         clientCommandRunner,
		Cache:      cache,
		K8SClient:  client,
		Ctx:        ctx,
		Config:     config2,
		K8SScheme:  scheme,
		Uploader:   uploader,
	}
	return task, nil
}

var (
	_wireLoggerValue = logger
)

func NewReporter(task *Task) (*MarketplaceReporter, error) {
	reporterConfig := task.Config
	client := task.K8SClient
	contextContext := task.Ctx
	scheme := task.K8SScheme
	logrLogger := _wireLogrLoggerValue
	clientCommandRunner := reconcileutils.NewClientCommand(client, scheme, logrLogger)
	reportName := task.ReportName
	meterReport, err := getMarketplaceReport(contextContext, clientCommandRunner, reportName)
	if err != nil {
		return nil, err
	}
	marketplaceConfig, err := getMarketplaceConfig(contextContext, clientCommandRunner)
	if err != nil {
		return nil, err
	}
	v, err := getMeterDefinitions(contextContext, meterReport, clientCommandRunner)
	if err != nil {
		return nil, err
	}
	service, err := getPrometheusService(contextContext, meterReport, clientCommandRunner)
	if err != nil {
		return nil, err
	}
	apiClient, err := provideApiClient(meterReport, service, reporterConfig)
	if err != nil {
		return nil, err
	}
	marketplaceReporter, err := NewMarketplaceReporter(reporterConfig, client, meterReport, marketplaceConfig, v, service, apiClient)
	if err != nil {
		return nil, err
	}
	return marketplaceReporter, nil
}

var (
	_wireLogrLoggerValue = logger
)
