package script

import (
	"context"
	"os/exec"
	"strings"

	"github.com/kyma-incubator/compass/components/formation-watcher/pkg/log"
	"github.com/kyma-incubator/compass/components/formation-watcher/templates"
	"github.com/kyma-incubator/compass/components/formation-watcher/types"
	"github.com/pkg/errors"
)

type Runner struct {
	ScriptsLocation string
	Resolver        templates.Resolver
}

func (r *Runner) Run(ctx context.Context, scriptName string, args ...string) error {
	_, err := r.RunWithOutput(ctx, scriptName, args...)
	return err
}

func (r *Runner) RunWithOutput(ctx context.Context, scriptName string, args ...string) (string, error) {
	cmdArgs := append([]string{r.ScriptsLocation + "/" + scriptName})
	if len(args) > 0 {
		cmdArgs = append(cmdArgs, args...)
	}
	log.C(ctx).Infof("Executing %s %s", scriptName, strings.Join(args, " "))

	cmd := exec.Command("/bin/bash", cmdArgs...)
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Errorf("failed executing script %s: %s: %s", scriptName, err, string(outputBytes))
	}

	output := string(outputBytes)
	log.C(ctx).Infof("[=== %s ===]", output)

	return output, nil
}

func (r *Runner) ApplyDependency(ctx context.Context, dep types.Dependency, remote string) error {
	dependencyString := r.Resolver.ResolveDependency(dep)
	return r.Run(ctx, "dependency_applier.sh", dependencyString, remote)
}

func (r *Runner) DeleteDependency(ctx context.Context, dep string, remote1, remote2 string) error {
	return r.Run(ctx, "dependency_cleaner.sh", dep, remote1, remote2)
}

func (r *Runner) ApplyResource(ctx context.Context, resourceFile string) error {
	return r.Run(ctx, "resource_applier.sh", resourceFile)
}

func (r *Runner) DeleteResource(ctx context.Context, resourceFile string) error {
	return r.Run(ctx, "resource_cleaner.sh", resourceFile)
}

func (r *Runner) RegisterRuntime(ctx context.Context, remote string) error {
	return r.Run(ctx, "register_consumer_cluster.sh", remote)
}

func (r *Runner) RegisterApplication(ctx context.Context, remote string) error {
	return r.Run(ctx, "register_provider_cluster.sh", remote)
}

func (r *Runner) DeleteRuntime(ctx context.Context, remote string) error {
	return r.Run(ctx, "cleanup_remote_consumer_cluster.sh", remote)
}

func (r *Runner) DeleteApplication(ctx context.Context, remote1, remote2 string) error {
	return r.Run(ctx, "cleanup_remote_provider_cluster.sh", remote1, remote2)
}

func (r *Runner) GetExistingServices(ctx context.Context) ([]string, error) {
	output, err := r.RunWithOutput(ctx, "get_services.sh")
	if err != nil {
		return nil, err
	}

	if output == "" {
		return []string{}, nil
	}

	return strings.Split(output, " "), nil
}

func (r *Runner) DependencyExists(ctx context.Context, dep string) (bool, error) {
	output, err := r.RunWithOutput(ctx, "dependency_exists.sh", dep)
	if err != nil {
		return false, err
	}

	return strings.Contains(output, dep), nil
}
