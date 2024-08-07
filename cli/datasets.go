// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package cli

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/ultravioletrs/cocos/agent"
)

func (cli *CLI) NewDatasetsCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "data",
		Short:   "Upload a dataset CSV file",
		Example: "data <dataset.csv> <private_key_file_path>",
		Args:    cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			datasetFile := args[0]

			log.Println("Uploading dataset CSV:", datasetFile)

			dataset, err := os.ReadFile(datasetFile)
			if err != nil {
				log.Fatalf("Error reading dataset file: %v", err)
			}

			dataReq := agent.Dataset{
				Dataset: dataset,
			}

			privKeyFile, err := os.ReadFile(args[1])
			if err != nil {
				log.Fatalf("Error reading private key file: %v", err)
			}

			pemBlock, _ := pem.Decode(privKeyFile)

			privKey := decodeKey(pemBlock)

			if err := cli.agentSDK.Data(cmd.Context(), dataReq, privKey); err != nil {
				log.Fatalf("Error uploading dataset: %v", err)
			}

			log.Println("Successfully uploaded dataset")
		},
	}
}

func decodeKey(b *pem.Block) interface{} {
	switch b.Type {
	case rsaKeyType:
		privKey, err := x509.ParsePKCS8PrivateKey(b.Bytes)
		if err != nil {
			privKey, err = x509.ParsePKCS1PrivateKey(b.Bytes)
			if err != nil {
				log.Fatalf("Error parsing private key: %v", err)
			}
		}
		return privKey
	case ecdsaKeyType:
		privKey, err := x509.ParseECPrivateKey(b.Bytes)
		if err != nil {
			log.Fatalf("Error parsing private key: %v", err)
		}
		return privKey
	default:
		log.Fatalf("Error decoding key")
		return nil
	}
}
