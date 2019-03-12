// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cli

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/google/wire"
	"github.com/windmilleng/tilt/internal/build"
	"github.com/windmilleng/tilt/internal/container"
	"github.com/windmilleng/tilt/internal/demo"
	"github.com/windmilleng/tilt/internal/docker"
	"github.com/windmilleng/tilt/internal/dockercompose"
	"github.com/windmilleng/tilt/internal/dockerfile"
	"github.com/windmilleng/tilt/internal/engine"
	"github.com/windmilleng/tilt/internal/hud"
	"github.com/windmilleng/tilt/internal/hud/server"
	"github.com/windmilleng/tilt/internal/k8s"
	"github.com/windmilleng/tilt/internal/minikube"
	"github.com/windmilleng/tilt/internal/store"
	"github.com/windmilleng/tilt/internal/tiltfile"
	"k8s.io/apimachinery/pkg/version"
	"time"
)

// Injectors from wire.go:

func wireDemo(ctx context.Context, branch demo.RepoBranch) (demo.Script, error) {
	v := provideClock()
	renderer := hud.NewRenderer(v)
	headsUpDisplay, err := hud.NewDefaultHeadsUpDisplay(renderer)
	if err != nil {
		return demo.Script{}, err
	}
	clientConfig := k8s.ProvideClientConfig()
	kubeContext, err := k8s.ProvideKubeContext(clientConfig)
	if err != nil {
		return demo.Script{}, err
	}
	env := k8s.ProvideEnv(kubeContext)
	portForwarder := k8s.ProvidePortForwarder()
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, portForwarder, namespace, kubectlRunner, clientConfig)
	podWatcher := engine.NewPodWatcher(client)
	nodeIP, err := k8s.DetectNodeIP(ctx, env)
	if err != nil {
		return demo.Script{}, err
	}
	serviceWatcher := engine.NewServiceWatcher(client, nodeIP)
	reducer := _wireReducerValue
	storeLogActionsFlag := provideLogActions()
	storeStore := store.NewStore(reducer, storeLogActionsFlag)
	podLogManager := engine.NewPodLogManager(client)
	portForwardController := engine.NewPortForwardController(client)
	fsWatcherMaker := engine.ProvideFsWatcherMaker()
	timerMaker := engine.ProvideTimerMaker()
	watchManager := engine.NewWatchManager(fsWatcherMaker, timerMaker)
	syncletManager := engine.NewSyncletManager(client)
	syncletBuildAndDeployer := engine.NewSyncletBuildAndDeployer(syncletManager, client)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	minikubeClient := minikube.ProvideMinikubeClient()
	dockerEnv, err := docker.ProvideEnv(ctx, env, runtime, minikubeClient)
	if err != nil {
		return demo.Script{}, err
	}
	clientClient, err := docker.ProvideDockerClient(ctx, dockerEnv)
	if err != nil {
		return demo.Script{}, err
	}
	version, err := docker.ProvideDockerVersion(ctx, clientClient)
	if err != nil {
		return demo.Script{}, err
	}
	cli, err := docker.DefaultClient(ctx, clientClient, version)
	if err != nil {
		return demo.Script{}, err
	}
	containerUpdater := build.NewContainerUpdater(cli)
	analytics, err := provideAnalytics()
	if err != nil {
		return demo.Script{}, err
	}
	localContainerBuildAndDeployer := engine.NewLocalContainerBuildAndDeployer(containerUpdater, analytics, env)
	labels := _wireLabelsValue
	dockerImageBuilder := build.NewDockerImageBuilder(cli, labels)
	imageBuilder := build.DefaultImageBuilder(dockerImageBuilder)
	cacheBuilder := build.NewCacheBuilder(cli)
	clock := build.ProvideClock()
	execCustomBuilder := build.NewExecCustomBuilder(cli, dockerEnv, clock)
	engineUpdateModeFlag := provideUpdateModeFlag()
	updateMode, err := engine.ProvideUpdateMode(engineUpdateModeFlag, env, runtime)
	if err != nil {
		return demo.Script{}, err
	}
	imageBuildAndDeployer := engine.NewImageBuildAndDeployer(imageBuilder, cacheBuilder, execCustomBuilder, client, env, analytics, updateMode, clock, runtime)
	dockerComposeClient := dockercompose.NewDockerComposeClient(dockerEnv)
	imageAndCacheBuilder := engine.NewImageAndCacheBuilder(imageBuilder, cacheBuilder, execCustomBuilder, updateMode)
	dockerComposeBuildAndDeployer := engine.NewDockerComposeBuildAndDeployer(dockerComposeClient, cli, imageAndCacheBuilder, clock)
	buildOrder := engine.DefaultBuildOrder(syncletBuildAndDeployer, localContainerBuildAndDeployer, imageBuildAndDeployer, dockerComposeBuildAndDeployer, env, updateMode, runtime)
	compositeBuildAndDeployer := engine.NewCompositeBuildAndDeployer(buildOrder)
	buildController := engine.NewBuildController(compositeBuildAndDeployer)
	imageReaper := build.NewImageReaper(cli)
	imageController := engine.NewImageController(imageReaper)
	globalYAMLBuildController := engine.NewGlobalYAMLBuildController(client)
	tiltfileLoader := tiltfile.ProvideTiltfileLoader(analytics, dockerComposeClient)
	configsController := engine.NewConfigsController(tiltfileLoader)
	dockerComposeEventWatcher := engine.NewDockerComposeEventWatcher(dockerComposeClient)
	dockerComposeLogManager := engine.NewDockerComposeLogManager(dockerComposeClient)
	profilerManager := engine.NewProfilerManager()
	analyticsReporter := engine.ProvideAnalyticsReporter(analytics, storeStore)
	upper := engine.NewUpper(ctx, headsUpDisplay, podWatcher, serviceWatcher, storeStore, podLogManager, portForwardController, watchManager, buildController, imageController, globalYAMLBuildController, configsController, dockerComposeEventWatcher, dockerComposeLogManager, profilerManager, syncletManager, analyticsReporter)
	script := demo.NewScript(upper, headsUpDisplay, client, env, storeStore, branch, runtime, tiltfileLoader)
	return script, nil
}

