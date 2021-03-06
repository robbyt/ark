/*
Copyright 2017 the Heptio Ark Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/heptio/ark/pkg/buildinfo"
)

// Config returns a *rest.Config, using either the kubeconfig (if specified) or an in-cluster
// configuration.
func Config(kubeconfig, baseName string) (*rest.Config, error) {
	loader := clientcmd.NewDefaultClientConfigLoadingRules()
	loader.ExplicitPath = kubeconfig
	clientConfig, err := clientcmd.BuildConfigFromKubeconfigGetter("", loader.Load)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	clientConfig.UserAgent = buildUserAgent(
		baseName,
		buildinfo.Version,
		buildinfo.FormattedGitSHA(),
		runtime.GOOS,
		runtime.GOARCH,
	)

	return clientConfig, nil
}

// buildUserAgent builds a User-Agent string from given args.
func buildUserAgent(command, version, formattedSha, os, arch string) string {
	return fmt.Sprintf(
		"%s/%s (%s/%s) %s", command, version, os, arch, formattedSha)
}
