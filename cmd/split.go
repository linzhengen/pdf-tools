package cmd

import (
	"fmt"
	pdfapi "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/spf13/cobra"
	"path/filepath"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Spit PDF pages",
	RunE: func(cmd *cobra.Command, args []string) error {
		outDir := GetFlagString(cmd, "out-dir")
		inFile := GetFlagString(cmd, "in-file")
		if outDir == "" {
			outDir = filepath.Dir(inFile)
		}
		fmt.Sprintf("outout to %s", outDir)
		return pdfapi.SplitFile(inFile, outDir, 1, nil)
	},
}

func init() {
	splitCmd.Flags().String("in-file", "", "Path to the input PDF file.")
	splitCmd.Flags().String("out-dir", "", "Output directory to save the split PDF files. By default, it will be added to the same directory as the input file.")
	splitCmd.MarkFlagRequired("in-file")
	rootCmd.AddCommand(splitCmd)
}