var (
	_wireReducerValue = engine.UpperReducer
	_wireLabelsValue  = dockerfile.Labels{}
)

func wireThreads(ctx context.Context) (Threads, error) {
	v := provideClock()
	renderer := hud.NewRenderer(v)
	headsUpDisplay, err := hud.NewDefaultHeadsUpDisplay(renderer)
	if err != nil {
		return Threads{}, err
	}
	clientConfig := k8s.ProvideClientConfig()
	kubeContext, err := k8s.ProvideKubeContext(clientConfig)
	if err != nil {
		return Threads{}, err
	}
	env := k8s.ProvideEnv(kubeContext)
	portForwarder := k8s.ProvidePortForwarder()
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, portForwarder, namespace, kubectlRunner, clientConfig)
	podWatcher := engine.NewPodWatcher(client)
	nodeIP, err := k8s.DetectNodeIP(ctx, env)
	if err != nil {
		return Threads{}, err
	}
	serviceWatcher := engine.NewServiceWatcher(client, nodeIP)
	reducer := _wireReducerValue
	storeLogActionsFlag := provideLogActions()
	storeStore := store.NewStore(reducer, storeLogActionsFlag)
	podLogManager := engine.NewPodLogManager(client)
	portForwardController := engine.NewPortForwardController(client)
	fsWatcherMaker := engine.ProvideFsWatcherMaker()
	timerMaker := engine.ProvideTimerMaker()
	watchManager := engine.NewWatchManager(fsWatcherMaker, timerMaker)
	syncletManager := engine.NewSyncletManager(client)
	syncletBuildAndDeployer := engine.NewSyncletBuildAndDeployer(syncletManager, client)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	minikubeClient := minikube.ProvideMinikubeClient()
	dockerEnv, err := docker.ProvideEnv(ctx, env, runtime, minikubeClient)
	if err != nil {
		return Threads{}, err
	}
	clientClient, err := docker.ProvideDockerClient(ctx, dockerEnv)
	if err != nil {
		return Threads{}, err
	}
	version, err := docker.ProvideDockerVersion(ctx, clientClient)
	if err != nil {
		return Threads{}, err
	}
	cli, err := docker.DefaultClient(ctx, clientClient, version)
	if err != nil {
		return Threads{}, err
	}
	containerUpdater := build.NewContainerUpdater(cli)
	analytics, err := provideAnalytics()
	if err != nil {
		return Threads{}, err
	}
	localContainerBuildAndDeployer := engine.NewLocalContainerBuildAndDeployer(containerUpdater, analytics, env)
	labels := _wireLabelsValue
	dockerImageBuilder := build.NewDockerImageBuilder(cli, labels)
	imageBuilder := build.DefaultImageBuilder(dockerImageBuilder)
	cacheBuilder := build.NewCacheBuilder(cli)
	clock := build.ProvideClock()
	execCustomBuilder := build.NewExecCustomBuilder(cli, dockerEnv, clock)
	engineUpdateModeFlag := provideUpdateModeFlag()
	updateMode, err := engine.ProvideUpdateMode(engineUpdateModeFlag, env, runtime)
	if err != nil {
		return Threads{}, err
	}
	imageBuildAndDeployer := engine.NewImageBuildAndDeployer(imageBuilder, cacheBuilder, execCustomBuilder, client, env, analytics, updateMode, clock, runtime)
	dockerComposeClient := dockercompose.NewDockerComposeClient(dockerEnv)
	imageAndCacheBuilder := engine.NewImageAndCacheBuilder(imageBuilder, cacheBuilder, execCustomBuilder, updateMode)
	dockerComposeBuildAndDeployer := engine.NewDockerComposeBuildAndDeployer(dockerComposeClient, cli, imageAndCacheBuilder, clock)
	buildOrder := engine.DefaultBuildOrder(syncletBuildAndDeployer, localContainerBuildAndDeployer, imageBuildAndDeployer, dockerComposeBuildAndDeployer, env, updateMode, runtime)
	compositeBuildAndDeployer := engine.NewCompositeBuildAndDeployer(buildOrder)
	buildController := engine.NewBuildController(compositeBuildAndDeployer)
	imageReaper := build.NewImageReaper(cli)
	imageController := engine.NewImageController(imageReaper)
	globalYAMLBuildController := engine.NewGlobalYAMLBuildController(client)
	tiltfileLoader := tiltfile.ProvideTiltfileLoader(analytics, dockerComposeClient)
	configsController := engine.NewConfigsController(tiltfileLoader)
	dockerComposeEventWatcher := engine.NewDockerComposeEventWatcher(dockerComposeClient)
	dockerComposeLogManager := engine.NewDockerComposeLogManager(dockerComposeClient)
	profilerManager := engine.NewProfilerManager()
	analyticsReporter := engine.ProvideAnalyticsReporter(analytics, storeStore)
	upper := engine.NewUpper(ctx, headsUpDisplay, podWatcher, serviceWatcher, storeStore, podLogManager, portForwardController, watchManager, buildController, imageController, globalYAMLBuildController, configsController, dockerComposeEventWatcher, dockerComposeLogManager, profilerManager, syncletManager, analyticsReporter)
	webMode := provideWebMode()
	webVersion := provideWebVersion()
	assetServer, err := server.ProvideAssetServer(ctx, webMode, webVersion)
	if err != nil {
		return Threads{}, err
	}
	headsUpServer := server.ProvideHeadsUpServer(storeStore, assetServer)
	threads := provideThreads(headsUpDisplay, upper, headsUpServer, assetServer)
	return threads, nil
}

