// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package controller

import (
	"context"
	"github.com/aquaproj/aqua/pkg/config"
	"github.com/aquaproj/aqua/pkg/config-finder"
	"github.com/aquaproj/aqua/pkg/config-reader"
	"github.com/aquaproj/aqua/pkg/controller/cp"
	exec2 "github.com/aquaproj/aqua/pkg/controller/exec"
	"github.com/aquaproj/aqua/pkg/controller/generate"
	"github.com/aquaproj/aqua/pkg/controller/generate-registry"
	"github.com/aquaproj/aqua/pkg/controller/initcmd"
	"github.com/aquaproj/aqua/pkg/controller/install"
	"github.com/aquaproj/aqua/pkg/controller/list"
	"github.com/aquaproj/aqua/pkg/controller/which"
	"github.com/aquaproj/aqua/pkg/download"
	"github.com/aquaproj/aqua/pkg/exec"
	"github.com/aquaproj/aqua/pkg/github"
	"github.com/aquaproj/aqua/pkg/install-registry"
	"github.com/aquaproj/aqua/pkg/installpackage"
	"github.com/aquaproj/aqua/pkg/link"
	"github.com/aquaproj/aqua/pkg/runtime"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/go-osenv/osenv"
	"net/http"
)

// Injectors from wire.go:

func InitializeListCommandController(ctx context.Context, param *config.Param, httpClient *http.Client) *list.Controller {
	fs := afero.NewOsFs()
	configFinder := finder.NewConfigFinder(fs)
	configReader := reader.New(fs)
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	installer := registry.New(param, gitHubContentFileDownloader, fs)
	controller := list.NewController(configFinder, configReader, installer)
	return controller
}

func InitializeGenerateRegistryCommandController(ctx context.Context, param *config.Param, httpClient *http.Client) *genrgst.Controller {
	fs := afero.NewOsFs()
	repositoriesService := github.New(ctx)
	controller := genrgst.NewController(fs, repositoriesService)
	return controller
}

func InitializeInitCommandController(ctx context.Context, param *config.Param) *initcmd.Controller {
	repositoriesService := github.New(ctx)
	fs := afero.NewOsFs()
	controller := initcmd.New(repositoriesService, fs)
	return controller
}

func InitializeGenerateCommandController(ctx context.Context, param *config.Param, httpClient *http.Client) *generate.Controller {
	fs := afero.NewOsFs()
	configFinder := finder.NewConfigFinder(fs)
	configReader := reader.New(fs)
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	installer := registry.New(param, gitHubContentFileDownloader, fs)
	fuzzyFinder := generate.NewFuzzyFinder()
	versionSelector := generate.NewVersionSelector()
	controller := generate.New(configFinder, configReader, installer, repositoriesService, fs, fuzzyFinder, versionSelector)
	return controller
}

func InitializeInstallCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) *install.Controller {
	fs := afero.NewOsFs()
	configFinder := finder.NewConfigFinder(fs)
	configReader := reader.New(fs)
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	installer := registry.New(param, gitHubContentFileDownloader, fs)
	packageDownloader := download.NewPackageDownloader(repositoriesService, rt, httpDownloader)
	linker := link.New()
	executor := exec.New()
	installpackageInstaller := installpackage.New(param, packageDownloader, rt, fs, linker, executor)
	controller := install.New(param, configFinder, configReader, installer, installpackageInstaller, fs, rt)
	return controller
}

func InitializeWhichCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) which.Controller {
	fs := afero.NewOsFs()
	configFinder := finder.NewConfigFinder(fs)
	configReader := reader.New(fs)
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	installer := registry.New(param, gitHubContentFileDownloader, fs)
	osEnv := osenv.New()
	linker := link.New()
	controller := which.New(param, configFinder, configReader, installer, rt, osEnv, fs, linker)
	return controller
}

func InitializeExecCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) *exec2.Controller {
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	packageDownloader := download.NewPackageDownloader(repositoriesService, rt, httpDownloader)
	fs := afero.NewOsFs()
	linker := link.New()
	executor := exec.New()
	installer := installpackage.New(param, packageDownloader, rt, fs, linker, executor)
	configFinder := finder.NewConfigFinder(fs)
	configReader := reader.New(fs)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	registryInstaller := registry.New(param, gitHubContentFileDownloader, fs)
	osEnv := osenv.New()
	controller := which.New(param, configFinder, configReader, registryInstaller, rt, osEnv, fs, linker)
	execController := exec2.New(installer, controller, executor, osEnv, fs)
	return execController
}

func InitializeCopyCommandController(ctx context.Context, param *config.Param, httpClient *http.Client, rt *runtime.Runtime) *cp.Controller {
	repositoriesService := github.New(ctx)
	httpDownloader := download.NewHTTPDownloader(httpClient)
	packageDownloader := download.NewPackageDownloader(repositoriesService, rt, httpDownloader)
	fs := afero.NewOsFs()
	linker := link.New()
	executor := exec.New()
	installer := installpackage.New(param, packageDownloader, rt, fs, linker, executor)
	configFinder := finder.NewConfigFinder(fs)
	configReader := reader.New(fs)
	gitHubContentFileDownloader := download.NewGitHubContentFileDownloader(repositoriesService, httpDownloader)
	registryInstaller := registry.New(param, gitHubContentFileDownloader, fs)
	osEnv := osenv.New()
	controller := which.New(param, configFinder, configReader, registryInstaller, rt, osEnv, fs, linker)
	cpController := cp.New(param, installer, fs, rt, controller)
	return cpController
}
