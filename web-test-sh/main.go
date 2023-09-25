package main

import (
	"context"
	"fmt"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/sigstore/cosign/v2/cmd/cosign/cli/options"
	ociremote "github.com/sigstore/cosign/v2/pkg/oci/remote"
)

func main() {
	key := "akurniko/gensbom_alpine"
	//	key := "scribesecuriy.jfrog.io/scribe-docker-public-local/test/gensbom_alpine_input:latest"
	//	key := "scribesecuriy.jfrog.io/scribe-docker-public-local/test/gensbom_alpine_input@sha256:862df2564ad7da1b91fc8854e9235bad0f2e93fb752fec89e4b6c4e403d32074"
	fmt.Println("verify signature for:", key)
	ref, err := name.ParseReference(key)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	fmt.Printf("ref %v\n", ref.Identifier())

	ctx := context.Background()
	ro := options.RegistryOptions{
		Keychain: authn.DefaultKeychain,
	}
	co, err := ro.ClientOpts(ctx)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	platform, err := v1.ParsePlatform("linux/arm64/v8")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	//co = append(co, remote.WithPlatform(platform))
	img, err := remote.Image(ref, remote.WithPlatform(*platform))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	iId, err := img.ConfigName()
	if err != nil {
		fmt.Printf("Error config name: %v", err)
		return
	}
	fmt.Printf("img config name id: %v\n", iId)
	m, err := img.Manifest()

	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	fmt.Printf("image digest %v %v\n", m.Config.Digest)

	digest, err := ociremote.ResolveDigest(ref, co...)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	fmt.Printf("digest id %v\n", digest.Identifier())

}