func wireK8sClient(ctx context.Context) (k8s.Client, error) {
	clientConfig := k8s.ProvideClientConfig()
	kubeContext, err := k8s.ProvideKubeContext(clientConfig)
	if err != nil {
		return nil, err
	}
	env := k8s.ProvideEnv(kubeContext)
	portForwarder := k8s.ProvidePortForwarder()
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, portForwarder, namespace, kubectlRunner, clientConfig)
	return client, nil
}

func wireKubeContext(ctx context.Context) (k8s.KubeContext, error) {
	clientConfig := k8s.ProvideClientConfig()
	kubeContext, err := k8s.ProvideKubeContext(clientConfig)
	if err != nil {
		return "", err
	}
	return kubeContext, nil
}

func wireEnv(ctx context.Context) (k8s.Env, error) {
	clientConfig := k8s.ProvideClientConfig()
	kubeContext, err := k8s.ProvideKubeContext(clientConfig)
	if err != nil {
		return "", err
	}
	env := k8s.ProvideEnv(kubeContext)
	return env, nil
}

func wireNamespace(ctx context.Context) (k8s.Namespace, error) {
	clientConfig := k8s.ProvideClientConfig()
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	return namespace, nil
}

func wireRuntime(ctx context.Context) (container.Runtime, error) {
	clientConfig := k8s.ProvideClientConfig()
	kubeContext, err := k8s.ProvideKubeContext(clientConfig)
	if err != nil {
		return "", err
	}
	env := k8s.ProvideEnv(kubeContext)
	portForwarder := k8s.ProvidePortForwarder()
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, portForwarder, namespace, kubectlRunner, clientConfig)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	return runtime, nil
}

func wireK8sVersion(ctx context.Context) (*version.Info, error) {
	clientConfig := k8s.ProvideClientConfig()
	config, err := k8s.ProvideRESTConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	clientset, err := k8s.ProvideClientSet(config)
	if err != nil {
		return nil, err
	}
	info, err := k8s.ProvideServerVersion(clientset)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func wireDockerVersion(ctx context.Context) (types.Version, error) {
	clientConfig := k8s.ProvideClientConfig()
	kubeContext, err := k8s.ProvideKubeContext(clientConfig)
	if err != nil {
		return types.Version{}, err
	}
	env := k8s.ProvideEnv(kubeContext)
	portForwarder := k8s.ProvidePortForwarder()
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, portForwarder, namespace, kubectlRunner, clientConfig)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	minikubeClient := minikube.ProvideMinikubeClient()
	dockerEnv, err := docker.ProvideEnv(ctx, env, runtime, minikubeClient)
	if err != nil {
		return types.Version{}, err
	}
	clientClient, err := docker.ProvideDockerClient(ctx, dockerEnv)
	if err != nil {
		return types.Version{}, err
	}
	typesVersion, err := docker.ProvideDockerVersion(ctx, clientClient)
	if err != nil {
		return types.Version{}, err
	}
	return typesVersion, nil
}

