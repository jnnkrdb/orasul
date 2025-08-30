package oci

import (
	"context"

	"github.com/jnnkrdb/orasul/bin/orasul/config"
	"github.com/jnnkrdb/orasul/pkg/logging"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/file"
	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
)

// upload a specific file to the env registered oci registry
func UploadToRegistry(ref, mType, f string) error {

	var ctx = context.Background()

	store, err := file.New(config.Cfg.Local.RegistryPath)
	if err != nil {
		logging.Default.Error("unable to create file store", "err", err)
		return err
	}
	defer store.Close()

	// Datei als Artefakt hinzufügen
	desc, err := store.Add(ctx, f, mType, f)
	if err != nil {
		logging.Default.Error("error adding file", "err", err)
		return err
	}

	// Remote-Repository ansprechen
	repo, err := remote.NewRepository(ref)
	if err != nil {
		logging.Default.Error("error creating remote repository", "err", err)
		return err
	}

	// Falls Authentifizierung benötigt wird:
	repo.Client = &auth.Client{
		Credential: auth.CredentialFunc(func(ctx context.Context, hostport string) (auth.Credential, error) {
			return auth.Credential{
				Username: config.Cfg.Oci.Username,
				Password: config.Cfg.Oci.Password,
			}, nil
		}),
	}

	// Push-Vorgang starten
	manifest, err := oras.Copy(ctx, store, desc.Digest.String(), repo, ref, oras.DefaultCopyOptions)
	if err != nil {
		logging.Default.Error("error pushing to registry", "err", err)
		return err
	}

	logging.Default.Info("successful uploaded file", "digest", manifest.Digest)
	return nil
}