func wireDockerEnv(ctx context.Context) (docker.Env, error) {
	clientConfig := k8s.ProvideClientConfig()
	kubeContext, err := k8s.ProvideKubeContext(clientConfig)
	if err != nil {
		return docker.Env{}, err
	}
	env := k8s.ProvideEnv(kubeContext)
	portForwarder := k8s.ProvidePortForwarder()
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, portForwarder, namespace, kubectlRunner, clientConfig)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	minikubeClient := minikube.ProvideMinikubeClient()
	dockerEnv, err := docker.ProvideEnv(ctx, env, runtime, minikubeClient)
	if err != nil {
		return docker.Env{}, err
	}
	return dockerEnv, nil
}

func wireDownDeps(ctx context.Context) (DownDeps, error) {
	analytics, err := provideAnalytics()
	if err != nil {
		return DownDeps{}, err
	}
	clientConfig := k8s.ProvideClientConfig()
	kubeContext, err := k8s.ProvideKubeContext(clientConfig)
	if err != nil {
		return DownDeps{}, err
	}
	env := k8s.ProvideEnv(kubeContext)
	portForwarder := k8s.ProvidePortForwarder()
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	kubectlRunner := k8s.ProvideKubectlRunner(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, portForwarder, namespace, kubectlRunner, clientConfig)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	minikubeClient := minikube.ProvideMinikubeClient()
	dockerEnv, err := docker.ProvideEnv(ctx, env, runtime, minikubeClient)
	if err != nil {
		return DownDeps{}, err
	}
	dockerComposeClient := dockercompose.NewDockerComposeClient(dockerEnv)
	tiltfileLoader := tiltfile.ProvideTiltfileLoader(analytics, dockerComposeClient)
	downDeps := ProvideDownDeps(tiltfileLoader, dockerComposeClient, client)
	return downDeps, nil
}

// wire.go:

var K8sWireSet = wire.NewSet(k8s.ProvideEnv, k8s.DetectNodeIP, k8s.ProvideKubeContext, k8s.ProvideClientConfig, k8s.ProvideClientSet, k8s.ProvideRESTConfig, k8s.ProvidePortForwarder, k8s.ProvideConfigNamespace, k8s.ProvideKubectlRunner, k8s.ProvideContainerRuntime, k8s.ProvideServerVersion, k8s.ProvideK8sClient)

var BaseWireSet = wire.NewSet(
	K8sWireSet, docker.ProvideDockerClient, docker.ProvideDockerVersion, docker.DefaultClient, wire.Bind(new(docker.Client), new(docker.Cli)), dockercompose.NewDockerComposeClient, build.NewImageReaper, tiltfile.ProvideTiltfileLoader, engine.DeployerWireSet, engine.NewPodLogManager, engine.NewPortForwardController, engine.NewBuildController, engine.NewPodWatcher, engine.NewServiceWatcher, engine.NewImageController, engine.NewConfigsController, engine.NewDockerComposeEventWatcher, engine.NewDockerComposeLogManager, engine.NewProfilerManager, provideClock, hud.NewRenderer, hud.NewDefaultHeadsUpDisplay, provideLogActions, store.NewStore, wire.Bind(new(store.RStore), new(store.Store)), engine.NewUpper, provideAnalytics, engine.ProvideAnalyticsReporter, provideUpdateModeFlag, engine.NewWatchManager, engine.ProvideFsWatcherMaker, engine.ProvideTimerMaker, provideWebVersion,
	provideWebMode, server.ProvideHeadsUpServer, server.ProvideAssetServer, provideThreads,
)

type Threads struct {
	hud         hud.HeadsUpDisplay
	upper       engine.Upper
	server      server.HeadsUpServer
	assetServer server.AssetServer
}

func provideThreads(h hud.HeadsUpDisplay, upper engine.Upper, server2 server.HeadsUpServer, assetServer server.AssetServer) Threads {
	return Threads{h, upper, server2, assetServer}
}

type DownDeps struct {
	tfl      tiltfile.TiltfileLoader
	dcClient dockercompose.DockerComposeClient
	kClient  k8s.Client
}

func ProvideDownDeps(
	tfl tiltfile.TiltfileLoader,
	dcClient dockercompose.DockerComposeClient,
	kClient k8s.Client) DownDeps {
	return DownDeps{
		tfl:      tfl,
		dcClient: dcClient,
		kClient:  kClient,
	}
}

func provideClock() func() time.Time {
	return time.Now
}
